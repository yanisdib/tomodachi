package user

import "errors"

var (
	ErrUserNotFound        = errors.New("this user doesn't exist")
	ErrUserCreationFailed  = errors.New("an error occurred while creating the user")
	ErrUserUpdateFailed    = errors.New("an error occurred while updating the user")
	ErrUserDeletionFailed  = errors.New("an error occurred while deleting the user")
	ErrUserAlreadyExists   = errors.New("this user already exists")
	ErrUserNotActive       = errors.New("this user is not active")
	ErrUserNotVerified     = errors.New("this user is not verified")
	ErrUserInvalidPassword = errors.New("this password is invalid")
	ErrUserInvalidEmail    = errors.New("this email is invalid")
	ErrUserInvalidUsername = errors.New("this username is invalid")
	ErrUserInvalidRole     = errors.New("this role is invalid")
	ErrUserInvalidID       = errors.New("this user ID is invalid")
)
