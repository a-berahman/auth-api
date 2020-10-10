// Package classification of Auth API
//
// Documentation for Auth API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package handler

import "github.com/a-berahman/auth-api/model"

// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handlers

// swagger:parameters getToken
type getTokenRequestParamsWrapper struct {
	// GetTokenRequest data structure
	// in: body
	// required: true
	Body model.TokenRequest
}

// Data structure generate token response
// swagger:response getTokenResponse
type getTokenResponseWrapper struct {
	// GetTokenResponse data structure
	// in: body
	// required: true
	Body model.TokenResponse
}

// swagger:parameters validateToken
type validateTokenRequestParamsWrapper struct {
	// ValidateTokenRequest data structure
	// in: body
	// required: true
	Body model.ValidateRequest
}

// Data structure validate token response
// swagger:response validateTokenResponse
type validateTokenResponseWrapper struct {
	// ValidaTokenResponse data structure
	// in: body
	// required: true
	Body model.ValidateResponse
}

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	//GenericError data structure
	//in: body
	Body model.GenericError
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}
