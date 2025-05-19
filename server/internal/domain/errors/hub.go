package errors

import "errors"

var (
	ErrHubNotFound       = errors.New("This hub doesn't exist")
	ErrHubCreationFailed = errors.New("An error occurred while creating the hub")
	ErrHubUpdateFailed   = errors.New("An error occurred while updating the hub")
	ErrHubDeletionFailed = errors.New("An error occurred while deleting the hub")
	ErrHubAlreadyExists  = errors.New("This hub already exists")
	ErrHubNotActive      = errors.New("This hub is closed or inactive")
	ErrHubNotPublic      = errors.New("This hub is not open to the public")
	ErrHubBadPassword    = errors.New("This password is incorrect")
	ErrHubAccessDenied   = errors.New("You don't have access to this hub")
	ErrHubInvalidID      = errors.New("This hub ID is invalid")
)
