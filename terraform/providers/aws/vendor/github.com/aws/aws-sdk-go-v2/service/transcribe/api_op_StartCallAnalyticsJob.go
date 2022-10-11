// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Transcribes the audio from a customer service call and applies any additional
// Request Parameters you choose to include in your request. In addition to many of
// the standard transcription features, Call Analytics provides you with call
// characteristics, call summarization, speaker sentiment, and optional redaction
// of your text transcript and your audio file. You can also apply custom
// categories to flag specified conditions. To learn more about these features and
// insights, refer to Analyzing call center audio with Call Analytics
// (https://docs.aws.amazon.com/transcribe/latest/dg/call-analytics.html). If you
// want to apply categories to your Call Analytics job, you must create them before
// submitting your job request. Categories cannot be retroactively applied to a
// job. To create a new category, use the operation. To learn more about Call
// Analytics categories, see Creating categories
// (https://docs.aws.amazon.com/transcribe/latest/dg/call-analytics-create-categories.html).
// To make a StartCallAnalyticsJob request, you must first upload your media file
// into an Amazon S3 bucket; you can then specify the Amazon S3 location of the
// file using the Media parameter. You must include the following parameters in
// your StartCallAnalyticsJob request:
//
// * region: The Amazon Web Services Region
// where you are making your request. For a list of Amazon Web Services Regions
// supported with Amazon Transcribe, refer to Amazon Transcribe endpoints and
// quotas (https://docs.aws.amazon.com/general/latest/gr/transcribe.html).
//
// *
// CallAnalyticsJobName: A custom name you create for your transcription job that
// is unique within your Amazon Web Services account.
//
// * DataAccessRoleArn: The
// Amazon Resource Name (ARN) of an IAM role that has permissions to access the
// Amazon S3 bucket that contains your input files.
//
// * Media (MediaFileUri or
// RedactedMediaFileUri): The Amazon S3 location of your media file.
//
// With Call
// Analytics, you can redact the audio contained in your media file by including
// RedactedMediaFileUri, instead of MediaFileUri, to specify the location of your
// input audio. If you choose to redact your audio, you can find your redacted
// media at the location specified in the RedactedMediaFileUri field of your
// response.
func (c *Client) StartCallAnalyticsJob(ctx context.Context, params *StartCallAnalyticsJobInput, optFns ...func(*Options)) (*StartCallAnalyticsJobOutput, error) {
	if params == nil {
		params = &StartCallAnalyticsJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartCallAnalyticsJob", params, optFns, c.addOperationStartCallAnalyticsJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartCallAnalyticsJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartCallAnalyticsJobInput struct {

	// A unique name, chosen by you, for your Call Analytics job. This name is case
	// sensitive, cannot contain spaces, and must be unique within an Amazon Web
	// Services account. If you try to create a new job with the same name as an
	// existing job, you get a ConflictException error.
	//
	// This member is required.
	CallAnalyticsJobName *string

	// Describes the Amazon S3 location of the media file you want to use in your
	// request.
	//
	// This member is required.
	Media *types.Media

	// Allows you to specify which speaker is on which channel. For example, if your
	// agent is the first participant to speak, you would set ChannelId to 0 (to
	// indicate the first channel) and ParticipantRole to AGENT (to indicate that it's
	// the agent speaking).
	ChannelDefinitions []types.ChannelDefinition

	// The Amazon Resource Name (ARN) of an IAM role that has permissions to access the
	// Amazon S3 bucket that contains your input files. If the role you specify doesn’t
	// have the appropriate permissions to access the specified Amazon S3 location,
	// your request fails. IAM role ARNs have the format
	// arn:partition:iam::account:role/role-name-with-path. For example:
	// arn:aws:iam::111122223333:role/Admin. For more information, see IAM ARNs
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-arns).
	DataAccessRoleArn *string

	// The KMS key you want to use to encrypt your Call Analytics output. If using a
	// key located in the current Amazon Web Services account, you can specify your KMS
	// key in one of four ways:
	//
	// * Use the KMS key ID itself. For example,
	// 1234abcd-12ab-34cd-56ef-1234567890ab.
	//
	// * Use an alias for the KMS key ID. For
	// example, alias/ExampleAlias.
	//
	// * Use the Amazon Resource Name (ARN) for the KMS
	// key ID. For example,
	// arn:aws:kms:region:account-ID:key/1234abcd-12ab-34cd-56ef-1234567890ab.
	//
	// * Use
	// the ARN for the KMS key alias. For example,
	// arn:aws:kms:region:account-ID:alias/ExampleAlias.
	//
	// If using a key located in a
	// different Amazon Web Services account than the current Amazon Web Services
	// account, you can specify your KMS key in one of two ways:
	//
	// * Use the ARN for the
	// KMS key ID. For example,
	// arn:aws:kms:region:account-ID:key/1234abcd-12ab-34cd-56ef-1234567890ab.
	//
	// * Use
	// the ARN for the KMS key alias. For example,
	// arn:aws:kms:region:account-ID:alias/ExampleAlias.
	//
	// If you don't specify an
	// encryption key, your output is encrypted with the default Amazon S3 key
	// (SSE-S3). If you specify a KMS key to encrypt your output, you must also specify
	// an output location using the OutputLocation parameter. Note that the user making
	// the request must have permission to use the specified KMS key.
	OutputEncryptionKMSKeyId *string

	// The Amazon S3 location where you want your Call Analytics transcription output
	// stored. You can use any of the following formats to specify the output
	// location:
	//
	// * s3://DOC-EXAMPLE-BUCKET
	//
	// *
	// s3://DOC-EXAMPLE-BUCKET/my-output-folder/
	//
	// *
	// s3://DOC-EXAMPLE-BUCKET/my-output-folder/my-call-analytics-job.json
	//
	// Unless you
	// specify a file name (option 3), the name of your output file has a default value
	// that matches the name you specified for your transcription job using the
	// CallAnalyticsJobName parameter. You can specify a KMS key to encrypt your output
	// using the OutputEncryptionKMSKeyId parameter. If you don't specify a KMS key,
	// Amazon Transcribe uses the default Amazon S3 key for server-side encryption. If
	// you don't specify OutputLocation, your transcript is placed in a service-managed
	// Amazon S3 bucket and you are provided with a URI to access your transcript.
	OutputLocation *string

	// Specify additional optional settings in your request, including content
	// redaction; allows you to apply custom language models, vocabulary filters, and
	// custom vocabularies to your Call Analytics job.
	Settings *types.CallAnalyticsJobSettings

	noSmithyDocumentSerde
}

type StartCallAnalyticsJobOutput struct {

	// Provides detailed information about the current Call Analytics job, including
	// job status and, if applicable, failure reason.
	CallAnalyticsJob *types.CallAnalyticsJob

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartCallAnalyticsJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartCallAnalyticsJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartCallAnalyticsJob{}, middleware.After)
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
	if err = addOpStartCallAnalyticsJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartCallAnalyticsJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartCallAnalyticsJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "StartCallAnalyticsJob",
	}
}
