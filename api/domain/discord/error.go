package discord

import "errors"

var (
	ErrDiscordServerNotFound       = errors.New("this Discord server doesn't exist")
	ErrDiscordServerCreationFailed = errors.New("an error occurred while creating the Discord server")
	ErrDiscordServerUpdateFailed   = errors.New("an error occurred while updating the Discord server")
	ErrDiscordServerDeletionFailed = errors.New("an error occurred while deleting the Discord server")
	ErrDiscordServerAlreadyExists  = errors.New("this Discord server already exists")
)
