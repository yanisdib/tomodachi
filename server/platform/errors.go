package platform

import "errors"

var (
	ErrDuplicate    = errors.New("this record already exists")
	ErrNotFound     = errors.New("this record does not exist")
	ErrCreateFailed = errors.New("an error occured while trying to create the platform")
	ErrUpdateFailed = errors.New("an error occured while trying to update this platform")
	ErrDeleteFailed = errors.New("an error occured while trying to delete this platform")
)
