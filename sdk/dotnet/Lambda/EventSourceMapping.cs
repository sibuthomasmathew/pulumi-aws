// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Lambda
{
    /// <summary>
    /// Provides a Lambda event source mapping. This allows Lambda functions to get events from Kinesis, DynamoDB, SQS and Managed Streaming for Apache Kafka (MSK).
    /// 
    /// For information about Lambda and how to use it, see [What is AWS Lambda?](http://docs.aws.amazon.com/lambda/latest/dg/welcome.html).
    /// For information about event source mappings, see [CreateEventSourceMapping](http://docs.aws.amazon.com/lambda/latest/dg/API_CreateEventSourceMapping.html) in the API docs.
    /// 
    /// ## Example Usage
    /// ### DynamoDB
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.Lambda.EventSourceMapping("example", new Aws.Lambda.EventSourceMappingArgs
    ///         {
    ///             EventSourceArn = aws_dynamodb_table.Example.Stream_arn,
    ///             FunctionName = aws_lambda_function.Example.Arn,
    ///             StartingPosition = "LATEST",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Kinesis
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.Lambda.EventSourceMapping("example", new Aws.Lambda.EventSourceMappingArgs
    ///         {
    ///             EventSourceArn = aws_kinesis_stream.Example.Arn,
    ///             FunctionName = aws_lambda_function.Example.Arn,
    ///             StartingPosition = "LATEST",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Managed Streaming for Apache Kafka (MSK)
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.Lambda.EventSourceMapping("example", new Aws.Lambda.EventSourceMappingArgs
    ///         {
    ///             EventSourceArn = aws_msk_cluster.Example.Arn,
    ///             FunctionName = aws_lambda_function.Example.Arn,
    ///             Topics = 
    ///             {
    ///                 "Example",
    ///             },
    ///             StartingPosition = "TRIM_HORIZON",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### SQS
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.Lambda.EventSourceMapping("example", new Aws.Lambda.EventSourceMappingArgs
    ///         {
    ///             EventSourceArn = aws_sqs_queue.Sqs_queue_test.Arn,
    ///             FunctionName = aws_lambda_function.Example.Arn,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Managed Streaming for Kafka (MSK)
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.Lambda.EventSourceMapping("example", new Aws.Lambda.EventSourceMappingArgs
    ///         {
    ///             EventSourceArn = aws_msk_cluster.Example.Arn,
    ///             FunctionName = aws_lambda_function.Example.Arn,
    ///             Topics = 
    ///             {
    ///                 "Example",
    ///             },
    ///             StartingPosition = "TRIM_HORIZON",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Lambda event source mappings can be imported using the `UUID` (event source mapping identifier), e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:lambda/eventSourceMapping:EventSourceMapping event_source_mapping 12345kxodurf3443
    /// ```
    /// 
    ///  [3]https://docs.aws.amazon.com/lambda/latest/dg/API_GetEventSourceMapping.html
    /// </summary>
    [AwsResourceType("aws:lambda/eventSourceMapping:EventSourceMapping")]
    public partial class EventSourceMapping : Pulumi.CustomResource
    {
        /// <summary>
        /// The largest number of records that Lambda will retrieve from your event source at the time of invocation. Defaults to `100` for DynamoDB, Kinesis and MSK, `10` for SQS.
        /// </summary>
        [Output("batchSize")]
        public Output<int?> BatchSize { get; private set; } = null!;

        [Output("bisectBatchOnFunctionError")]
        public Output<bool?> BisectBatchOnFunctionError { get; private set; } = null!;

        [Output("destinationConfig")]
        public Output<Outputs.EventSourceMappingDestinationConfig?> DestinationConfig { get; private set; } = null!;

        /// <summary>
        /// Determines if the mapping will be enabled on creation. Defaults to `true`.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// The event source ARN - can be a Kinesis stream, DynamoDB stream, SQS queue or MSK cluster.
        /// </summary>
        [Output("eventSourceArn")]
        public Output<string> EventSourceArn { get; private set; } = null!;

        /// <summary>
        /// The the ARN of the Lambda function the event source mapping is sending events to. (Note: this is a computed value that differs from `function_name` above.)
        /// </summary>
        [Output("functionArn")]
        public Output<string> FunctionArn { get; private set; } = null!;

        /// <summary>
        /// The name or the ARN of the Lambda function that will be subscribing to events.
        /// </summary>
        [Output("functionName")]
        public Output<string> FunctionName { get; private set; } = null!;

        /// <summary>
        /// The date this resource was last modified.
        /// </summary>
        [Output("lastModified")]
        public Output<string> LastModified { get; private set; } = null!;

        /// <summary>
        /// The result of the last AWS Lambda invocation of your Lambda function.
        /// </summary>
        [Output("lastProcessingResult")]
        public Output<string> LastProcessingResult { get; private set; } = null!;

        /// <summary>
        /// The maximum amount of time to gather records before invoking the function, in seconds (between 0 and 300). Records will continue to buffer (or accumulate in the case of an SQS queue event source) until either `maximum_batching_window_in_seconds` expires or `batch_size` has been met. For streaming event sources, defaults to as soon as records are available in the stream. If the batch it reads from the stream/queue only has one record in it, Lambda only sends one record to the function. Only available for stream sources (DynamoDB and Kinesis) and SQS standard queues.
        /// </summary>
        [Output("maximumBatchingWindowInSeconds")]
        public Output<int?> MaximumBatchingWindowInSeconds { get; private set; } = null!;

        [Output("maximumRecordAgeInSeconds")]
        public Output<int> MaximumRecordAgeInSeconds { get; private set; } = null!;

        [Output("maximumRetryAttempts")]
        public Output<int> MaximumRetryAttempts { get; private set; } = null!;

        [Output("parallelizationFactor")]
        public Output<int> ParallelizationFactor { get; private set; } = null!;

        /// <summary>
        /// The position in the stream where AWS Lambda should start reading. Must be one of `AT_TIMESTAMP` (Kinesis only), `LATEST` or `TRIM_HORIZON` if getting events from Kinesis, DynamoDB or MSK. Must not be provided if getting events from SQS. More information about these positions can be found in the [AWS DynamoDB Streams API Reference](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_streams_GetShardIterator.html) and [AWS Kinesis API Reference](https://docs.aws.amazon.com/kinesis/latest/APIReference/API_GetShardIterator.html#Kinesis-GetShardIterator-request-ShardIteratorType).
        /// </summary>
        [Output("startingPosition")]
        public Output<string?> StartingPosition { get; private set; } = null!;

        /// <summary>
        /// A timestamp in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) of the data record which to start reading when using `starting_position` set to `AT_TIMESTAMP`. If a record with this exact timestamp does not exist, the next later record is chosen. If the timestamp is older than the current trim horizon, the oldest available record is chosen.
        /// * `parallelization_factor`: - (Optional) The number of batches to process from each shard concurrently. Only available for stream sources (DynamoDB and Kinesis). Minimum and default of 1, maximum of 10.
        /// * `maximum_retry_attempts`: - (Optional) The maximum number of times to retry when the function returns an error. Only available for stream sources (DynamoDB and Kinesis). Minimum of 0, maximum and default of 10000.
        /// * `maximum_record_age_in_seconds`: - (Optional) The maximum age of a record that Lambda sends to a function for processing. Only available for stream sources (DynamoDB and Kinesis). Minimum of 60, maximum and default of 604800.
        /// * `bisect_batch_on_function_error`: - (Optional) If the function returns an error, split the batch in two and retry. Only available for stream sources (DynamoDB and Kinesis). Defaults to `false`.
        /// </summary>
        [Output("startingPositionTimestamp")]
        public Output<string?> StartingPositionTimestamp { get; private set; } = null!;

        /// <summary>
        /// The state of the event source mapping.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The reason the event source mapping is in its current state.
        /// </summary>
        [Output("stateTransitionReason")]
        public Output<string> StateTransitionReason { get; private set; } = null!;

        /// <summary>
        /// The name of the Kafka topics. Only available for MSK sources. A single topic name must be specified.
        /// * `destination_config`: - (Optional) An Amazon SQS queue or Amazon SNS topic destination for failed records. Only available for stream sources (DynamoDB and Kinesis). Detailed below.
        /// </summary>
        [Output("topics")]
        public Output<ImmutableArray<string>> Topics { get; private set; } = null!;

        /// <summary>
        /// The UUID of the created event source mapping.
        /// </summary>
        [Output("uuid")]
        public Output<string> Uuid { get; private set; } = null!;


        /// <summary>
        /// Create a EventSourceMapping resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EventSourceMapping(string name, EventSourceMappingArgs args, CustomResourceOptions? options = null)
            : base("aws:lambda/eventSourceMapping:EventSourceMapping", name, args ?? new EventSourceMappingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EventSourceMapping(string name, Input<string> id, EventSourceMappingState? state = null, CustomResourceOptions? options = null)
            : base("aws:lambda/eventSourceMapping:EventSourceMapping", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing EventSourceMapping resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EventSourceMapping Get(string name, Input<string> id, EventSourceMappingState? state = null, CustomResourceOptions? options = null)
        {
            return new EventSourceMapping(name, id, state, options);
        }
    }

    public sealed class EventSourceMappingArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The largest number of records that Lambda will retrieve from your event source at the time of invocation. Defaults to `100` for DynamoDB, Kinesis and MSK, `10` for SQS.
        /// </summary>
        [Input("batchSize")]
        public Input<int>? BatchSize { get; set; }

        [Input("bisectBatchOnFunctionError")]
        public Input<bool>? BisectBatchOnFunctionError { get; set; }

        [Input("destinationConfig")]
        public Input<Inputs.EventSourceMappingDestinationConfigArgs>? DestinationConfig { get; set; }

        /// <summary>
        /// Determines if the mapping will be enabled on creation. Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The event source ARN - can be a Kinesis stream, DynamoDB stream, SQS queue or MSK cluster.
        /// </summary>
        [Input("eventSourceArn", required: true)]
        public Input<string> EventSourceArn { get; set; } = null!;

        /// <summary>
        /// The name or the ARN of the Lambda function that will be subscribing to events.
        /// </summary>
        [Input("functionName", required: true)]
        public Input<string> FunctionName { get; set; } = null!;

        /// <summary>
        /// The maximum amount of time to gather records before invoking the function, in seconds (between 0 and 300). Records will continue to buffer (or accumulate in the case of an SQS queue event source) until either `maximum_batching_window_in_seconds` expires or `batch_size` has been met. For streaming event sources, defaults to as soon as records are available in the stream. If the batch it reads from the stream/queue only has one record in it, Lambda only sends one record to the function. Only available for stream sources (DynamoDB and Kinesis) and SQS standard queues.
        /// </summary>
        [Input("maximumBatchingWindowInSeconds")]
        public Input<int>? MaximumBatchingWindowInSeconds { get; set; }

        [Input("maximumRecordAgeInSeconds")]
        public Input<int>? MaximumRecordAgeInSeconds { get; set; }

        [Input("maximumRetryAttempts")]
        public Input<int>? MaximumRetryAttempts { get; set; }

        [Input("parallelizationFactor")]
        public Input<int>? ParallelizationFactor { get; set; }

        /// <summary>
        /// The position in the stream where AWS Lambda should start reading. Must be one of `AT_TIMESTAMP` (Kinesis only), `LATEST` or `TRIM_HORIZON` if getting events from Kinesis, DynamoDB or MSK. Must not be provided if getting events from SQS. More information about these positions can be found in the [AWS DynamoDB Streams API Reference](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_streams_GetShardIterator.html) and [AWS Kinesis API Reference](https://docs.aws.amazon.com/kinesis/latest/APIReference/API_GetShardIterator.html#Kinesis-GetShardIterator-request-ShardIteratorType).
        /// </summary>
        [Input("startingPosition")]
        public Input<string>? StartingPosition { get; set; }

        /// <summary>
        /// A timestamp in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) of the data record which to start reading when using `starting_position` set to `AT_TIMESTAMP`. If a record with this exact timestamp does not exist, the next later record is chosen. If the timestamp is older than the current trim horizon, the oldest available record is chosen.
        /// * `parallelization_factor`: - (Optional) The number of batches to process from each shard concurrently. Only available for stream sources (DynamoDB and Kinesis). Minimum and default of 1, maximum of 10.
        /// * `maximum_retry_attempts`: - (Optional) The maximum number of times to retry when the function returns an error. Only available for stream sources (DynamoDB and Kinesis). Minimum of 0, maximum and default of 10000.
        /// * `maximum_record_age_in_seconds`: - (Optional) The maximum age of a record that Lambda sends to a function for processing. Only available for stream sources (DynamoDB and Kinesis). Minimum of 60, maximum and default of 604800.
        /// * `bisect_batch_on_function_error`: - (Optional) If the function returns an error, split the batch in two and retry. Only available for stream sources (DynamoDB and Kinesis). Defaults to `false`.
        /// </summary>
        [Input("startingPositionTimestamp")]
        public Input<string>? StartingPositionTimestamp { get; set; }

        [Input("topics")]
        private InputList<string>? _topics;

        /// <summary>
        /// The name of the Kafka topics. Only available for MSK sources. A single topic name must be specified.
        /// * `destination_config`: - (Optional) An Amazon SQS queue or Amazon SNS topic destination for failed records. Only available for stream sources (DynamoDB and Kinesis). Detailed below.
        /// </summary>
        public InputList<string> Topics
        {
            get => _topics ?? (_topics = new InputList<string>());
            set => _topics = value;
        }

        public EventSourceMappingArgs()
        {
        }
    }

    public sealed class EventSourceMappingState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The largest number of records that Lambda will retrieve from your event source at the time of invocation. Defaults to `100` for DynamoDB, Kinesis and MSK, `10` for SQS.
        /// </summary>
        [Input("batchSize")]
        public Input<int>? BatchSize { get; set; }

        [Input("bisectBatchOnFunctionError")]
        public Input<bool>? BisectBatchOnFunctionError { get; set; }

        [Input("destinationConfig")]
        public Input<Inputs.EventSourceMappingDestinationConfigGetArgs>? DestinationConfig { get; set; }

        /// <summary>
        /// Determines if the mapping will be enabled on creation. Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The event source ARN - can be a Kinesis stream, DynamoDB stream, SQS queue or MSK cluster.
        /// </summary>
        [Input("eventSourceArn")]
        public Input<string>? EventSourceArn { get; set; }

        /// <summary>
        /// The the ARN of the Lambda function the event source mapping is sending events to. (Note: this is a computed value that differs from `function_name` above.)
        /// </summary>
        [Input("functionArn")]
        public Input<string>? FunctionArn { get; set; }

        /// <summary>
        /// The name or the ARN of the Lambda function that will be subscribing to events.
        /// </summary>
        [Input("functionName")]
        public Input<string>? FunctionName { get; set; }

        /// <summary>
        /// The date this resource was last modified.
        /// </summary>
        [Input("lastModified")]
        public Input<string>? LastModified { get; set; }

        /// <summary>
        /// The result of the last AWS Lambda invocation of your Lambda function.
        /// </summary>
        [Input("lastProcessingResult")]
        public Input<string>? LastProcessingResult { get; set; }

        /// <summary>
        /// The maximum amount of time to gather records before invoking the function, in seconds (between 0 and 300). Records will continue to buffer (or accumulate in the case of an SQS queue event source) until either `maximum_batching_window_in_seconds` expires or `batch_size` has been met. For streaming event sources, defaults to as soon as records are available in the stream. If the batch it reads from the stream/queue only has one record in it, Lambda only sends one record to the function. Only available for stream sources (DynamoDB and Kinesis) and SQS standard queues.
        /// </summary>
        [Input("maximumBatchingWindowInSeconds")]
        public Input<int>? MaximumBatchingWindowInSeconds { get; set; }

        [Input("maximumRecordAgeInSeconds")]
        public Input<int>? MaximumRecordAgeInSeconds { get; set; }

        [Input("maximumRetryAttempts")]
        public Input<int>? MaximumRetryAttempts { get; set; }

        [Input("parallelizationFactor")]
        public Input<int>? ParallelizationFactor { get; set; }

        /// <summary>
        /// The position in the stream where AWS Lambda should start reading. Must be one of `AT_TIMESTAMP` (Kinesis only), `LATEST` or `TRIM_HORIZON` if getting events from Kinesis, DynamoDB or MSK. Must not be provided if getting events from SQS. More information about these positions can be found in the [AWS DynamoDB Streams API Reference](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_streams_GetShardIterator.html) and [AWS Kinesis API Reference](https://docs.aws.amazon.com/kinesis/latest/APIReference/API_GetShardIterator.html#Kinesis-GetShardIterator-request-ShardIteratorType).
        /// </summary>
        [Input("startingPosition")]
        public Input<string>? StartingPosition { get; set; }

        /// <summary>
        /// A timestamp in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) of the data record which to start reading when using `starting_position` set to `AT_TIMESTAMP`. If a record with this exact timestamp does not exist, the next later record is chosen. If the timestamp is older than the current trim horizon, the oldest available record is chosen.
        /// * `parallelization_factor`: - (Optional) The number of batches to process from each shard concurrently. Only available for stream sources (DynamoDB and Kinesis). Minimum and default of 1, maximum of 10.
        /// * `maximum_retry_attempts`: - (Optional) The maximum number of times to retry when the function returns an error. Only available for stream sources (DynamoDB and Kinesis). Minimum of 0, maximum and default of 10000.
        /// * `maximum_record_age_in_seconds`: - (Optional) The maximum age of a record that Lambda sends to a function for processing. Only available for stream sources (DynamoDB and Kinesis). Minimum of 60, maximum and default of 604800.
        /// * `bisect_batch_on_function_error`: - (Optional) If the function returns an error, split the batch in two and retry. Only available for stream sources (DynamoDB and Kinesis). Defaults to `false`.
        /// </summary>
        [Input("startingPositionTimestamp")]
        public Input<string>? StartingPositionTimestamp { get; set; }

        /// <summary>
        /// The state of the event source mapping.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The reason the event source mapping is in its current state.
        /// </summary>
        [Input("stateTransitionReason")]
        public Input<string>? StateTransitionReason { get; set; }

        [Input("topics")]
        private InputList<string>? _topics;

        /// <summary>
        /// The name of the Kafka topics. Only available for MSK sources. A single topic name must be specified.
        /// * `destination_config`: - (Optional) An Amazon SQS queue or Amazon SNS topic destination for failed records. Only available for stream sources (DynamoDB and Kinesis). Detailed below.
        /// </summary>
        public InputList<string> Topics
        {
            get => _topics ?? (_topics = new InputList<string>());
            set => _topics = value;
        }

        /// <summary>
        /// The UUID of the created event source mapping.
        /// </summary>
        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        public EventSourceMappingState()
        {
        }
    }
}
