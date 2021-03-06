// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "dummy-secrets": Application Contexts
//
// Command:
// $ goagen
// --design=github.com/julianvilas/dummy-secrets/design
// --out=$(GOPATH)/src/github.com/julianvilas/dummy-secrets
// --version=v1.1.0-dirty

package app

import (
	"context"
	"github.com/goadesign/goa"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

// CreateSecretsContext provides the Secrets create action context.
type CreateSecretsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *SecretPayload
}

// NewCreateSecretsContext parses the incoming request URL and body, performs validations and creates the
// context used by the Secrets controller create action.
func NewCreateSecretsContext(ctx context.Context, r *http.Request, service *goa.Service) (*CreateSecretsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := CreateSecretsContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateSecretsContext) Created() error {
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *CreateSecretsContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *CreateSecretsContext) InternalServerError(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// ShowSecretsContext provides the Secrets show action context.
type ShowSecretsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID uuid.UUID
}

// NewShowSecretsContext parses the incoming request URL and body, performs validations and creates the
// context used by the Secrets controller show action.
func NewShowSecretsContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowSecretsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowSecretsContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		if id, err2 := uuid.FromString(rawID); err2 == nil {
			rctx.ID = id
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("id", rawID, "uuid"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowSecretsContext) OK(r *Secret) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.secret+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// Created sends a HTTP response with status code 201.
func (ctx *ShowSecretsContext) Created() error {
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ShowSecretsContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowSecretsContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ShowSecretsContext) InternalServerError(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}
