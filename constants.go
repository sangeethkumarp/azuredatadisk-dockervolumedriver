package main

const (
	// ErrorInternal - Internal error occured. Contact the developer with the logs.
	ErrorInternal = -1
	// ErrorInvalidArgument - Invalid argument passed by the user
	ErrorInvalidArgument = -2
	// ErrorSocketFileFailure - Failure in listening to the Unix socket file
	ErrorSocketFileFailure = -3
)

// UnixSocketFileName - Name of the unix socket file that will be used by this volume driver
const UnixSocketFileName = "azureDataDisk"
