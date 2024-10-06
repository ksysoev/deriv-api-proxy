package core

import (
	"context"
	"errors"
	"fmt"
	"html/template"
	"sync"
)

var ErrIterDone = errors.New("iteration done")

type CallHandler struct {
	calls map[string]*CallRunConfig
}

type CallRunConfig struct {
	params   map[string]string
	requests map[string]*RequestRunConfig
}

type RequestRunConfig struct {
	tmplt        *template.Template
	allow        []string
	responseBody string
}

type TemplateData struct {
	Params map[string]any
	ReqID  int64
}

func NewCallHandler(config *Config) (*CallHandler, error) {
	h := &CallHandler{
		calls: make(map[string]*CallRunConfig),
	}

	for _, call := range config.Calls {
		c := &CallRunConfig{
			requests: make(map[string]*RequestRunConfig),
		}
		for _, req := range call.Backend {
			tmplt, err := template.New("request").Parse(req.RequestTemplate)
			if err != nil {
				return nil, err
			}
			c.requests[req.ResponseBody] = &RequestRunConfig{
				tmplt:        tmplt,
				allow:        req.Allow,
				responseBody: req.ResponseBody,
			}
		}
		h.calls[call.Method] = c
	}

	return h, nil
}

func (h *CallHandler) Process(req *Request) (*RequesIter, error) {
	method := req.RoutingKey()

	call, ok := h.calls[method]

	if !ok {
		return nil, fmt.Errorf("unsupported method: %s", method)
	}

	requests := make([]*RequestProcessor, 0, len(call.requests))
	for _, req := range call.requests {
		requests = append(requests, &RequestProcessor{
			tempate:      req.tmplt,
			allow:        req.allow,
			responseBody: req.responseBody,
		})
	}

	ctx, cancel := context.WithCancel(req.Context())

	return &RequesIter{
		ctx:       ctx,
		cancel:    cancel,
		reqs:      requests,
		params:    req.Params,
		finalResp: make(map[string]any),
		composer:  NewComposer(),
	}, nil
}

type RequesIter struct {
	ctx       context.Context
	cancel    context.CancelFunc
	pos       int
	reqs      []*RequestProcessor
	params    map[string]any
	finalResp map[string]any
	err       error
	mu        sync.Mutex
	composer  *Composer
}

func (r *RequesIter) HasNext() bool {
	return r.pos < len(r.reqs)
}

func (r *RequesIter) Next(id int64, respChan <-chan []byte) ([]byte, error) {
	if r.pos >= len(r.reqs) {
		return nil, ErrIterDone
	}

	req := r.reqs[r.pos]
	r.pos++

	data := TemplateData{
		Params: r.params,
		ReqID:  id,
	}

	body, err := req.Render(data)
	if err != nil {
		return nil, fmt.Errorf("fail to render request %s: %w", req.responseBody, err)
	}

	go r.composer.WaitResponse(r.ctx, req, respChan)

	return body, nil
}

func (r *RequesIter) WaitResp(req_id *int) ([]byte, error) {
	return r.composer.Response(req_id)
}
