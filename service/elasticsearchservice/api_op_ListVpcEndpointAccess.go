// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves information about each principal that is allowed to access a given
// Amazon OpenSearch Service domain through the use of an interface VPC endpoint.
func (c *Client) ListVpcEndpointAccess(ctx context.Context, params *ListVpcEndpointAccessInput, optFns ...func(*Options)) (*ListVpcEndpointAccessOutput, error) {
	if params == nil {
		params = &ListVpcEndpointAccessInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListVpcEndpointAccess", params, optFns, c.addOperationListVpcEndpointAccessMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListVpcEndpointAccessOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Retrieves information about each principal that is allowed to access a given
// Amazon OpenSearch Service domain through the use of an interface VPC endpoint
type ListVpcEndpointAccessInput struct {

	// The name of the OpenSearch Service domain to retrieve access information for.
	//
	// This member is required.
	DomainName *string

	// Provides an identifier to allow retrieval of paginated results.
	NextToken *string

	noSmithyDocumentSerde
}

// Container for response parameters to the ListVpcEndpointAccess operation.
// Returns a list of accounts id and account type authorized to manage VPC
// endpoints.
type ListVpcEndpointAccessOutput struct {

	// List of AuthorizedPrincipal describing the details of the permissions to manage
	// VPC endpoints against the specified domain.
	//
	// This member is required.
	AuthorizedPrincipalList []types.AuthorizedPrincipal

	// Provides an identifier to allow retrieval of paginated results.
	//
	// This member is required.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListVpcEndpointAccessMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListVpcEndpointAccess{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListVpcEndpointAccess{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListVpcEndpointAccessValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListVpcEndpointAccess(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opListVpcEndpointAccess(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "ListVpcEndpointAccess",
	}
}