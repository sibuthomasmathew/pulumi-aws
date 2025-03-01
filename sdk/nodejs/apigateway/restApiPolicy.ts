// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an API Gateway REST API Policy.
 *
 * > **Note:** Amazon API Gateway Version 1 resources are used for creating and deploying REST APIs. To create and deploy WebSocket and HTTP APIs, use Amazon API Gateway Version 2 resources.
 *
 * ## Example Usage
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const testRestApi = new aws.apigateway.RestApi("testRestApi", {});
 * const testRestApiPolicy = new aws.apigateway.RestApiPolicy("testRestApiPolicy", {
 *     restApiId: testRestApi.id,
 *     policy: pulumi.interpolate`{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Effect": "Allow",
 *       "Principal": {
 *         "AWS": "*"
 *       },
 *       "Action": "execute-api:Invoke",
 *       "Resource": "${testRestApi.executionArn}",
 *       "Condition": {
 *         "IpAddress": {
 *           "aws:SourceIp": "123.123.123.123/32"
 *         }
 *       }
 *     }
 *   ]
 * }
 * `,
 * });
 * ```
 *
 * ## Import
 *
 * `aws_api_gateway_rest_api_policy` can be imported by using the REST API ID, e.g.
 *
 * ```sh
 *  $ pulumi import aws:apigateway/restApiPolicy:RestApiPolicy example 12345abcde
 * ```
 */
export class RestApiPolicy extends pulumi.CustomResource {
    /**
     * Get an existing RestApiPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RestApiPolicyState, opts?: pulumi.CustomResourceOptions): RestApiPolicy {
        return new RestApiPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:apigateway/restApiPolicy:RestApiPolicy';

    /**
     * Returns true if the given object is an instance of RestApiPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RestApiPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RestApiPolicy.__pulumiType;
    }

    /**
     * JSON formatted policy document that controls access to the API Gateway.
     */
    public readonly policy!: pulumi.Output<string>;
    /**
     * The ID of the REST API.
     */
    public readonly restApiId!: pulumi.Output<string>;

    /**
     * Create a RestApiPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RestApiPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RestApiPolicyArgs | RestApiPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RestApiPolicyState | undefined;
            inputs["policy"] = state ? state.policy : undefined;
            inputs["restApiId"] = state ? state.restApiId : undefined;
        } else {
            const args = argsOrState as RestApiPolicyArgs | undefined;
            if ((!args || args.policy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policy'");
            }
            if ((!args || args.restApiId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'restApiId'");
            }
            inputs["policy"] = args ? args.policy : undefined;
            inputs["restApiId"] = args ? args.restApiId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(RestApiPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RestApiPolicy resources.
 */
export interface RestApiPolicyState {
    /**
     * JSON formatted policy document that controls access to the API Gateway.
     */
    readonly policy?: pulumi.Input<string>;
    /**
     * The ID of the REST API.
     */
    readonly restApiId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RestApiPolicy resource.
 */
export interface RestApiPolicyArgs {
    /**
     * JSON formatted policy document that controls access to the API Gateway.
     */
    readonly policy: pulumi.Input<string>;
    /**
     * The ID of the REST API.
     */
    readonly restApiId: pulumi.Input<string>;
}
