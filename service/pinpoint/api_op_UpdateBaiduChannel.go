// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateBaiduChannelInput struct {
	_ struct{} `type:"structure" payload:"BaiduChannelRequest"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`

	// Specifies the status and settings of the Baidu (Baidu Cloud Push) channel
	// for an application.
	//
	// BaiduChannelRequest is a required field
	BaiduChannelRequest *BaiduChannelRequest `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateBaiduChannelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateBaiduChannelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateBaiduChannelInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if s.BaiduChannelRequest == nil {
		invalidParams.Add(aws.NewErrParamRequired("BaiduChannelRequest"))
	}
	if s.BaiduChannelRequest != nil {
		if err := s.BaiduChannelRequest.Validate(); err != nil {
			invalidParams.AddNested("BaiduChannelRequest", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateBaiduChannelInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BaiduChannelRequest != nil {
		v := s.BaiduChannelRequest

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "BaiduChannelRequest", v, metadata)
	}
	return nil
}

type UpdateBaiduChannelOutput struct {
	_ struct{} `type:"structure" payload:"BaiduChannelResponse"`

	// Provides information about the status and settings of the Baidu (Baidu Cloud
	// Push) channel for an application.
	//
	// BaiduChannelResponse is a required field
	BaiduChannelResponse *BaiduChannelResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateBaiduChannelOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateBaiduChannelOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.BaiduChannelResponse != nil {
		v := s.BaiduChannelResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "BaiduChannelResponse", v, metadata)
	}
	return nil
}

const opUpdateBaiduChannel = "UpdateBaiduChannel"

// UpdateBaiduChannelRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Updates the settings of the Baidu channel for an application.
//
//    // Example sending a request using UpdateBaiduChannelRequest.
//    req := client.UpdateBaiduChannelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/UpdateBaiduChannel
func (c *Client) UpdateBaiduChannelRequest(input *UpdateBaiduChannelInput) UpdateBaiduChannelRequest {
	op := &aws.Operation{
		Name:       opUpdateBaiduChannel,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/apps/{application-id}/channels/baidu",
	}

	if input == nil {
		input = &UpdateBaiduChannelInput{}
	}

	req := c.newRequest(op, input, &UpdateBaiduChannelOutput{})
	return UpdateBaiduChannelRequest{Request: req, Input: input, Copy: c.UpdateBaiduChannelRequest}
}

// UpdateBaiduChannelRequest is the request type for the
// UpdateBaiduChannel API operation.
type UpdateBaiduChannelRequest struct {
	*aws.Request
	Input *UpdateBaiduChannelInput
	Copy  func(*UpdateBaiduChannelInput) UpdateBaiduChannelRequest
}

// Send marshals and sends the UpdateBaiduChannel API request.
func (r UpdateBaiduChannelRequest) Send(ctx context.Context) (*UpdateBaiduChannelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateBaiduChannelResponse{
		UpdateBaiduChannelOutput: r.Request.Data.(*UpdateBaiduChannelOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateBaiduChannelResponse is the response type for the
// UpdateBaiduChannel API operation.
type UpdateBaiduChannelResponse struct {
	*UpdateBaiduChannelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateBaiduChannel request.
func (r *UpdateBaiduChannelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}