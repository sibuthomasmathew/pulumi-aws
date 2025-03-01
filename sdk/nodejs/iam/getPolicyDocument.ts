// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Generates an IAM policy document in JSON format for use with resources that expect policy documents such as `aws.iam.Policy`.
 *
 * Using this data source to generate policy documents is *optional*. It is also valid to use literal JSON strings in your configuration or to use the `file` interpolation function to read a raw JSON policy document from a file.
 *
 * ## Example Usage
 * ### Example with Both Source and Override Documents
 *
 * You can also combine `sourceJson` and `overrideJson` in the same document.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const source = aws.iam.getPolicyDocument({
 *     statements: [{
 *         sid: "OverridePlaceholder",
 *         actions: ["ec2:DescribeAccountAttributes"],
 *         resources: ["*"],
 *     }],
 * });
 * const override = aws.iam.getPolicyDocument({
 *     statements: [{
 *         sid: "OverridePlaceholder",
 *         actions: ["s3:GetObject"],
 *         resources: ["*"],
 *     }],
 * });
 * const politik = Promise.all([source, override]).then(([source, override]) => aws.iam.getPolicyDocument({
 *     sourceJson: source.json,
 *     overrideJson: override.json,
 * }));
 * ```
 *
 * `data.aws_iam_policy_document.politik.json` will evaluate to:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * ```
 */
export function getPolicyDocument(args?: GetPolicyDocumentArgs, opts?: pulumi.InvokeOptions): Promise<GetPolicyDocumentResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:iam/getPolicyDocument:getPolicyDocument", {
        "overrideJson": args.overrideJson,
        "overridePolicyDocuments": args.overridePolicyDocuments,
        "policyId": args.policyId,
        "sourceJson": args.sourceJson,
        "sourcePolicyDocuments": args.sourcePolicyDocuments,
        "statements": args.statements,
        "version": args.version,
    }, opts);
}

/**
 * A collection of arguments for invoking getPolicyDocument.
 */
export interface GetPolicyDocumentArgs {
    /**
     * IAM policy document whose statements with non-blank `sid`s will override statements with the same `sid` from documents assigned to the `sourceJson`, `sourcePolicyDocuments`, and `overridePolicyDocuments` arguments. Non-overriding statements will be added to the exported document.
     */
    readonly overrideJson?: string;
    /**
     * List of IAM policy documents that are merged together into the exported document. In merging, statements with non-blank `sid`s will override statements with the same `sid` from earlier documents in the list. Statements with non-blank `sid`s will also override statements with the same `sid` from documents provided in the `sourceJson` and `sourcePolicyDocuments` arguments.  Non-overriding statements will be added to the exported document.
     */
    readonly overridePolicyDocuments?: string[];
    /**
     * ID for the policy document.
     */
    readonly policyId?: string;
    /**
     * IAM policy document used as a base for the exported policy document. Statements with the same `sid` from documents assigned to the `overrideJson` and `overridePolicyDocuments` arguments will override source statements.
     */
    readonly sourceJson?: string;
    /**
     * List of IAM policy documents that are merged together into the exported document. Statements defined in `sourcePolicyDocuments` or `sourceJson` must have unique `sid`s. Statements with the same `sid` from documents assigned to the `overrideJson` and `overridePolicyDocuments` arguments will override source statements.
     */
    readonly sourcePolicyDocuments?: string[];
    /**
     * Configuration block for a policy statement. Detailed below.
     */
    readonly statements?: inputs.iam.GetPolicyDocumentStatement[];
    /**
     * IAM policy document version. Valid values are `2008-10-17` and `2012-10-17`. Defaults to `2012-10-17`. For more information, see the [AWS IAM User Guide](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_version.html).
     */
    readonly version?: string;
}

/**
 * A collection of values returned by getPolicyDocument.
 */
export interface GetPolicyDocumentResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Standard JSON policy document rendered based on the arguments above.
     */
    readonly json: string;
    readonly overrideJson?: string;
    readonly overridePolicyDocuments?: string[];
    readonly policyId?: string;
    readonly sourceJson?: string;
    readonly sourcePolicyDocuments?: string[];
    readonly statements?: outputs.iam.GetPolicyDocumentStatement[];
    readonly version?: string;
}
