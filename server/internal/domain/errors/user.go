package errors

import "errors"

var (
	ErrUserNotFound        = errors.New("This user doesn't exist")
	ErrUserCreationFailed  = errors.New("An error occurred while creating the user")
	ErrUserUpdateFailed    = errors.New("An error occurred while updating the user")
	ErrUserDeletionFailed  = errors.New("An error occurred while deleting the user")
	ErrUserAlreadyExists   = errors.New("This user already exists")
	ErrUserNotActive       = errors.New("This user is not active")
	ErrUserNotVerified     = errors.New("This user is not verified")
	ErrUserInvalidPassword = errors.New("This password is invalid")
	ErrUserInvalidEmail    = errors.New("This email is invalid")
	ErrUserInvalidUsername = errors.New("This username is invalid")
	ErrUserInvalidRole     = errors.New("This role is invalid")
	ErrUserInvalidID       = errors.New("This user ID is invalid")
)
