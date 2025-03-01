// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pinpoint

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Pinpoint ADM (Amazon Device Messaging) Channel resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/pinpoint"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		app, err := pinpoint.NewApp(ctx, "app", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = pinpoint.NewAdmChannel(ctx, "channel", &pinpoint.AdmChannelArgs{
// 			ApplicationId: app.ApplicationId,
// 			ClientId:      pulumi.String(""),
// 			ClientSecret:  pulumi.String(""),
// 			Enabled:       pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Pinpoint ADM Channel can be imported using the `application-id`, e.g.
//
// ```sh
//  $ pulumi import aws:pinpoint/admChannel:AdmChannel channel application-id
// ```
type AdmChannel struct {
	pulumi.CustomResourceState

	// The application ID.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// Client ID (part of OAuth Credentials) obtained via Amazon Developer Account.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// Client Secret (part of OAuth Credentials) obtained via Amazon Developer Account.
	ClientSecret pulumi.StringOutput `pulumi:"clientSecret"`
	// Specifies whether to enable the channel. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
}

// NewAdmChannel registers a new resource with the given unique name, arguments, and options.
func NewAdmChannel(ctx *pulumi.Context,
	name string, args *AdmChannelArgs, opts ...pulumi.ResourceOption) (*AdmChannel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	if args.ClientId == nil {
		return nil, errors.New("invalid value for required argument 'ClientId'")
	}
	if args.ClientSecret == nil {
		return nil, errors.New("invalid value for required argument 'ClientSecret'")
	}
	var resource AdmChannel
	err := ctx.RegisterResource("aws:pinpoint/admChannel:AdmChannel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAdmChannel gets an existing AdmChannel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAdmChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AdmChannelState, opts ...pulumi.ResourceOption) (*AdmChannel, error) {
	var resource AdmChannel
	err := ctx.ReadResource("aws:pinpoint/admChannel:AdmChannel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AdmChannel resources.
type admChannelState struct {
	// The application ID.
	ApplicationId *string `pulumi:"applicationId"`
	// Client ID (part of OAuth Credentials) obtained via Amazon Developer Account.
	ClientId *string `pulumi:"clientId"`
	// Client Secret (part of OAuth Credentials) obtained via Amazon Developer Account.
	ClientSecret *string `pulumi:"clientSecret"`
	// Specifies whether to enable the channel. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
}

type AdmChannelState struct {
	// The application ID.
	ApplicationId pulumi.StringPtrInput
	// Client ID (part of OAuth Credentials) obtained via Amazon Developer Account.
	ClientId pulumi.StringPtrInput
	// Client Secret (part of OAuth Credentials) obtained via Amazon Developer Account.
	ClientSecret pulumi.StringPtrInput
	// Specifies whether to enable the channel. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
}

func (AdmChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*admChannelState)(nil)).Elem()
}

type admChannelArgs struct {
	// The application ID.
	ApplicationId string `pulumi:"applicationId"`
	// Client ID (part of OAuth Credentials) obtained via Amazon Developer Account.
	ClientId string `pulumi:"clientId"`
	// Client Secret (part of OAuth Credentials) obtained via Amazon Developer Account.
	ClientSecret string `pulumi:"clientSecret"`
	// Specifies whether to enable the channel. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
}

// The set of arguments for constructing a AdmChannel resource.
type AdmChannelArgs struct {
	// The application ID.
	ApplicationId pulumi.StringInput
	// Client ID (part of OAuth Credentials) obtained via Amazon Developer Account.
	ClientId pulumi.StringInput
	// Client Secret (part of OAuth Credentials) obtained via Amazon Developer Account.
	ClientSecret pulumi.StringInput
	// Specifies whether to enable the channel. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
}

func (AdmChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*admChannelArgs)(nil)).Elem()
}

type AdmChannelInput interface {
	pulumi.Input

	ToAdmChannelOutput() AdmChannelOutput
	ToAdmChannelOutputWithContext(ctx context.Context) AdmChannelOutput
}

func (*AdmChannel) ElementType() reflect.Type {
	return reflect.TypeOf((*AdmChannel)(nil))
}

func (i *AdmChannel) ToAdmChannelOutput() AdmChannelOutput {
	return i.ToAdmChannelOutputWithContext(context.Background())
}

func (i *AdmChannel) ToAdmChannelOutputWithContext(ctx context.Context) AdmChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdmChannelOutput)
}

func (i *AdmChannel) ToAdmChannelPtrOutput() AdmChannelPtrOutput {
	return i.ToAdmChannelPtrOutputWithContext(context.Background())
}

func (i *AdmChannel) ToAdmChannelPtrOutputWithContext(ctx context.Context) AdmChannelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdmChannelPtrOutput)
}

type AdmChannelPtrInput interface {
	pulumi.Input

	ToAdmChannelPtrOutput() AdmChannelPtrOutput
	ToAdmChannelPtrOutputWithContext(ctx context.Context) AdmChannelPtrOutput
}

type admChannelPtrType AdmChannelArgs

func (*admChannelPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AdmChannel)(nil))
}

func (i *admChannelPtrType) ToAdmChannelPtrOutput() AdmChannelPtrOutput {
	return i.ToAdmChannelPtrOutputWithContext(context.Background())
}

func (i *admChannelPtrType) ToAdmChannelPtrOutputWithContext(ctx context.Context) AdmChannelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdmChannelPtrOutput)
}

// AdmChannelArrayInput is an input type that accepts AdmChannelArray and AdmChannelArrayOutput values.
// You can construct a concrete instance of `AdmChannelArrayInput` via:
//
//          AdmChannelArray{ AdmChannelArgs{...} }
type AdmChannelArrayInput interface {
	pulumi.Input

	ToAdmChannelArrayOutput() AdmChannelArrayOutput
	ToAdmChannelArrayOutputWithContext(context.Context) AdmChannelArrayOutput
}

type AdmChannelArray []AdmChannelInput

func (AdmChannelArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*AdmChannel)(nil))
}

func (i AdmChannelArray) ToAdmChannelArrayOutput() AdmChannelArrayOutput {
	return i.ToAdmChannelArrayOutputWithContext(context.Background())
}

func (i AdmChannelArray) ToAdmChannelArrayOutputWithContext(ctx context.Context) AdmChannelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdmChannelArrayOutput)
}

// AdmChannelMapInput is an input type that accepts AdmChannelMap and AdmChannelMapOutput values.
// You can construct a concrete instance of `AdmChannelMapInput` via:
//
//          AdmChannelMap{ "key": AdmChannelArgs{...} }
type AdmChannelMapInput interface {
	pulumi.Input

	ToAdmChannelMapOutput() AdmChannelMapOutput
	ToAdmChannelMapOutputWithContext(context.Context) AdmChannelMapOutput
}

type AdmChannelMap map[string]AdmChannelInput

func (AdmChannelMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*AdmChannel)(nil))
}

func (i AdmChannelMap) ToAdmChannelMapOutput() AdmChannelMapOutput {
	return i.ToAdmChannelMapOutputWithContext(context.Background())
}

func (i AdmChannelMap) ToAdmChannelMapOutputWithContext(ctx context.Context) AdmChannelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdmChannelMapOutput)
}

type AdmChannelOutput struct {
	*pulumi.OutputState
}

func (AdmChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdmChannel)(nil))
}

func (o AdmChannelOutput) ToAdmChannelOutput() AdmChannelOutput {
	return o
}

func (o AdmChannelOutput) ToAdmChannelOutputWithContext(ctx context.Context) AdmChannelOutput {
	return o
}

func (o AdmChannelOutput) ToAdmChannelPtrOutput() AdmChannelPtrOutput {
	return o.ToAdmChannelPtrOutputWithContext(context.Background())
}

func (o AdmChannelOutput) ToAdmChannelPtrOutputWithContext(ctx context.Context) AdmChannelPtrOutput {
	return o.ApplyT(func(v AdmChannel) *AdmChannel {
		return &v
	}).(AdmChannelPtrOutput)
}

type AdmChannelPtrOutput struct {
	*pulumi.OutputState
}

func (AdmChannelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AdmChannel)(nil))
}

func (o AdmChannelPtrOutput) ToAdmChannelPtrOutput() AdmChannelPtrOutput {
	return o
}

func (o AdmChannelPtrOutput) ToAdmChannelPtrOutputWithContext(ctx context.Context) AdmChannelPtrOutput {
	return o
}

type AdmChannelArrayOutput struct{ *pulumi.OutputState }

func (AdmChannelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AdmChannel)(nil))
}

func (o AdmChannelArrayOutput) ToAdmChannelArrayOutput() AdmChannelArrayOutput {
	return o
}

func (o AdmChannelArrayOutput) ToAdmChannelArrayOutputWithContext(ctx context.Context) AdmChannelArrayOutput {
	return o
}

func (o AdmChannelArrayOutput) Index(i pulumi.IntInput) AdmChannelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AdmChannel {
		return vs[0].([]AdmChannel)[vs[1].(int)]
	}).(AdmChannelOutput)
}

type AdmChannelMapOutput struct{ *pulumi.OutputState }

func (AdmChannelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AdmChannel)(nil))
}

func (o AdmChannelMapOutput) ToAdmChannelMapOutput() AdmChannelMapOutput {
	return o
}

func (o AdmChannelMapOutput) ToAdmChannelMapOutputWithContext(ctx context.Context) AdmChannelMapOutput {
	return o
}

func (o AdmChannelMapOutput) MapIndex(k pulumi.StringInput) AdmChannelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AdmChannel {
		return vs[0].(map[string]AdmChannel)[vs[1].(string)]
	}).(AdmChannelOutput)
}

func init() {
	pulumi.RegisterOutputType(AdmChannelOutput{})
	pulumi.RegisterOutputType(AdmChannelPtrOutput{})
	pulumi.RegisterOutputType(AdmChannelArrayOutput{})
	pulumi.RegisterOutputType(AdmChannelMapOutput{})
}
