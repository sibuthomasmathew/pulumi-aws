// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

import {RestApi} from "./index";

/**
 * Provides a Model for a REST API Gateway.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const myDemoAPI = new aws.apigateway.RestApi("myDemoAPI", {description: "This is my API for demonstration purposes"});
 * const myDemoModel = new aws.apigateway.Model("myDemoModel", {
 *     restApi: myDemoAPI.id,
 *     description: "a JSON schema",
 *     contentType: "application/json",
 *     schema: `{
 *   "type": "object"
 * }
 * `,
 * });
 * ```
 *
 * ## Import
 *
 * `aws_api_gateway_model` can be imported using `REST-API-ID/NAME`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:apigateway/model:Model example 12345abcde/example
 * ```
 */
export class Model extends pulumi.CustomResource {
    /**
     * Get an existing Model resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ModelState, opts?: pulumi.CustomResourceOptions): Model {
        return new Model(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:apigateway/model:Model';

    /**
     * Returns true if the given object is an instance of Model.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Model {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Model.__pulumiType;
    }

    /**
     * The content type of the model
     */
    public readonly contentType!: pulumi.Output<string>;
    /**
     * The description of the model
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the model
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the associated REST API
     */
    public readonly restApi!: pulumi.Output<string>;
    /**
     * The schema of the model in a JSON form
     */
    public readonly schema!: pulumi.Output<string | undefined>;

    /**
     * Create a Model resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ModelArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ModelArgs | ModelState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ModelState | undefined;
            inputs["contentType"] = state ? state.contentType : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["restApi"] = state ? state.restApi : undefined;
            inputs["schema"] = state ? state.schema : undefined;
        } else {
            const args = argsOrState as ModelArgs | undefined;
            if ((!args || args.contentType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'contentType'");
            }
            if ((!args || args.restApi === undefined) && !opts.urn) {
                throw new Error("Missing required property 'restApi'");
            }
            inputs["contentType"] = args ? args.contentType : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["restApi"] = args ? args.restApi : undefined;
            inputs["schema"] = args ? args.schema : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Model.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Model resources.
 */
export interface ModelState {
    /**
     * The content type of the model
     */
    readonly contentType?: pulumi.Input<string>;
    /**
     * The description of the model
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the model
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the associated REST API
     */
    readonly restApi?: pulumi.Input<string | RestApi>;
    /**
     * The schema of the model in a JSON form
     */
    readonly schema?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Model resource.
 */
export interface ModelArgs {
    /**
     * The content type of the model
     */
    readonly contentType: pulumi.Input<string>;
    /**
     * The description of the model
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the model
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the associated REST API
     */
    readonly restApi: pulumi.Input<string | RestApi>;
    /**
     * The schema of the model in a JSON form
     */
    readonly schema?: pulumi.Input<string>;
}
