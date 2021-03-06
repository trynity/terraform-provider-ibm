// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteVpcsIDParams creates a new DeleteVpcsIDParams object
// with the default values initialized.
func NewDeleteVpcsIDParams() *DeleteVpcsIDParams {
	var ()
	return &DeleteVpcsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteVpcsIDParamsWithTimeout creates a new DeleteVpcsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteVpcsIDParamsWithTimeout(timeout time.Duration) *DeleteVpcsIDParams {
	var ()
	return &DeleteVpcsIDParams{

		timeout: timeout,
	}
}

// NewDeleteVpcsIDParamsWithContext creates a new DeleteVpcsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteVpcsIDParamsWithContext(ctx context.Context) *DeleteVpcsIDParams {
	var ()
	return &DeleteVpcsIDParams{

		Context: ctx,
	}
}

// NewDeleteVpcsIDParamsWithHTTPClient creates a new DeleteVpcsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteVpcsIDParamsWithHTTPClient(client *http.Client) *DeleteVpcsIDParams {
	var ()
	return &DeleteVpcsIDParams{
		HTTPClient: client,
	}
}

/*DeleteVpcsIDParams contains all the parameters to send to the API endpoint
for the delete vpcs ID operation typically these are written to a http.Request
*/
type DeleteVpcsIDParams struct {

	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*ID
	  The VPC identifier

	*/
	ID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete vpcs ID params
func (o *DeleteVpcsIDParams) WithTimeout(timeout time.Duration) *DeleteVpcsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete vpcs ID params
func (o *DeleteVpcsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete vpcs ID params
func (o *DeleteVpcsIDParams) WithContext(ctx context.Context) *DeleteVpcsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete vpcs ID params
func (o *DeleteVpcsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete vpcs ID params
func (o *DeleteVpcsIDParams) WithHTTPClient(client *http.Client) *DeleteVpcsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete vpcs ID params
func (o *DeleteVpcsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGeneration adds the generation to the delete vpcs ID params
func (o *DeleteVpcsIDParams) WithGeneration(generation int64) *DeleteVpcsIDParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the delete vpcs ID params
func (o *DeleteVpcsIDParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithID adds the id to the delete vpcs ID params
func (o *DeleteVpcsIDParams) WithID(id string) *DeleteVpcsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete vpcs ID params
func (o *DeleteVpcsIDParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the delete vpcs ID params
func (o *DeleteVpcsIDParams) WithVersion(version string) *DeleteVpcsIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete vpcs ID params
func (o *DeleteVpcsIDParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteVpcsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param generation
	qrGeneration := o.Generation
	qGeneration := swag.FormatInt64(qrGeneration)
	if qGeneration != "" {
		if err := r.SetQueryParam("generation", qGeneration); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// query param version
	qrVersion := o.Version
	qVersion := qrVersion
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
