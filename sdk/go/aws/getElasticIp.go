// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// `ec2.Eip` provides details about a specific Elastic IP.
//
// ## Example Usage
// ### Search By Allocation ID (VPC only)
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
// 		opt0 := "eipalloc-12345678"
// 		_, err := ec2.GetElasticIp(ctx, &ec2.GetElasticIpArgs{
// 			Id: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Search By Filters (EC2-Classic or VPC)
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
// 		_, err := ec2.GetElasticIp(ctx, &ec2.GetElasticIpArgs{
// 			Filters: []ec2.GetElasticIpFilter{
// 				ec2.GetElasticIpFilter{
// 					Name: "tag:Name",
// 					Values: []string{
// 						"exampleNameTagValue",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Search By Public IP (EC2-Classic or VPC)
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
// 		opt0 := "1.2.3.4"
// 		_, err := ec2.GetElasticIp(ctx, &ec2.GetElasticIpArgs{
// 			PublicIp: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Search By Tags (EC2-Classic or VPC)
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
// 		_, err := ec2.GetElasticIp(ctx, &ec2.GetElasticIpArgs{
// 			Tags: map[string]interface{}{
// 				"Name": "exampleNameTagValue",
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Deprecated: aws.getElasticIp has been deprecated in favor of aws.ec2.getElasticIp
func GetElasticIp(ctx *pulumi.Context, args *GetElasticIpArgs, opts ...pulumi.InvokeOption) (*GetElasticIpResult, error) {
	var rv GetElasticIpResult
	err := ctx.Invoke("aws:index/getElasticIp:getElasticIp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getElasticIp.
type GetElasticIpArgs struct {
	// One or more name/value pairs to use as filters. There are several valid keys, for a full reference, check out the [EC2 API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeAddresses.html).
	Filters []GetElasticIpFilter `pulumi:"filters"`
	// The allocation id of the specific VPC EIP to retrieve. If a classic EIP is required, do NOT set `id`, only set `publicIp`
	Id *string `pulumi:"id"`
	// The public IP of the specific EIP to retrieve.
	PublicIp *string `pulumi:"publicIp"`
	// A map of tags, each pair of which must exactly match a pair on the desired Elastic IP
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getElasticIp.
type GetElasticIpResult struct {
	// The ID representing the association of the address with an instance in a VPC.
	AssociationId string `pulumi:"associationId"`
	// The carrier IP address.
	CarrierIp string `pulumi:"carrierIp"`
	// Customer Owned IP.
	CustomerOwnedIp string `pulumi:"customerOwnedIp"`
	// The ID of a Customer Owned IP Pool. For more on customer owned IP addressed check out [Customer-owned IP addresses guide](https://docs.aws.amazon.com/outposts/latest/userguide/outposts-networking-components.html#ip-addressing)
	CustomerOwnedIpv4Pool string `pulumi:"customerOwnedIpv4Pool"`
	// Indicates whether the address is for use in EC2-Classic (standard) or in a VPC (vpc).
	Domain  string               `pulumi:"domain"`
	Filters []GetElasticIpFilter `pulumi:"filters"`
	// If VPC Elastic IP, the allocation identifier. If EC2-Classic Elastic IP, the public IP address.
	Id string `pulumi:"id"`
	// The ID of the instance that the address is associated with (if any).
	InstanceId string `pulumi:"instanceId"`
	// The ID of the network interface.
	NetworkInterfaceId string `pulumi:"networkInterfaceId"`
	// The ID of the AWS account that owns the network interface.
	NetworkInterfaceOwnerId string `pulumi:"networkInterfaceOwnerId"`
	// The Private DNS associated with the Elastic IP address.
	PrivateDns string `pulumi:"privateDns"`
	// The private IP address associated with the Elastic IP address.
	PrivateIp string `pulumi:"privateIp"`
	// Public DNS associated with the Elastic IP address.
	PublicDns string `pulumi:"publicDns"`
	// Public IP address of Elastic IP.
	PublicIp string `pulumi:"publicIp"`
	// The ID of an address pool.
	PublicIpv4Pool string `pulumi:"publicIpv4Pool"`
	// Key-value map of tags associated with Elastic IP.
	Tags map[string]string `pulumi:"tags"`
}
