// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a WAF Rule Resource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const ipset = new aws.waf.IpSet("ipset", {ipSetDescriptors: [{
 *     type: "IPV4",
 *     value: "192.0.7.0/24",
 * }]});
 * const wafrule = new aws.waf.Rule("wafrule", {
 *     metricName: "tfWAFRule",
 *     predicates: [{
 *         dataId: ipset.id,
 *         negated: false,
 *         type: "IPMatch",
 *     }],
 * }, {
 *     dependsOn: [ipset],
 * });
 * ```
 *
 * ## Import
 *
 * WAF rules can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import aws:waf/rule:Rule example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
 * ```
 */
export class Rule extends pulumi.CustomResource {
    /**
     * Get an existing Rule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RuleState, opts?: pulumi.CustomResourceOptions): Rule {
        return new Rule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:waf/rule:Rule';

    /**
     * Returns true if the given object is an instance of Rule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Rule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Rule.__pulumiType;
    }

    /**
     * The ARN of the WAF rule.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The name or description for the Amazon CloudWatch metric of this rule. The name can contain only alphanumeric characters (A-Z, a-z, 0-9); the name can't contain whitespace.
     */
    public readonly metricName!: pulumi.Output<string>;
    /**
     * The name or description of the rule.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The objects to include in a rule (documented below).
     */
    public readonly predicates!: pulumi.Output<outputs.waf.RulePredicate[] | undefined>;
    /**
     * Key-value map of resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Rule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RuleArgs | RuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RuleState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["metricName"] = state ? state.metricName : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["predicates"] = state ? state.predicates : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as RuleArgs | undefined;
            if ((!args || args.metricName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'metricName'");
            }
            inputs["metricName"] = args ? args.metricName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["predicates"] = args ? args.predicates : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Rule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Rule resources.
 */
export interface RuleState {
    /**
     * The ARN of the WAF rule.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The name or description for the Amazon CloudWatch metric of this rule. The name can contain only alphanumeric characters (A-Z, a-z, 0-9); the name can't contain whitespace.
     */
    readonly metricName?: pulumi.Input<string>;
    /**
     * The name or description of the rule.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The objects to include in a rule (documented below).
     */
    readonly predicates?: pulumi.Input<pulumi.Input<inputs.waf.RulePredicate>[]>;
    /**
     * Key-value map of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Rule resource.
 */
export interface RuleArgs {
    /**
     * The name or description for the Amazon CloudWatch metric of this rule. The name can contain only alphanumeric characters (A-Z, a-z, 0-9); the name can't contain whitespace.
     */
    readonly metricName: pulumi.Input<string>;
    /**
     * The name or description of the rule.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The objects to include in a rule (documented below).
     */
    readonly predicates?: pulumi.Input<pulumi.Input<inputs.waf.RulePredicate>[]>;
    /**
     * Key-value map of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
