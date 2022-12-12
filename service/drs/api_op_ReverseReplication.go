// Code generated by smithy-go-codegen DO NOT EDIT.

package drs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Start replication to origin / target region - applies only to protected
// instances that originated in EC2. For recovery instances on target region -
// starts replication back to origin region. For failback instances on origin
// region - starts replication to target region to re-protect them.
func (c *Client) ReverseReplication(ctx context.Context, params *ReverseReplicationInput, optFns ...func(*Options)) (*ReverseReplicationOutput, error) {
	if params == nil {
		params = &ReverseReplicationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ReverseReplication", params, optFns, c.addOperationReverseReplicationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ReverseReplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ReverseReplicationInput struct {

	// The ID of the Recovery Instance that we want to reverse the replication for.
	//
	// This member is required.
	RecoveryInstanceID *string

	noSmithyDocumentSerde
}

type ReverseReplicationOutput struct {

	// ARN of created SourceServer.
	ReversedDirectionSourceServerArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationReverseReplicationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpReverseReplication{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpReverseReplication{}, middleware.After)
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
	if err = addOpReverseReplicationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opReverseReplication(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opReverseReplication(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "drs",
		OperationName: "ReverseReplication",
	}
}