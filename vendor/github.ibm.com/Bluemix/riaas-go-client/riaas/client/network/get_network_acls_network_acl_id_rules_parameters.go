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

// NewGetNetworkAclsNetworkACLIDRulesParams creates a new GetNetworkAclsNetworkACLIDRulesParams object
// with the default values initialized.
func NewGetNetworkAclsNetworkACLIDRulesParams() *GetNetworkAclsNetworkACLIDRulesParams {
	var (
		limitDefault = int32(50)
	)
	return &GetNetworkAclsNetworkACLIDRulesParams{
		Limit: &limitDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkAclsNetworkACLIDRulesParamsWithTimeout creates a new GetNetworkAclsNetworkACLIDRulesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworkAclsNetworkACLIDRulesParamsWithTimeout(timeout time.Duration) *GetNetworkAclsNetworkACLIDRulesParams {
	var (
		limitDefault = int32(50)
	)
	return &GetNetworkAclsNetworkACLIDRulesParams{
		Limit: &limitDefault,

		timeout: timeout,
	}
}

// NewGetNetworkAclsNetworkACLIDRulesParamsWithContext creates a new GetNetworkAclsNetworkACLIDRulesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworkAclsNetworkACLIDRulesParamsWithContext(ctx context.Context) *GetNetworkAclsNetworkACLIDRulesParams {
	var (
		limitDefault = int32(50)
	)
	return &GetNetworkAclsNetworkACLIDRulesParams{
		Limit: &limitDefault,

		Context: ctx,
	}
}

// NewGetNetworkAclsNetworkACLIDRulesParamsWithHTTPClient creates a new GetNetworkAclsNetworkACLIDRulesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworkAclsNetworkACLIDRulesParamsWithHTTPClient(client *http.Client) *GetNetworkAclsNetworkACLIDRulesParams {
	var (
		limitDefault = int32(50)
	)
	return &GetNetworkAclsNetworkACLIDRulesParams{
		Limit:      &limitDefault,
		HTTPClient: client,
	}
}

/*GetNetworkAclsNetworkACLIDRulesParams contains all the parameters to send to the API endpoint
for the get network acls network ACL ID rules operation typically these are written to a http.Request
*/
type GetNetworkAclsNetworkACLIDRulesParams struct {

	/*Direction
	  Filters the collection to rules with the specified direction

	*/
	Direction *string
	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*Limit
	  The number of resources to return on a page

	*/
	Limit *int32
	/*NetworkACLID
	  The network ACL identifier

	*/
	NetworkACLID string
	/*Start
	  A server-supplied token determining what resource to start the page on

	*/
	Start *string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get network acls network ACL ID rules params
func (o *GetNetworkAclsNetworkACLIDRulesParams) WithTimeout(timeout time.Duration) *GetNetworkAclsNetworkACLIDRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network acls network ACL ID rules params
func (o *GetNetworkAclsNetworkACLIDRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network acls network ACL ID rules params
func (o *GetNetworkAclsNetworkACLIDRulesParams) WithContext(ctx context.Context) *GetNetworkAclsNetworkACLIDRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network acls network ACL ID rules params
func (o *GetNetworkAclsNetworkACLIDRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network acls network ACL ID rules params
func (o *GetNetworkAclsNetworkACLIDRulesParams) WithHTTPClient(client *http.Client) *GetNetworkAclsNetworkACLIDRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network acls network ACL ID rules params
func (o *GetNetworkAclsNetworkACLIDRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDirection adds the direction to the get network acls network ACL ID rules params
func (o *GetNetworkAclsNetworkACLIDRulesParams) WithDirection(direction *string) *GetNetworkAclsNetworkACLIDRulesParams {
	o.SetDirection(direction)
	return o
}

// SetDirection adds the direction to the get network acls network ACL ID rules params
func (o *GetNetworkAclsNetworkACLIDRulesParams) SetDirection(direction *string) {
	o.Direction = direction
}

// WithGeneration adds the generation to the get network acls network ACL ID rules params
func (o *GetNetworkAclsNetworkACLIDRulesParams) WithGeneration(generation int64) *GetNetworkAclsNetworkACLIDRulesParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the get network acls network ACL ID rules params
func (o *GetNetworkAclsNetworkACLIDRulesParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithLimit adds the limit to the get network acls network ACL ID rules params
func (o *GetNetworkAclsNetworkACLIDRulesParams) WithLimit(limit *int32) *GetNetworkAclsNetworkACLIDRulesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get network acls network ACL ID rules params
func (o *GetNetworkAclsNetworkACLIDRulesParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithNetworkACLID adds the networkACLID to the get network acls network ACL ID rules params
func (o *GetNetworkAclsNetworkACLIDRulesParams) WithNetworkACLID(networkACLID string) *GetNetworkAclsNetworkACLIDRulesParams {
	o.SetNetworkACLID(networkACLID)
	return o
}

// SetNetworkACLID adds the networkAclId to the get network acls network ACL ID rules params
func (o *GetNetworkAclsNetworkACLIDRulesParams) SetNetworkACLID(networkACLID string) {
	o.NetworkACLID = networkACLID
}

// WithStart adds the start to the get network acls network ACL ID rules params
func (o *GetNetworkAclsNetworkACLIDRulesParams) WithStart(start *string) *GetNetworkAclsNetworkACLIDRulesParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get network acls network ACL ID rules params
func (o *GetNetworkAclsNetworkACLIDRulesParams) SetStart(start *string) {
	o.Start = start
}

// WithVersion adds the version to the get network acls network ACL ID rules params
func (o *GetNetworkAclsNetworkACLIDRulesParams) WithVersion(version string) *GetNetworkAclsNetworkACLIDRulesParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get network acls network ACL ID rules params
func (o *GetNetworkAclsNetworkACLIDRulesParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkAclsNetworkACLIDRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Direction != nil {

		// query param direction
		var qrDirection string
		if o.Direction != nil {
			qrDirection = *o.Direction
		}
		qDirection := qrDirection
		if qDirection != "" {
			if err := r.SetQueryParam("direction", qDirection); err != nil {
				return err
			}
		}

	}

	// query param generation
	qrGeneration := o.Generation
	qGeneration := swag.FormatInt64(qrGeneration)
	if qGeneration != "" {
		if err := r.SetQueryParam("generation", qGeneration); err != nil {
			return err
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	// path param network_acl_id
	if err := r.SetPathParam("network_acl_id", o.NetworkACLID); err != nil {
		return err
	}

	if o.Start != nil {

		// query param start
		var qrStart string
		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := qrStart
		if qStart != "" {
			if err := r.SetQueryParam("start", qStart); err != nil {
				return err
			}
		}

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
