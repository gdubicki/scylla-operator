// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewMessagingServiceMessagesRespondCompletedGetParams creates a new MessagingServiceMessagesRespondCompletedGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMessagingServiceMessagesRespondCompletedGetParams() *MessagingServiceMessagesRespondCompletedGetParams {
	return &MessagingServiceMessagesRespondCompletedGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMessagingServiceMessagesRespondCompletedGetParamsWithTimeout creates a new MessagingServiceMessagesRespondCompletedGetParams object
// with the ability to set a timeout on a request.
func NewMessagingServiceMessagesRespondCompletedGetParamsWithTimeout(timeout time.Duration) *MessagingServiceMessagesRespondCompletedGetParams {
	return &MessagingServiceMessagesRespondCompletedGetParams{
		timeout: timeout,
	}
}

// NewMessagingServiceMessagesRespondCompletedGetParamsWithContext creates a new MessagingServiceMessagesRespondCompletedGetParams object
// with the ability to set a context for a request.
func NewMessagingServiceMessagesRespondCompletedGetParamsWithContext(ctx context.Context) *MessagingServiceMessagesRespondCompletedGetParams {
	return &MessagingServiceMessagesRespondCompletedGetParams{
		Context: ctx,
	}
}

// NewMessagingServiceMessagesRespondCompletedGetParamsWithHTTPClient creates a new MessagingServiceMessagesRespondCompletedGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewMessagingServiceMessagesRespondCompletedGetParamsWithHTTPClient(client *http.Client) *MessagingServiceMessagesRespondCompletedGetParams {
	return &MessagingServiceMessagesRespondCompletedGetParams{
		HTTPClient: client,
	}
}

/*
MessagingServiceMessagesRespondCompletedGetParams contains all the parameters to send to the API endpoint

	for the messaging service messages respond completed get operation.

	Typically these are written to a http.Request.
*/
type MessagingServiceMessagesRespondCompletedGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the messaging service messages respond completed get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MessagingServiceMessagesRespondCompletedGetParams) WithDefaults() *MessagingServiceMessagesRespondCompletedGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the messaging service messages respond completed get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MessagingServiceMessagesRespondCompletedGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the messaging service messages respond completed get params
func (o *MessagingServiceMessagesRespondCompletedGetParams) WithTimeout(timeout time.Duration) *MessagingServiceMessagesRespondCompletedGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the messaging service messages respond completed get params
func (o *MessagingServiceMessagesRespondCompletedGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the messaging service messages respond completed get params
func (o *MessagingServiceMessagesRespondCompletedGetParams) WithContext(ctx context.Context) *MessagingServiceMessagesRespondCompletedGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the messaging service messages respond completed get params
func (o *MessagingServiceMessagesRespondCompletedGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the messaging service messages respond completed get params
func (o *MessagingServiceMessagesRespondCompletedGetParams) WithHTTPClient(client *http.Client) *MessagingServiceMessagesRespondCompletedGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the messaging service messages respond completed get params
func (o *MessagingServiceMessagesRespondCompletedGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *MessagingServiceMessagesRespondCompletedGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}