// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// CreateAccountURL generates an URL for the create account operation
type CreateAccountURL struct {
	ClusterID string
	Username  string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *CreateAccountURL) WithBasePath(bp string) *CreateAccountURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *CreateAccountURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *CreateAccountURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/clusters/{clusterId}/account/{username}"

	clusterID := o.ClusterID
	if clusterID != "" {
		_path = strings.Replace(_path, "{clusterId}", clusterID, -1)
	} else {
		return nil, errors.New("ClusterID is required on CreateAccountURL")
	}

	username := o.Username
	if username != "" {
		_path = strings.Replace(_path, "{username}", username, -1)
	} else {
		return nil, errors.New("Username is required on CreateAccountURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/v1"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *CreateAccountURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *CreateAccountURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *CreateAccountURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on CreateAccountURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on CreateAccountURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *CreateAccountURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
