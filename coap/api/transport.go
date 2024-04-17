// Copyright (c) Abstract Machines
// SPDX-License-Identifier: Apache-2.0

package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/absmach/magistrala"
	"github.com/absmach/magistrala/coap"
	"github.com/absmach/magistrala/pkg/errors"
	svcerr "github.com/absmach/magistrala/pkg/errors/service"
	"github.com/absmach/magistrala/pkg/messaging"
	"github.com/go-chi/chi/v5"
	"github.com/plgd-dev/go-coap/v3/message"
	"github.com/plgd-dev/go-coap/v3/message/codes"
	"github.com/plgd-dev/go-coap/v3/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	protocol     = "coap"
	authQuery    = "auth"
	startObserve = 0 // observe option value that indicates start of observation
)

var channelPartRegExp = regexp.MustCompile(`^/channels/([\w\-]+)/messages(/[^?]*)?(\?.*)?$`)

const (
	numGroups    = 3 // entire expression + channel group + subtopic group
	channelGroup = 2 // channel group is second in channel regexp
)

var (
	errMalformedSubtopic = errors.New("malformed subtopic")
	errBadOptions        = errors.New("bad options")
)

var (
	logger  *slog.Logger
	service coap.Service
)

// MakeHandler returns a HTTP handler for API endpoints.
func MakeHandler(instanceID string) http.Handler {
	b := chi.NewRouter()
	b.Get("/health", magistrala.Health(protocol, instanceID))
	b.Handle("/metrics", promhttp.Handler())

	return b
}

// MakeCoAPHandler creates handler for CoAP messages.
func MakeCoAPHandler(svc coap.Service, l *slog.Logger) mux.HandlerFunc {
	logger = l
	service = svc

	return handler
}

func sendResp(w mux.ResponseWriter, resp *message.Message, isObs bool) {
	pm := w.Conn().AcquireMessage(w.Conn().Context())
	defer w.Conn().ReleaseMessage(pm)
	pm.SetCode(resp.Code)
	pm.SetToken(resp.Token)
	pm.SetBody(bytes.NewReader(resp.Payload))
	for _, opt := range resp.Options {
		pm.SetOptionBytes(opt.ID, opt.Value)
	}
	if !isObs {
		if err := w.Conn().WriteMessage(pm); err != nil {
			logger.Warn(fmt.Sprintf("Can't set response: %s", err))
		}
	}
}

func handler(w mux.ResponseWriter, m *mux.Message) {
	resp := message.Message{
		Code:    codes.Content,
		Token:   m.Token(),
		Options: make(message.Options, 0, 16),
	}
	msg, err := decodeMessage(m)
	if err != nil {
		logger.Warn(fmt.Sprintf("Error decoding message: %s", err))
		resp.Code = codes.BadRequest
		sendResp(w, &resp, false)
		return
	}
	key, err := parseKey(m)
	if err != nil {
		logger.Warn(fmt.Sprintf("Error parsing auth: %s", err))
		resp.Code = codes.Unauthorized
		sendResp(w, &resp, false)
		return
	}
	var isObs bool = false
	switch m.Code() {
	case codes.GET:
		isObs, err = handleGet(context.Background(), m, w, msg, key)
	case codes.POST:
		resp.Code = codes.Created
		err = service.Publish(m.Context(), key, msg)
	default:
		err = svcerr.ErrNotFound
	}
	defer sendResp(w, &resp, isObs)
	if err != nil {
		switch {
		case err == errBadOptions:
			resp.Code = codes.BadOption
		case err == svcerr.ErrNotFound:
			resp.Code = codes.NotFound
		case errors.Contains(err, svcerr.ErrAuthorization),
			errors.Contains(err, svcerr.ErrAuthentication):
			resp.Code = codes.Unauthorized
		default:
			resp.Code = codes.InternalServerError
		}
	}
}

func handleGet(ctx context.Context, m *mux.Message, c mux.ResponseWriter, msg *messaging.Message, key string) (bool, error) {
	var obs uint32
	obs, err := m.Observe()
	if err != nil {
		logger.Warn(fmt.Sprintf("Error reading observe option: %s", err))
		return false, errBadOptions
	}
	if obs == startObserve {
		c := coap.NewClient(c, m.Token(), logger)
		return true, service.Subscribe(ctx, key, msg.GetChannel(), msg.GetSubtopic(), c)
	}
	return false, service.Unsubscribe(ctx, key, msg.GetChannel(), msg.GetSubtopic(), m.Token().String())
}

func decodeMessage(msg *mux.Message) (*messaging.Message, error) {
	if msg.Options() == nil {
		return &messaging.Message{}, errBadOptions
	}
	path, err := msg.Path()
	if err != nil {
		return &messaging.Message{}, err
	}
	channelParts := channelPartRegExp.FindStringSubmatch(path)
	if len(channelParts) < numGroups {
		return &messaging.Message{}, errMalformedSubtopic
	}

	st, err := parseSubtopic(channelParts[channelGroup])
	if err != nil {
		return &messaging.Message{}, err
	}
	ret := &messaging.Message{
		Protocol: protocol,
		Channel:  channelParts[1],
		Subtopic: st,
		Payload:  []byte{},
		Created:  time.Now().UnixNano(),
	}

	if msg.Body() != nil {
		buff, err := io.ReadAll(msg.Body())
		if err != nil {
			return ret, err
		}
		ret.Payload = buff
	}
	return ret, nil
}

func parseKey(msg *mux.Message) (string, error) {
	if obs, _ := msg.Observe(); obs != 0 && msg.Code() == codes.GET {
		return "", nil
	}
	authKey, err := msg.Options().GetString(message.URIQuery)
	if err != nil {
		return "", err
	}
	vars := strings.Split(authKey, "=")
	if len(vars) != 2 || vars[0] != authQuery {
		return "", svcerr.ErrAuthorization
	}
	return vars[1], nil
}

func parseSubtopic(subtopic string) (string, error) {
	if subtopic == "" {
		return subtopic, nil
	}

	subtopic, err := url.QueryUnescape(subtopic)
	if err != nil {
		return "", errMalformedSubtopic
	}
	subtopic = strings.ReplaceAll(subtopic, "/", ".")

	elems := strings.Split(subtopic, ".")
	filteredElems := []string{}
	for _, elem := range elems {
		if elem == "" {
			continue
		}

		if len(elem) > 1 && (strings.Contains(elem, "*") || strings.Contains(elem, ">")) {
			return "", errMalformedSubtopic
		}

		filteredElems = append(filteredElems, elem)
	}

	subtopic = strings.Join(filteredElems, ".")
	return subtopic, nil
}
