// Code generated by smithy-go-codegen DO NOT EDIT.

package verifiedpermissions

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/verifiedpermissions/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Modifies a Cedar static policy in the specified policy store. You can change
// only certain elements of the UpdatePolicyDefinition (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_UpdatePolicyInput.html#amazonverifiedpermissions-UpdatePolicy-request-UpdatePolicyDefinition)
// parameter. You can directly update only static policies. To change a
// template-linked policy, you must update the template instead, using
// UpdatePolicyTemplate (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_UpdatePolicyTemplate.html)
// . If policy validation is enabled in the policy store, then updating a static
// policy causes Verified Permissions to validate the policy against the schema in
// the policy store. If the updated static policy doesn't pass validation, the
// operation fails and the update isn't stored.
func (c *Client) UpdatePolicy(ctx context.Context, params *UpdatePolicyInput, optFns ...func(*Options)) (*UpdatePolicyOutput, error) {
	if params == nil {
		params = &UpdatePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdatePolicy", params, optFns, c.addOperationUpdatePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdatePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdatePolicyInput struct {

	// Specifies the updated policy content that you want to replace on the specified
	// policy. The content must be valid Cedar policy language text. You can change
	// only the following elements from the policy definition:
	//   - The action referenced by the policy.
	//   - Any conditional clauses, such as when or unless clauses.
	// You can't change the following elements:
	//   - Changing from static to templateLinked .
	//   - Changing the effect of the policy from permit or forbid .
	//   - The principal referenced by the policy.
	//   - The resource referenced by the policy.
	//
	// This member is required.
	Definition types.UpdatePolicyDefinition

	// Specifies the ID of the policy that you want to update. To find this value, you
	// can use ListPolicies (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_ListPolicies.html)
	// .
	//
	// This member is required.
	PolicyId *string

	// Specifies the ID of the policy store that contains the policy that you want to
	// update.
	//
	// This member is required.
	PolicyStoreId *string

	noSmithyDocumentSerde
}

type UpdatePolicyOutput struct {

	// The date and time that the policy was originally created.
	//
	// This member is required.
	CreatedDate *time.Time

	// The date and time that the policy was most recently updated.
	//
	// This member is required.
	LastUpdatedDate *time.Time

	// The ID of the policy that was updated.
	//
	// This member is required.
	PolicyId *string

	// The ID of the policy store that contains the policy that was updated.
	//
	// This member is required.
	PolicyStoreId *string

	// The type of the policy that was updated.
	//
	// This member is required.
	PolicyType types.PolicyType

	// The principal specified in the policy's scope. This element isn't included in
	// the response when Principal isn't present in the policy content.
	Principal *types.EntityIdentifier

	// The resource specified in the policy's scope. This element isn't included in
	// the response when Resource isn't present in the policy content.
	Resource *types.EntityIdentifier

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdatePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdatePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdatePolicy{}, middleware.After)
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
	if err = addOpUpdatePolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePolicy(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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

func newServiceMetadataMiddleware_opUpdatePolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "verifiedpermissions",
		OperationName: "UpdatePolicy",
	}
}