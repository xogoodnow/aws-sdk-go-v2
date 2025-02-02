// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Ends a game session that's currently in progress. You can use this action to
// terminate any game session that isn't in TERMINATED or TERMINATING status.
// Terminating a game session is the most efficient way to free up a server process
// when it's hosting a game session that's in a bad state or not ending naturally.
// You can use this action to terminate a game session that's being hosted on any
// type of Amazon GameLift fleet compute, including computes for managed EC2,
// managed container, and Anywhere fleets.
//
// There are two potential methods for terminating a game session:
//
//   - With a graceful termination, the Amazon GameLift service prompts the server
//     process to initiate its normal game session shutdown sequence. This sequence is
//     implemented in the game server code and might involve a variety of actions to
//     gracefully end a game session, such as notifying players, and stop the server
//     process.
//
//   - With a forceful termination, the Amazon GameLift service takes immediate
//     action to terminate the game session by stopping the server process. Termination
//     occurs without the normal game session shutdown sequence.
//
// Request options
//
//   - Request termination for a single game session. Provide the game session ID
//     and the termination method.
//
// # Results
//
// If successful, game session termination is initiated, which includes changing
// the game session status to TERMINATING . As a result of this action, and
// depending on the implementation of OnProcessTerminate() , the server process
// either becomes available to host a new game session, or it's recycled and a new
// server process started with availability to host a game session. The game
// session status is changed to TERMINATED , with a status reason that indicates
// the termination method used.
func (c *Client) TerminateGameSession(ctx context.Context, params *TerminateGameSessionInput, optFns ...func(*Options)) (*TerminateGameSessionOutput, error) {
	if params == nil {
		params = &TerminateGameSessionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TerminateGameSession", params, optFns, c.addOperationTerminateGameSessionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TerminateGameSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TerminateGameSessionInput struct {

	// A unique identifier for the game session to be terminated. A game session ARN
	// has the following format: arn:aws:gamelift:::gamesession// .
	//
	// This member is required.
	GameSessionId *string

	// The method to use to terminate the game session. Available methods include:
	//
	//   - TRIGGER_ON_PROCESS_TERMINATE – Sends an OnProcessTerminate() callback to the
	//   server process to initiate the normal game session shutdown sequence. At a
	//   minimum, the callback method must include a call to the server SDK action
	//   ProcessEnding() , which is how the server process signals that a game session
	//   is ending. If the server process doesn't call ProcessEnding() , this
	//   termination method won't be successful.
	//
	//   - FORCE_TERMINATE – Takes action to stop the server process, using existing
	//   methods to control how server processes run on an Amazon GameLift managed
	//   compute.
	//
	// This method is not available for game sessions that are running on Anywhere
	//   fleets unless the fleet is deployed with the Amazon GameLift Agent. In this
	//   scenario, a force terminate request results in an invalid or bad request
	//   exception.
	//
	// This member is required.
	TerminationMode types.TerminationMode

	noSmithyDocumentSerde
}

type TerminateGameSessionOutput struct {

	// Properties describing a game session.
	//
	// A game session in ACTIVE status can host players. When a game session ends, its
	// status is set to TERMINATED .
	//
	// Amazon GameLift retains a game session resource for 30 days after the game
	// session ends. You can reuse idempotency token values after this time. Game
	// session logs are retained for 14 days.
	//
	// [All APIs by task]
	//
	// [All APIs by task]: https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets
	GameSession *types.GameSession

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationTerminateGameSessionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpTerminateGameSession{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpTerminateGameSession{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "TerminateGameSession"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addSpanRetryLoop(stack, options); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpTerminateGameSessionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTerminateGameSession(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opTerminateGameSession(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "TerminateGameSession",
	}
}
