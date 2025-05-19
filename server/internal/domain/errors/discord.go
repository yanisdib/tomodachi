package errors

import "errors"

var (
	ErrDiscordServerNotFound       = errors.New("This Discord server doesn't exist")
	ErrDiscordServerCreationFailed = errors.New("An error occurred while creating the Discord server")
	ErrDiscordServerUpdateFailed   = errors.New("An error occurred while updating the Discord server")
	ErrDiscordServerDeletionFailed = errors.New("An error occurred while deleting the Discord server")
	ErrDiscordServerAlreadyExists  = errors.New("This Discord server already exists")
)
