// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * The "AMI copy" resource allows duplication of an Amazon Machine Image (AMI),
 * including cross-region copies.
 *
 * If the source AMI has associated EBS snapshots, those will also be duplicated
 * along with the AMI.
 *
 * This is useful for taking a single AMI provisioned in one region and making
 * it available in another for a multi-region deployment.
 *
 * Copying an AMI can take several minutes. The creation of this resource will
 * block until the new AMI is available for use on new instances.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.ec2.AmiCopy("example", {
 *     description: "A copy of ami-xxxxxxxx",
 *     sourceAmiId: "ami-xxxxxxxx",
 *     sourceAmiRegion: "us-west-1",
 *     tags: {
 *         Name: "HelloWorld",
 *     },
 * });
 * ```
 */
export class AmiCopy extends pulumi.CustomResource {
    /**
     * Get an existing AmiCopy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AmiCopyState, opts?: pulumi.CustomResourceOptions): AmiCopy {
        return new AmiCopy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2/amiCopy:AmiCopy';

    /**
     * Returns true if the given object is an instance of AmiCopy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AmiCopy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AmiCopy.__pulumiType;
    }

    /**
     * Machine architecture for created instances. Defaults to "x8664".
     */
    public /*out*/ readonly architecture!: pulumi.Output<string>;
    /**
     * The ARN of the AMI.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * A longer, human-readable description for the AMI.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Nested block describing an EBS block device that should be
     * attached to created instances. The structure of this block is described below.
     */
    public readonly ebsBlockDevices!: pulumi.Output<outputs.ec2.AmiCopyEbsBlockDevice[]>;
    /**
     * Specifies whether enhanced networking with ENA is enabled. Defaults to `false`.
     */
    public /*out*/ readonly enaSupport!: pulumi.Output<boolean>;
    /**
     * Boolean controlling whether the created EBS volumes will be encrypted. Can't be used with `snapshotId`.
     */
    public readonly encrypted!: pulumi.Output<boolean | undefined>;
    /**
     * Nested block describing an ephemeral block device that
     * should be attached to created instances. The structure of this block is described below.
     */
    public readonly ephemeralBlockDevices!: pulumi.Output<outputs.ec2.AmiCopyEphemeralBlockDevice[]>;
    public /*out*/ readonly hypervisor!: pulumi.Output<string>;
    /**
     * Path to an S3 object containing an image manifest, e.g. created
     * by the `ec2-upload-bundle` command in the EC2 command line tools.
     */
    public /*out*/ readonly imageLocation!: pulumi.Output<string>;
    public /*out*/ readonly imageOwnerAlias!: pulumi.Output<string>;
    public /*out*/ readonly imageType!: pulumi.Output<string>;
    /**
     * The id of the kernel image (AKI) that will be used as the paravirtual
     * kernel in created instances.
     */
    public /*out*/ readonly kernelId!: pulumi.Output<string>;
    /**
     * The full ARN of the AWS Key Management Service (AWS KMS) CMK to use when encrypting the snapshots of
     * an image during a copy operation. This parameter is only required if you want to use a non-default CMK;
     * if this parameter is not specified, the default CMK for EBS is used
     */
    public readonly kmsKeyId!: pulumi.Output<string>;
    public /*out*/ readonly manageEbsSnapshots!: pulumi.Output<boolean>;
    /**
     * A region-unique name for the AMI.
     */
    public readonly name!: pulumi.Output<string>;
    public /*out*/ readonly ownerId!: pulumi.Output<string>;
    public /*out*/ readonly platform!: pulumi.Output<string>;
    public /*out*/ readonly platformDetails!: pulumi.Output<string>;
    public /*out*/ readonly public!: pulumi.Output<boolean>;
    /**
     * The id of an initrd image (ARI) that will be used when booting the
     * created instances.
     */
    public /*out*/ readonly ramdiskId!: pulumi.Output<string>;
    /**
     * The name of the root device (for example, `/dev/sda1`, or `/dev/xvda`).
     */
    public /*out*/ readonly rootDeviceName!: pulumi.Output<string>;
    public /*out*/ readonly rootSnapshotId!: pulumi.Output<string>;
    /**
     * The id of the AMI to copy. This id must be valid in the region
     * given by `sourceAmiRegion`.
     */
    public readonly sourceAmiId!: pulumi.Output<string>;
    /**
     * The region from which the AMI will be copied. This may be the
     * same as the AWS provider region in order to create a copy within the same region.
     */
    public readonly sourceAmiRegion!: pulumi.Output<string>;
    /**
     * When set to "simple" (the default), enables enhanced networking
     * for created instances. No other value is supported at this time.
     */
    public /*out*/ readonly sriovNetSupport!: pulumi.Output<string>;
    /**
     * A map of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    public /*out*/ readonly usageOperation!: pulumi.Output<string>;
    /**
     * Keyword to choose what virtualization mode created instances
     * will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
     * changes the set of further arguments that are required, as described below.
     */
    public /*out*/ readonly virtualizationType!: pulumi.Output<string>;

    /**
     * Create a AmiCopy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AmiCopyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AmiCopyArgs | AmiCopyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AmiCopyState | undefined;
            inputs["architecture"] = state ? state.architecture : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["ebsBlockDevices"] = state ? state.ebsBlockDevices : undefined;
            inputs["enaSupport"] = state ? state.enaSupport : undefined;
            inputs["encrypted"] = state ? state.encrypted : undefined;
            inputs["ephemeralBlockDevices"] = state ? state.ephemeralBlockDevices : undefined;
            inputs["hypervisor"] = state ? state.hypervisor : undefined;
            inputs["imageLocation"] = state ? state.imageLocation : undefined;
            inputs["imageOwnerAlias"] = state ? state.imageOwnerAlias : undefined;
            inputs["imageType"] = state ? state.imageType : undefined;
            inputs["kernelId"] = state ? state.kernelId : undefined;
            inputs["kmsKeyId"] = state ? state.kmsKeyId : undefined;
            inputs["manageEbsSnapshots"] = state ? state.manageEbsSnapshots : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["ownerId"] = state ? state.ownerId : undefined;
            inputs["platform"] = state ? state.platform : undefined;
            inputs["platformDetails"] = state ? state.platformDetails : undefined;
            inputs["public"] = state ? state.public : undefined;
            inputs["ramdiskId"] = state ? state.ramdiskId : undefined;
            inputs["rootDeviceName"] = state ? state.rootDeviceName : undefined;
            inputs["rootSnapshotId"] = state ? state.rootSnapshotId : undefined;
            inputs["sourceAmiId"] = state ? state.sourceAmiId : undefined;
            inputs["sourceAmiRegion"] = state ? state.sourceAmiRegion : undefined;
            inputs["sriovNetSupport"] = state ? state.sriovNetSupport : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["usageOperation"] = state ? state.usageOperation : undefined;
            inputs["virtualizationType"] = state ? state.virtualizationType : undefined;
        } else {
            const args = argsOrState as AmiCopyArgs | undefined;
            if ((!args || args.sourceAmiId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceAmiId'");
            }
            if ((!args || args.sourceAmiRegion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceAmiRegion'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["ebsBlockDevices"] = args ? args.ebsBlockDevices : undefined;
            inputs["encrypted"] = args ? args.encrypted : undefined;
            inputs["ephemeralBlockDevices"] = args ? args.ephemeralBlockDevices : undefined;
            inputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["sourceAmiId"] = args ? args.sourceAmiId : undefined;
            inputs["sourceAmiRegion"] = args ? args.sourceAmiRegion : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["architecture"] = undefined /*out*/;
            inputs["arn"] = undefined /*out*/;
            inputs["enaSupport"] = undefined /*out*/;
            inputs["hypervisor"] = undefined /*out*/;
            inputs["imageLocation"] = undefined /*out*/;
            inputs["imageOwnerAlias"] = undefined /*out*/;
            inputs["imageType"] = undefined /*out*/;
            inputs["kernelId"] = undefined /*out*/;
            inputs["manageEbsSnapshots"] = undefined /*out*/;
            inputs["ownerId"] = undefined /*out*/;
            inputs["platform"] = undefined /*out*/;
            inputs["platformDetails"] = undefined /*out*/;
            inputs["public"] = undefined /*out*/;
            inputs["ramdiskId"] = undefined /*out*/;
            inputs["rootDeviceName"] = undefined /*out*/;
            inputs["rootSnapshotId"] = undefined /*out*/;
            inputs["sriovNetSupport"] = undefined /*out*/;
            inputs["usageOperation"] = undefined /*out*/;
            inputs["virtualizationType"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AmiCopy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AmiCopy resources.
 */
export interface AmiCopyState {
    /**
     * Machine architecture for created instances. Defaults to "x8664".
     */
    readonly architecture?: pulumi.Input<string>;
    /**
     * The ARN of the AMI.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * A longer, human-readable description for the AMI.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Nested block describing an EBS block device that should be
     * attached to created instances. The structure of this block is described below.
     */
    readonly ebsBlockDevices?: pulumi.Input<pulumi.Input<inputs.ec2.AmiCopyEbsBlockDevice>[]>;
    /**
     * Specifies whether enhanced networking with ENA is enabled. Defaults to `false`.
     */
    readonly enaSupport?: pulumi.Input<boolean>;
    /**
     * Boolean controlling whether the created EBS volumes will be encrypted. Can't be used with `snapshotId`.
     */
    readonly encrypted?: pulumi.Input<boolean>;
    /**
     * Nested block describing an ephemeral block device that
     * should be attached to created instances. The structure of this block is described below.
     */
    readonly ephemeralBlockDevices?: pulumi.Input<pulumi.Input<inputs.ec2.AmiCopyEphemeralBlockDevice>[]>;
    readonly hypervisor?: pulumi.Input<string>;
    /**
     * Path to an S3 object containing an image manifest, e.g. created
     * by the `ec2-upload-bundle` command in the EC2 command line tools.
     */
    readonly imageLocation?: pulumi.Input<string>;
    readonly imageOwnerAlias?: pulumi.Input<string>;
    readonly imageType?: pulumi.Input<string>;
    /**
     * The id of the kernel image (AKI) that will be used as the paravirtual
     * kernel in created instances.
     */
    readonly kernelId?: pulumi.Input<string>;
    /**
     * The full ARN of the AWS Key Management Service (AWS KMS) CMK to use when encrypting the snapshots of
     * an image during a copy operation. This parameter is only required if you want to use a non-default CMK;
     * if this parameter is not specified, the default CMK for EBS is used
     */
    readonly kmsKeyId?: pulumi.Input<string>;
    readonly manageEbsSnapshots?: pulumi.Input<boolean>;
    /**
     * A region-unique name for the AMI.
     */
    readonly name?: pulumi.Input<string>;
    readonly ownerId?: pulumi.Input<string>;
    readonly platform?: pulumi.Input<string>;
    readonly platformDetails?: pulumi.Input<string>;
    readonly public?: pulumi.Input<boolean>;
    /**
     * The id of an initrd image (ARI) that will be used when booting the
     * created instances.
     */
    readonly ramdiskId?: pulumi.Input<string>;
    /**
     * The name of the root device (for example, `/dev/sda1`, or `/dev/xvda`).
     */
    readonly rootDeviceName?: pulumi.Input<string>;
    readonly rootSnapshotId?: pulumi.Input<string>;
    /**
     * The id of the AMI to copy. This id must be valid in the region
     * given by `sourceAmiRegion`.
     */
    readonly sourceAmiId?: pulumi.Input<string>;
    /**
     * The region from which the AMI will be copied. This may be the
     * same as the AWS provider region in order to create a copy within the same region.
     */
    readonly sourceAmiRegion?: pulumi.Input<string>;
    /**
     * When set to "simple" (the default), enables enhanced networking
     * for created instances. No other value is supported at this time.
     */
    readonly sriovNetSupport?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly usageOperation?: pulumi.Input<string>;
    /**
     * Keyword to choose what virtualization mode created instances
     * will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
     * changes the set of further arguments that are required, as described below.
     */
    readonly virtualizationType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AmiCopy resource.
 */
export interface AmiCopyArgs {
    /**
     * A longer, human-readable description for the AMI.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Nested block describing an EBS block device that should be
     * attached to created instances. The structure of this block is described below.
     */
    readonly ebsBlockDevices?: pulumi.Input<pulumi.Input<inputs.ec2.AmiCopyEbsBlockDevice>[]>;
    /**
     * Boolean controlling whether the created EBS volumes will be encrypted. Can't be used with `snapshotId`.
     */
    readonly encrypted?: pulumi.Input<boolean>;
    /**
     * Nested block describing an ephemeral block device that
     * should be attached to created instances. The structure of this block is described below.
     */
    readonly ephemeralBlockDevices?: pulumi.Input<pulumi.Input<inputs.ec2.AmiCopyEphemeralBlockDevice>[]>;
    /**
     * The full ARN of the AWS Key Management Service (AWS KMS) CMK to use when encrypting the snapshots of
     * an image during a copy operation. This parameter is only required if you want to use a non-default CMK;
     * if this parameter is not specified, the default CMK for EBS is used
     */
    readonly kmsKeyId?: pulumi.Input<string>;
    /**
     * A region-unique name for the AMI.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The id of the AMI to copy. This id must be valid in the region
     * given by `sourceAmiRegion`.
     */
    readonly sourceAmiId: pulumi.Input<string>;
    /**
     * The region from which the AMI will be copied. This may be the
     * same as the AWS provider region in order to create a copy within the same region.
     */
    readonly sourceAmiRegion: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
