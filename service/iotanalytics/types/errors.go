// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// There was an internal failure.
type InternalFailureException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InternalFailureException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalFailureException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalFailureException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "InternalFailureException"
	}
	return *e.ErrorCodeOverride
}
func (e *InternalFailureException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The request was not valid.
type InvalidRequestException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidRequestException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "InvalidRequestException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The command caused an internal limit to be exceeded.
type LimitExceededException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "LimitExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A resource with the same name already exists.
type ResourceAlreadyExistsException struct {
	Message *string

	ErrorCodeOverride *string

	ResourceId  *string
	ResourceArn *string

	noSmithyDocumentSerde
}

func (e *ResourceAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceAlreadyExistsException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "ResourceAlreadyExistsException"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A resource with the specified name could not be found.
type ResourceNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "ResourceNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The service is temporarily unavailable.
type ServiceUnavailableException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ServiceUnavailableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceUnavailableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceUnavailableException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "ServiceUnavailableException"
	}
	return *e.ErrorCodeOverride
}
func (e *ServiceUnavailableException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The request was denied due to request throttling.
type ThrottlingException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ThrottlingException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ThrottlingException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ThrottlingException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "ThrottlingException"
	}
	return *e.ErrorCodeOverride
}
func (e *ThrottlingException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
