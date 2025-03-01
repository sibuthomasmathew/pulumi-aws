// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ssoadmin

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to get ARNs and Identity Store IDs of Single Sign-On (SSO) Instances.
func GetInstances(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetInstancesResult, error) {
	var rv GetInstancesResult
	err := ctx.Invoke("aws:ssoadmin/getInstances:getInstances", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getInstances.
type GetInstancesResult struct {
	// Set of Amazon Resource Names (ARNs) of the SSO Instances.
	Arns []string `pulumi:"arns"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Set of identifiers of the identity stores connected to the SSO Instances.
	IdentityStoreIds []string `pulumi:"identityStoreIds"`
}
