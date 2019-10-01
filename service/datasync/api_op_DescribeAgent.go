// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package datasync

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// DescribeAgent
type DescribeAgentInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the agent to describe.
	//
	// AgentArn is a required field
	AgentArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeAgentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeAgentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeAgentInput"}

	if s.AgentArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("AgentArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// DescribeAgentResponse
type DescribeAgentOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the agent.
	AgentArn *string `type:"string"`

	// The time that the agent was activated (that is, created in your account).
	CreationTime *time.Time `type:"timestamp"`

	// The type of endpoint that your agent is connected to. If the endpoint is
	// a VPC endpoint, the agent is not accessible over the public Internet.
	EndpointType EndpointType `type:"string" enum:"true"`

	// The time that the agent last connected to DataSyc.
	LastConnectionTime *time.Time `type:"timestamp"`

	// The name of the agent.
	Name *string `min:"1" type:"string"`

	// The VPC endpoint, subnet and security group that an agent uses to access
	// IP addresses in a VPC (Virtual Private Cloud).
	PrivateLinkConfig *PrivateLinkConfig `type:"structure"`

	// The status of the agent. If the status is ONLINE, then the agent is configured
	// properly and is available to use. The Running status is the normal running
	// status for an agent. If the status is OFFLINE, the agent's VM is turned off
	// or the agent is in an unhealthy state. When the issue that caused the unhealthy
	// state is resolved, the agent returns to ONLINE status.
	Status AgentStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s DescribeAgentOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeAgent = "DescribeAgent"

// DescribeAgentRequest returns a request value for making API operation for
// AWS DataSync.
//
// Returns metadata such as the name, the network interfaces, and the status
// (that is, whether the agent is running or not) for an agent. To specify which
// agent to describe, use the Amazon Resource Name (ARN) of the agent in your
// request.
//
//    // Example sending a request using DescribeAgentRequest.
//    req := client.DescribeAgentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datasync-2018-11-09/DescribeAgent
func (c *Client) DescribeAgentRequest(input *DescribeAgentInput) DescribeAgentRequest {
	op := &aws.Operation{
		Name:       opDescribeAgent,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeAgentInput{}
	}

	req := c.newRequest(op, input, &DescribeAgentOutput{})
	return DescribeAgentRequest{Request: req, Input: input, Copy: c.DescribeAgentRequest}
}

// DescribeAgentRequest is the request type for the
// DescribeAgent API operation.
type DescribeAgentRequest struct {
	*aws.Request
	Input *DescribeAgentInput
	Copy  func(*DescribeAgentInput) DescribeAgentRequest
}

// Send marshals and sends the DescribeAgent API request.
func (r DescribeAgentRequest) Send(ctx context.Context) (*DescribeAgentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAgentResponse{
		DescribeAgentOutput: r.Request.Data.(*DescribeAgentOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeAgentResponse is the response type for the
// DescribeAgent API operation.
type DescribeAgentResponse struct {
	*DescribeAgentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeAgent request.
func (r *DescribeAgentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}