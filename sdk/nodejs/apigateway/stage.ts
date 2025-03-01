// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

import {Deployment, RestApi} from "./index";

/**
 * Manages an API Gateway Stage. A stage is a named reference to a deployment, which can be done via the `aws.apigateway.Deployment` resource. Stages can be optionally managed further with the `aws.apigateway.BasePathMapping` resource, `aws.apigateway.DomainName` resource, and `awsApiMethodSettings` resource. For more information, see the [API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-stages.html).
 *
 * ## Example Usage
 * ### Managing the API Logging CloudWatch Log Group
 *
 * API Gateway provides the ability to [enable CloudWatch API logging](https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-logging.html). To manage the CloudWatch Log Group when this feature is enabled, the `aws.cloudwatch.LogGroup` resource can be used where the name matches the API Gateway naming convention. If the CloudWatch Log Group previously exists, the `aws.cloudwatch.LogGroup` resource can be imported as a one time operation and recreation of the environment can occur without import.
 *
 * > The below configuration uses [`dependsOn`](https://www.pulumi.com/docs/intro/concepts/programming-model/#dependson) to prevent ordering issues with API Gateway automatically creating the log group first and a variable for naming consistency. Other ordering and naming methodologies may be more appropriate for your environment.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const config = new pulumi.Config();
 * const stageName = config.get("stageName") || "example";
 * const exampleRestApi = new aws.apigateway.RestApi("exampleRestApi", {});
 * // ... other configuration ...
 * const exampleLogGroup = new aws.cloudwatch.LogGroup("exampleLogGroup", {retentionInDays: 7});
 * // ... potentially other configuration ...
 * const exampleStage = new aws.apigateway.Stage("exampleStage", {stageName: stageName}, {
 *     dependsOn: [exampleLogGroup],
 * });
 * // ... other configuration ...
 * ```
 *
 * ## Import
 *
 * `aws_api_gateway_stage` can be imported using `REST-API-ID/STAGE-NAME`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:apigateway/stage:Stage example 12345abcde/example
 * ```
 */
export class Stage extends pulumi.CustomResource {
    /**
     * Get an existing Stage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StageState, opts?: pulumi.CustomResourceOptions): Stage {
        return new Stage(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:apigateway/stage:Stage';

    /**
     * Returns true if the given object is an instance of Stage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Stage {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Stage.__pulumiType;
    }

    /**
     * Enables access logs for the API stage. Detailed below.
     */
    public readonly accessLogSettings!: pulumi.Output<outputs.apigateway.StageAccessLogSettings | undefined>;
    /**
     * Amazon Resource Name (ARN)
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Specifies whether a cache cluster is enabled for the stage
     */
    public readonly cacheClusterEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The size of the cache cluster for the stage, if enabled. Allowed values include `0.5`, `1.6`, `6.1`, `13.5`, `28.4`, `58.2`, `118` and `237`.
     */
    public readonly cacheClusterSize!: pulumi.Output<string | undefined>;
    /**
     * The identifier of a client certificate for the stage.
     */
    public readonly clientCertificateId!: pulumi.Output<string | undefined>;
    /**
     * The ID of the deployment that the stage points to
     */
    public readonly deployment!: pulumi.Output<string>;
    /**
     * The description of the stage
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The version of the associated API documentation
     */
    public readonly documentationVersion!: pulumi.Output<string | undefined>;
    /**
     * The execution ARN to be used in `lambdaPermission`'s `sourceArn`
     * when allowing API Gateway to invoke a Lambda function,
     * e.g. `arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j/prod`
     */
    public /*out*/ readonly executionArn!: pulumi.Output<string>;
    /**
     * The URL to invoke the API pointing to the stage,
     * e.g. `https://z4675bid1j.execute-api.eu-west-2.amazonaws.com/prod`
     */
    public /*out*/ readonly invokeUrl!: pulumi.Output<string>;
    /**
     * The ID of the associated REST API
     */
    public readonly restApi!: pulumi.Output<string>;
    /**
     * The name of the stage
     */
    public readonly stageName!: pulumi.Output<string>;
    /**
     * A map of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map that defines the stage variables
     */
    public readonly variables!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Whether active tracing with X-ray is enabled. Defaults to `false`.
     */
    public readonly xrayTracingEnabled!: pulumi.Output<boolean | undefined>;

    /**
     * Create a Stage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StageArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StageArgs | StageState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StageState | undefined;
            inputs["accessLogSettings"] = state ? state.accessLogSettings : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["cacheClusterEnabled"] = state ? state.cacheClusterEnabled : undefined;
            inputs["cacheClusterSize"] = state ? state.cacheClusterSize : undefined;
            inputs["clientCertificateId"] = state ? state.clientCertificateId : undefined;
            inputs["deployment"] = state ? state.deployment : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["documentationVersion"] = state ? state.documentationVersion : undefined;
            inputs["executionArn"] = state ? state.executionArn : undefined;
            inputs["invokeUrl"] = state ? state.invokeUrl : undefined;
            inputs["restApi"] = state ? state.restApi : undefined;
            inputs["stageName"] = state ? state.stageName : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["variables"] = state ? state.variables : undefined;
            inputs["xrayTracingEnabled"] = state ? state.xrayTracingEnabled : undefined;
        } else {
            const args = argsOrState as StageArgs | undefined;
            if ((!args || args.deployment === undefined) && !opts.urn) {
                throw new Error("Missing required property 'deployment'");
            }
            if ((!args || args.restApi === undefined) && !opts.urn) {
                throw new Error("Missing required property 'restApi'");
            }
            if ((!args || args.stageName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'stageName'");
            }
            inputs["accessLogSettings"] = args ? args.accessLogSettings : undefined;
            inputs["cacheClusterEnabled"] = args ? args.cacheClusterEnabled : undefined;
            inputs["cacheClusterSize"] = args ? args.cacheClusterSize : undefined;
            inputs["clientCertificateId"] = args ? args.clientCertificateId : undefined;
            inputs["deployment"] = args ? args.deployment : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["documentationVersion"] = args ? args.documentationVersion : undefined;
            inputs["restApi"] = args ? args.restApi : undefined;
            inputs["stageName"] = args ? args.stageName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["variables"] = args ? args.variables : undefined;
            inputs["xrayTracingEnabled"] = args ? args.xrayTracingEnabled : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["executionArn"] = undefined /*out*/;
            inputs["invokeUrl"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Stage.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Stage resources.
 */
export interface StageState {
    /**
     * Enables access logs for the API stage. Detailed below.
     */
    readonly accessLogSettings?: pulumi.Input<inputs.apigateway.StageAccessLogSettings>;
    /**
     * Amazon Resource Name (ARN)
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * Specifies whether a cache cluster is enabled for the stage
     */
    readonly cacheClusterEnabled?: pulumi.Input<boolean>;
    /**
     * The size of the cache cluster for the stage, if enabled. Allowed values include `0.5`, `1.6`, `6.1`, `13.5`, `28.4`, `58.2`, `118` and `237`.
     */
    readonly cacheClusterSize?: pulumi.Input<string>;
    /**
     * The identifier of a client certificate for the stage.
     */
    readonly clientCertificateId?: pulumi.Input<string>;
    /**
     * The ID of the deployment that the stage points to
     */
    readonly deployment?: pulumi.Input<string | Deployment>;
    /**
     * The description of the stage
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The version of the associated API documentation
     */
    readonly documentationVersion?: pulumi.Input<string>;
    /**
     * The execution ARN to be used in `lambdaPermission`'s `sourceArn`
     * when allowing API Gateway to invoke a Lambda function,
     * e.g. `arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j/prod`
     */
    readonly executionArn?: pulumi.Input<string>;
    /**
     * The URL to invoke the API pointing to the stage,
     * e.g. `https://z4675bid1j.execute-api.eu-west-2.amazonaws.com/prod`
     */
    readonly invokeUrl?: pulumi.Input<string>;
    /**
     * The ID of the associated REST API
     */
    readonly restApi?: pulumi.Input<string | RestApi>;
    /**
     * The name of the stage
     */
    readonly stageName?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map that defines the stage variables
     */
    readonly variables?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Whether active tracing with X-ray is enabled. Defaults to `false`.
     */
    readonly xrayTracingEnabled?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a Stage resource.
 */
export interface StageArgs {
    /**
     * Enables access logs for the API stage. Detailed below.
     */
    readonly accessLogSettings?: pulumi.Input<inputs.apigateway.StageAccessLogSettings>;
    /**
     * Specifies whether a cache cluster is enabled for the stage
     */
    readonly cacheClusterEnabled?: pulumi.Input<boolean>;
    /**
     * The size of the cache cluster for the stage, if enabled. Allowed values include `0.5`, `1.6`, `6.1`, `13.5`, `28.4`, `58.2`, `118` and `237`.
     */
    readonly cacheClusterSize?: pulumi.Input<string>;
    /**
     * The identifier of a client certificate for the stage.
     */
    readonly clientCertificateId?: pulumi.Input<string>;
    /**
     * The ID of the deployment that the stage points to
     */
    readonly deployment: pulumi.Input<string | Deployment>;
    /**
     * The description of the stage
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The version of the associated API documentation
     */
    readonly documentationVersion?: pulumi.Input<string>;
    /**
     * The ID of the associated REST API
     */
    readonly restApi: pulumi.Input<string | RestApi>;
    /**
     * The name of the stage
     */
    readonly stageName: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map that defines the stage variables
     */
    readonly variables?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Whether active tracing with X-ray is enabled. Defaults to `false`.
     */
    readonly xrayTracingEnabled?: pulumi.Input<boolean>;
}
