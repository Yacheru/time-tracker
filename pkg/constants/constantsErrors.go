package constants

import "errors"

var (
	ErrEmptyVar = errors.New("required variable not defined")

	ErrPeopleNotFound    = errors.New("no people found")
	FailedValidateParams = errors.New("failed validate params")
	ErrPeopleExist       = errors.New("people already exist")
	ErrCheckExist        = errors.New("people exist check failed")
	InvalidSurname       = errors.New("invalid surname or does not exist")
	InvalidName          = errors.New("invalid name or does not exist")
	InvalidSeries        = errors.New("invalid passport series or does not exist")
	InvalidNumber        = errors.New("invalid passport number or does not exist")
	ErrStartTask         = errors.New("failed to start task")
	ErrStopTask          = errors.New("failed to stop task")
	ErrGetAllTasks       = errors.New("failed to get all tasks")
	NoActiveTask         = errors.New("no active task")
	HaveActiveTask       = errors.New("already have active task")

	FailedParseBody = errors.New("failed to parse request body")
)
