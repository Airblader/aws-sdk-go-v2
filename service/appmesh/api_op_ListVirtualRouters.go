// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appmesh

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListVirtualRoutersInput struct {
	_ struct{} `type:"structure"`

	Limit *int64 `location:"querystring" locationName:"limit" min:"1" type:"integer"`

	// MeshName is a required field
	MeshName *string `location:"uri" locationName:"meshName" min:"1" type:"string" required:"true"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListVirtualRoutersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListVirtualRoutersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListVirtualRoutersInput"}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}

	if s.MeshName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MeshName"))
	}
	if s.MeshName != nil && len(*s.MeshName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MeshName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListVirtualRoutersInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MeshName != nil {
		v := *s.MeshName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "meshName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Limit != nil {
		v := *s.Limit

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "limit", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListVirtualRoutersOutput struct {
	_ struct{} `type:"structure"`

	NextToken *string `locationName:"nextToken" type:"string"`

	// VirtualRouters is a required field
	VirtualRouters []VirtualRouterRef `locationName:"virtualRouters" type:"list" required:"true"`
}

// String returns the string representation
func (s ListVirtualRoutersOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListVirtualRoutersOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VirtualRouters != nil {
		v := s.VirtualRouters

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "virtualRouters", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListVirtualRouters = "ListVirtualRouters"

// ListVirtualRoutersRequest returns a request value for making API operation for
// AWS App Mesh.
//
// Returns a list of existing virtual routers in a service mesh.
//
//    // Example sending a request using ListVirtualRoutersRequest.
//    req := client.ListVirtualRoutersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/ListVirtualRouters
func (c *Client) ListVirtualRoutersRequest(input *ListVirtualRoutersInput) ListVirtualRoutersRequest {
	op := &aws.Operation{
		Name:       opListVirtualRouters,
		HTTPMethod: "GET",
		HTTPPath:   "/v20190125/meshes/{meshName}/virtualRouters",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListVirtualRoutersInput{}
	}

	req := c.newRequest(op, input, &ListVirtualRoutersOutput{})
	return ListVirtualRoutersRequest{Request: req, Input: input, Copy: c.ListVirtualRoutersRequest}
}

// ListVirtualRoutersRequest is the request type for the
// ListVirtualRouters API operation.
type ListVirtualRoutersRequest struct {
	*aws.Request
	Input *ListVirtualRoutersInput
	Copy  func(*ListVirtualRoutersInput) ListVirtualRoutersRequest
}

// Send marshals and sends the ListVirtualRouters API request.
func (r ListVirtualRoutersRequest) Send(ctx context.Context) (*ListVirtualRoutersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListVirtualRoutersResponse{
		ListVirtualRoutersOutput: r.Request.Data.(*ListVirtualRoutersOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListVirtualRoutersRequestPaginator returns a paginator for ListVirtualRouters.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListVirtualRoutersRequest(input)
//   p := appmesh.NewListVirtualRoutersRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListVirtualRoutersPaginator(req ListVirtualRoutersRequest) ListVirtualRoutersPaginator {
	return ListVirtualRoutersPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListVirtualRoutersInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListVirtualRoutersPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListVirtualRoutersPaginator struct {
	aws.Pager
}

func (p *ListVirtualRoutersPaginator) CurrentPage() *ListVirtualRoutersOutput {
	return p.Pager.CurrentPage().(*ListVirtualRoutersOutput)
}

// ListVirtualRoutersResponse is the response type for the
// ListVirtualRouters API operation.
type ListVirtualRoutersResponse struct {
	*ListVirtualRoutersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListVirtualRouters request.
func (r *ListVirtualRoutersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}