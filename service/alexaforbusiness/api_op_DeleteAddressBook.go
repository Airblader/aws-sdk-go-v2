// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteAddressBookInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the address book to delete.
	//
	// AddressBookArn is a required field
	AddressBookArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteAddressBookInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteAddressBookInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteAddressBookInput"}

	if s.AddressBookArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("AddressBookArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteAddressBookOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteAddressBookOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteAddressBook = "DeleteAddressBook"

// DeleteAddressBookRequest returns a request value for making API operation for
// Alexa For Business.
//
// Deletes an address book by the address book ARN.
//
//    // Example sending a request using DeleteAddressBookRequest.
//    req := client.DeleteAddressBookRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/DeleteAddressBook
func (c *Client) DeleteAddressBookRequest(input *DeleteAddressBookInput) DeleteAddressBookRequest {
	op := &aws.Operation{
		Name:       opDeleteAddressBook,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteAddressBookInput{}
	}

	req := c.newRequest(op, input, &DeleteAddressBookOutput{})
	return DeleteAddressBookRequest{Request: req, Input: input, Copy: c.DeleteAddressBookRequest}
}

// DeleteAddressBookRequest is the request type for the
// DeleteAddressBook API operation.
type DeleteAddressBookRequest struct {
	*aws.Request
	Input *DeleteAddressBookInput
	Copy  func(*DeleteAddressBookInput) DeleteAddressBookRequest
}

// Send marshals and sends the DeleteAddressBook API request.
func (r DeleteAddressBookRequest) Send(ctx context.Context) (*DeleteAddressBookResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteAddressBookResponse{
		DeleteAddressBookOutput: r.Request.Data.(*DeleteAddressBookOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteAddressBookResponse is the response type for the
// DeleteAddressBook API operation.
type DeleteAddressBookResponse struct {
	*DeleteAddressBookOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteAddressBook request.
func (r *DeleteAddressBookResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}