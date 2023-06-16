// Code generated by smithy-go-codegen DO NOT EDIT.

package paymentcryptography

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/paymentcryptography/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the aliases for all keys in the caller's Amazon Web Services account and
// Amazon Web Services Region. You can filter the list of aliases. For more
// information, see Using aliases (https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-managealias.html)
// in the Amazon Web Services Payment Cryptography User Guide. This is a paginated
// operation, which means that each response might contain only a subset of all the
// aliases. When the response contains only a subset of aliases, it includes a
// NextToken value. Use this value in a subsequent ListAliases request to get more
// aliases. When you receive a response with no NextToken (or an empty or null
// value), that means there are no more aliases to get. Cross-account use: This
// operation can't be used across different Amazon Web Services accounts. Related
// operations:
//   - CreateAlias
//   - DeleteAlias
//   - GetAlias
//   - UpdateAlias
func (c *Client) ListAliases(ctx context.Context, params *ListAliasesInput, optFns ...func(*Options)) (*ListAliasesOutput, error) {
	if params == nil {
		params = &ListAliasesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAliases", params, optFns, c.addOperationListAliasesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAliasesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAliasesInput struct {

	// Use this parameter to specify the maximum number of items to return. When this
	// value is present, Amazon Web Services Payment Cryptography does not return more
	// than the specified number of items, but it might return fewer. This value is
	// optional. If you include a value, it must be between 1 and 100, inclusive. If
	// you do not include a value, it defaults to 50.
	MaxResults *int32

	// Use this parameter in a subsequent request after you receive a response with
	// truncated results. Set it to the value of NextToken from the truncated response
	// you just received.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAliasesOutput struct {

	// The list of aliases. Each alias describes the KeyArn contained within.
	//
	// This member is required.
	Aliases []types.Alias

	// The token for the next set of results, or an empty or null value if there are
	// no more results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAliasesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListAliases{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListAliases{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAliases(options.Region), middleware.Before); err != nil {
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

// ListAliasesAPIClient is a client that implements the ListAliases operation.
type ListAliasesAPIClient interface {
	ListAliases(context.Context, *ListAliasesInput, ...func(*Options)) (*ListAliasesOutput, error)
}

var _ ListAliasesAPIClient = (*Client)(nil)

// ListAliasesPaginatorOptions is the paginator options for ListAliases
type ListAliasesPaginatorOptions struct {
	// Use this parameter to specify the maximum number of items to return. When this
	// value is present, Amazon Web Services Payment Cryptography does not return more
	// than the specified number of items, but it might return fewer. This value is
	// optional. If you include a value, it must be between 1 and 100, inclusive. If
	// you do not include a value, it defaults to 50.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAliasesPaginator is a paginator for ListAliases
type ListAliasesPaginator struct {
	options   ListAliasesPaginatorOptions
	client    ListAliasesAPIClient
	params    *ListAliasesInput
	nextToken *string
	firstPage bool
}

// NewListAliasesPaginator returns a new ListAliasesPaginator
func NewListAliasesPaginator(client ListAliasesAPIClient, params *ListAliasesInput, optFns ...func(*ListAliasesPaginatorOptions)) *ListAliasesPaginator {
	if params == nil {
		params = &ListAliasesInput{}
	}

	options := ListAliasesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAliasesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAliasesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAliases page.
func (p *ListAliasesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAliasesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListAliases(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListAliases(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "payment-cryptography",
		OperationName: "ListAliases",
	}
}