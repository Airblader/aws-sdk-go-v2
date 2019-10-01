// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package athena

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StopQueryExecutionInput struct {
	_ struct{} `type:"structure"`

	// The unique ID of the query execution to stop.
	//
	// QueryExecutionId is a required field
	QueryExecutionId *string `type:"string" required:"true" idempotencyToken:"true"`
}

// String returns the string representation
func (s StopQueryExecutionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopQueryExecutionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopQueryExecutionInput"}

	if s.QueryExecutionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("QueryExecutionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StopQueryExecutionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StopQueryExecutionOutput) String() string {
	return awsutil.Prettify(s)
}

const opStopQueryExecution = "StopQueryExecution"

// StopQueryExecutionRequest returns a request value for making API operation for
// Amazon Athena.
//
// Stops a query execution. Requires you to have access to the workgroup in
// which the query ran.
//
// For code samples using the AWS SDK for Java, see Examples and Code Samples
// (http://docs.aws.amazon.com/athena/latest/ug/code-samples.html) in the Amazon
// Athena User Guide.
//
//    // Example sending a request using StopQueryExecutionRequest.
//    req := client.StopQueryExecutionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/athena-2017-05-18/StopQueryExecution
func (c *Client) StopQueryExecutionRequest(input *StopQueryExecutionInput) StopQueryExecutionRequest {
	op := &aws.Operation{
		Name:       opStopQueryExecution,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StopQueryExecutionInput{}
	}

	req := c.newRequest(op, input, &StopQueryExecutionOutput{})
	return StopQueryExecutionRequest{Request: req, Input: input, Copy: c.StopQueryExecutionRequest}
}

// StopQueryExecutionRequest is the request type for the
// StopQueryExecution API operation.
type StopQueryExecutionRequest struct {
	*aws.Request
	Input *StopQueryExecutionInput
	Copy  func(*StopQueryExecutionInput) StopQueryExecutionRequest
}

// Send marshals and sends the StopQueryExecution API request.
func (r StopQueryExecutionRequest) Send(ctx context.Context) (*StopQueryExecutionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopQueryExecutionResponse{
		StopQueryExecutionOutput: r.Request.Data.(*StopQueryExecutionOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopQueryExecutionResponse is the response type for the
// StopQueryExecution API operation.
type StopQueryExecutionResponse struct {
	*StopQueryExecutionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopQueryExecution request.
func (r *StopQueryExecutionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}