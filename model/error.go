package model

//GenericError us a generic error message return by a server
type GenericError struct {
	Message string `json:"message"`
}
