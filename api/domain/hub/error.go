package hub

import "errors"

var (
	ErrHubNotFound       = errors.New("this hub doesn't exist")
	ErrHubCreationFailed = errors.New("an error occurred while creating the hub")
	ErrHubUpdateFailed   = errors.New("an error occurred while updating the hub")
	ErrHubDeletionFailed = errors.New("an error occurred while deleting the hub")
	ErrHubAlreadyExists  = errors.New("this hub already exists")
	ErrHubNotActive      = errors.New("this hub is closed or inactive")
	ErrHubNotPublic      = errors.New("this hub is not open to the public")
	ErrHubBadPassword    = errors.New("this password is incorrect")
	ErrHubAccessDenied   = errors.New("you don't have access to this hub")
	ErrHubInvalidID      = errors.New("this hub ID is invalid")
	ErrHubsFetchFailed   = errors.New("an error occurred while fetching the hubs")
	ErrNoHubsFound       = errors.New("no hubs found")
)
