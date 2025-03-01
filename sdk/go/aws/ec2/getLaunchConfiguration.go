// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides information about a Launch Configuration.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ec2.LookupLaunchConfiguration(ctx, &ec2.LookupLaunchConfigurationArgs{
// 			Name: "test-launch-config",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupLaunchConfiguration(ctx *pulumi.Context, args *LookupLaunchConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupLaunchConfigurationResult, error) {
	var rv LookupLaunchConfigurationResult
	err := ctx.Invoke("aws:ec2/getLaunchConfiguration:getLaunchConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLaunchConfiguration.
type LookupLaunchConfigurationArgs struct {
	// The name of the launch configuration.
	Name string `pulumi:"name"`
}

// A collection of values returned by getLaunchConfiguration.
type LookupLaunchConfigurationResult struct {
	// The Amazon Resource Name of the launch configuration.
	Arn string `pulumi:"arn"`
	// Whether a Public IP address is associated with the instance.
	AssociatePublicIpAddress bool `pulumi:"associatePublicIpAddress"`
	// The EBS Block Devices attached to the instance.
	EbsBlockDevices []GetLaunchConfigurationEbsBlockDevice `pulumi:"ebsBlockDevices"`
	// Whether the launched EC2 instance will be EBS-optimized.
	EbsOptimized bool `pulumi:"ebsOptimized"`
	// Whether Detailed Monitoring is Enabled.
	EnableMonitoring bool `pulumi:"enableMonitoring"`
	// The Ephemeral volumes on the instance.
	EphemeralBlockDevices []GetLaunchConfigurationEphemeralBlockDevice `pulumi:"ephemeralBlockDevices"`
	// The IAM Instance Profile to associate with launched instances.
	IamInstanceProfile string `pulumi:"iamInstanceProfile"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The EC2 Image ID of the instance.
	ImageId string `pulumi:"imageId"`
	// The Instance Type of the instance to launch.
	InstanceType string `pulumi:"instanceType"`
	// The Key Name that should be used for the instance.
	KeyName string `pulumi:"keyName"`
	// The metadata options for the instance.
	MetadataOptions []GetLaunchConfigurationMetadataOption `pulumi:"metadataOptions"`
	// The Name of the launch configuration.
	Name string `pulumi:"name"`
	// The Tenancy of the instance.
	PlacementTenancy string `pulumi:"placementTenancy"`
	// The Root Block Device of the instance.
	RootBlockDevices []GetLaunchConfigurationRootBlockDevice `pulumi:"rootBlockDevices"`
	// A list of associated Security Group IDS.
	SecurityGroups []string `pulumi:"securityGroups"`
	// The Price to use for reserving Spot instances.
	SpotPrice string `pulumi:"spotPrice"`
	// The User Data of the instance.
	UserData string `pulumi:"userData"`
	// The ID of a ClassicLink-enabled VPC.
	VpcClassicLinkId string `pulumi:"vpcClassicLinkId"`
	// The IDs of one or more Security Groups for the specified ClassicLink-enabled VPC.
	VpcClassicLinkSecurityGroups []string `pulumi:"vpcClassicLinkSecurityGroups"`
}
