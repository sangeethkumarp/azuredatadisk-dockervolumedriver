package main

import (
	"fmt"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

// AzureSubscriptionCred represents the credentials required to perform ARM REST
// API request on behalf of an Azure subscription.
type AzureSubscriptionCred struct {
	ClientID       string
	ClientSecret   string
	SubscriptionID string
	TenantID       string
}

// Validate performs preliminary check of the credentials.
func (subsCred AzureSubscriptionCred) Validate() error {
	var err error

	// TODO-sangeethkumarp: Add strict typing checks, since most fields are GUID.
	err = validateInputArg(&subsCred.ClientID, "Client ID", err)
	err = validateInputArg(&subsCred.ClientSecret, "Client Secret", err)
	err = validateInputArg(&subsCred.SubscriptionID, "Subscription ID", err)
	err = validateInputArg(&subsCred.TenantID, "Tenant ID", err)

	return err
}

// AzureStorageCred represents the credentials required to perform operations on
// an Azurestorage account.
type AzureStorageCred struct {
	AccountName   string
	AccountKey    string
	ContainerName string
}

// Validate performs preliminary check of the credentials.
func (storCred AzureStorageCred) Validate() error {
	var err error

	// TODO-sangeethkumarp: Add strict typing checks, since most fields are GUID.
	err = validateInputArg(&storCred.AccountName, "Storage account name", err)
	err = validateInputArg(&storCred.AccountKey, "Storage account key", err)
	err = validateInputArg(&storCred.ContainerName, "Storage container name", err)

	return err
}

// Notice that this method also trims the extra space on front and back, along
// with the empty space check.
func validateInputArg(input *string, name string, prevError error) error {
	// Chained error check simplification for reduced LOC
	if prevError != nil {
		return prevError
	}

	if name == "" || input == nil {
		return makeError(ErrorInternal, "Unexpeted name or input argument")
	} else if *input = strings.TrimSpace(*input); *input == "" {
		return makeError(ErrorInvalidArgument, "'%s' shouldn't be empty!", name)
	}

	return nil
}

func makeError(errorCode int, format string, args ...interface{}) *cli.ExitError {
	return makeErrorWithFields(log.Fields{}, errorCode, format, args...)
}

func makeErrorFromErr(errorCode int, err error, format string, args ...interface{}) *cli.ExitError {
	return makeErrorWithFields(log.Fields{"Reason": err}, errorCode, format, args...)
}

func makeErrorWithFields(fields log.Fields, errorCode int, format string, args ...interface{}) *cli.ExitError {

	// Format the string with provided args
	formattedStr := fmt.Sprintf(format, args...)

	// Log the error
	fields["ErrorCode"] = errorCode
	log.WithFields(fields).Error(formattedStr)

	// Construct the error object
	return cli.NewExitError(formattedStr, errorCode)
}
