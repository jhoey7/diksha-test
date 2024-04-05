package utils

// ErrReqInvalid list error
var (
	ErrReqInvalid         = 400
	ErrNone               = 200
	ErrUserNotFound       = 404
	ErrPasswordNotMatch   = 405
	ErrUserAlreadyExist   = 406
	ErrIncorrectPassword  = 407
	ErrMaxCounterPassword = 408
	ErrDefault            = 508

	MsgUserAlreadyExist   = "The mdn you specified is already in use"
	MsgUserNotFound       = "User Not Found"
	MsgIncorrectPassword  = "Incorrect Password"
	MsgErrDefault         = "Don't worry, we are handling this"
	MsgPasswordNotMatch   = "Password Not Match."
	MsgMaxCounterPassword = "You have reached the maximum password input limit"
)
