package common

const (
	// Verify Error
	VerifyFailed = 1

	// Auth Error (1xx)
	AuthAccountExists    = 100
	AuthAccountNotFound  = 101
	AuthAccountNotActive = 102
	AuthAccountBlocked   = 103

	// Request Error (4xx)
	RequestValidationError = 400
	RequestUnauthorized    = 401
	RequestForbidden       = 403
	RequestNotFound        = 404

	// Server Error (5xx)
	ServerError     = 500
	ServerAuthError = 501
)
