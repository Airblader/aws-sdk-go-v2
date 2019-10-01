// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetObjectTorrentInput struct {
	_ struct{} `type:"structure"`

	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// Key is a required field
	Key *string `location:"uri" locationName:"Key" min:"1" type:"string" required:"true"`

	// Confirms that the requester knows that she or he will be charged for the
	// request. Bucket owners need not specify this parameter in their requests.
	// Documentation on downloading objects from requester pays buckets can be found
	// at http://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html
	RequestPayer RequestPayer `location:"header" locationName:"x-amz-request-payer" type:"string" enum:"true"`
}

// String returns the string representation
func (s GetObjectTorrentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetObjectTorrentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetObjectTorrentInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *GetObjectTorrentInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetObjectTorrentInput) MarshalFields(e protocol.FieldEncoder) error {

	if len(s.RequestPayer) > 0 {
		v := s.RequestPayer

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-request-payer", v, metadata)
	}
	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.Key != nil {
		v := *s.Key

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Key", protocol.StringValue(v), metadata)
	}
	return nil
}

type GetObjectTorrentOutput struct {
	_ struct{} `type:"structure" payload:"Body"`

	Body io.ReadCloser `type:"blob"`

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged RequestCharged `location:"header" locationName:"x-amz-request-charged" type:"string" enum:"true"`
}

// String returns the string representation
func (s GetObjectTorrentOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetObjectTorrentOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.RequestCharged) > 0 {
		v := s.RequestCharged

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-request-charged", v, metadata)
	}
	// Skipping Body Output type's body not valid.
	return nil
}

const opGetObjectTorrent = "GetObjectTorrent"

// GetObjectTorrentRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Return torrent files from a bucket.
//
//    // Example sending a request using GetObjectTorrentRequest.
//    req := client.GetObjectTorrentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetObjectTorrent
func (c *Client) GetObjectTorrentRequest(input *GetObjectTorrentInput) GetObjectTorrentRequest {
	op := &aws.Operation{
		Name:       opGetObjectTorrent,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}/{Key+}?torrent",
	}

	if input == nil {
		input = &GetObjectTorrentInput{}
	}

	req := c.newRequest(op, input, &GetObjectTorrentOutput{})
	return GetObjectTorrentRequest{Request: req, Input: input, Copy: c.GetObjectTorrentRequest}
}

// GetObjectTorrentRequest is the request type for the
// GetObjectTorrent API operation.
type GetObjectTorrentRequest struct {
	*aws.Request
	Input *GetObjectTorrentInput
	Copy  func(*GetObjectTorrentInput) GetObjectTorrentRequest
}

// Send marshals and sends the GetObjectTorrent API request.
func (r GetObjectTorrentRequest) Send(ctx context.Context) (*GetObjectTorrentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetObjectTorrentResponse{
		GetObjectTorrentOutput: r.Request.Data.(*GetObjectTorrentOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetObjectTorrentResponse is the response type for the
// GetObjectTorrent API operation.
type GetObjectTorrentResponse struct {
	*GetObjectTorrentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetObjectTorrent request.
func (r *GetObjectTorrentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}