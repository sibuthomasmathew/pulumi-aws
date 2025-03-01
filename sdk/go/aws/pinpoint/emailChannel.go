// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pinpoint

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Pinpoint Email Channel resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/iam"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/pinpoint"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ses"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		app, err := pinpoint.NewApp(ctx, "app", nil)
// 		if err != nil {
// 			return err
// 		}
// 		identity, err := ses.NewDomainIdentity(ctx, "identity", &ses.DomainIdentityArgs{
// 			Domain: pulumi.String("example.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		role, err := iam.NewRole(ctx, "role", &iam.RoleArgs{
// 			AssumeRolePolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Action\": \"sts:AssumeRole\",\n", "      \"Principal\": {\n", "        \"Service\": \"pinpoint.amazonaws.com\"\n", "      },\n", "      \"Effect\": \"Allow\",\n", "      \"Sid\": \"\"\n", "    }\n", "  ]\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = pinpoint.NewEmailChannel(ctx, "email", &pinpoint.EmailChannelArgs{
// 			ApplicationId: app.ApplicationId,
// 			FromAddress:   pulumi.String("user@example.com"),
// 			Identity:      identity.Arn,
// 			RoleArn:       role.Arn,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewRolePolicy(ctx, "rolePolicy", &iam.RolePolicyArgs{
// 			Role:   role.ID(),
// 			Policy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": {\n", "    \"Action\": [\n", "      \"mobileanalytics:PutEvents\",\n", "      \"mobileanalytics:PutItems\"\n", "    ],\n", "    \"Effect\": \"Allow\",\n", "    \"Resource\": [\n", "      \"*\"\n", "    ]\n", "  }\n", "}\n")),
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
// Pinpoint Email Channel can be imported using the `application-id`, e.g.
//
// ```sh
//  $ pulumi import aws:pinpoint/emailChannel:EmailChannel email application-id
// ```
type EmailChannel struct {
	pulumi.CustomResourceState

	// The application ID.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// The ARN of the Amazon SES configuration set that you want to apply to messages that you send through the channel.
	ConfigurationSet pulumi.StringPtrOutput `pulumi:"configurationSet"`
	// Whether the channel is enabled or disabled. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The email address used to send emails from.
	FromAddress pulumi.StringOutput `pulumi:"fromAddress"`
	// The ARN of an identity verified with SES.
	Identity pulumi.StringOutput `pulumi:"identity"`
	// Messages per second that can be sent.
	MessagesPerSecond pulumi.IntOutput `pulumi:"messagesPerSecond"`
	// The ARN of an IAM Role used to submit events to Mobile Analytics' event ingestion service.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
}

// NewEmailChannel registers a new resource with the given unique name, arguments, and options.
func NewEmailChannel(ctx *pulumi.Context,
	name string, args *EmailChannelArgs, opts ...pulumi.ResourceOption) (*EmailChannel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	if args.FromAddress == nil {
		return nil, errors.New("invalid value for required argument 'FromAddress'")
	}
	if args.Identity == nil {
		return nil, errors.New("invalid value for required argument 'Identity'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	var resource EmailChannel
	err := ctx.RegisterResource("aws:pinpoint/emailChannel:EmailChannel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEmailChannel gets an existing EmailChannel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEmailChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EmailChannelState, opts ...pulumi.ResourceOption) (*EmailChannel, error) {
	var resource EmailChannel
	err := ctx.ReadResource("aws:pinpoint/emailChannel:EmailChannel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EmailChannel resources.
type emailChannelState struct {
	// The application ID.
	ApplicationId *string `pulumi:"applicationId"`
	// The ARN of the Amazon SES configuration set that you want to apply to messages that you send through the channel.
	ConfigurationSet *string `pulumi:"configurationSet"`
	// Whether the channel is enabled or disabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The email address used to send emails from.
	FromAddress *string `pulumi:"fromAddress"`
	// The ARN of an identity verified with SES.
	Identity *string `pulumi:"identity"`
	// Messages per second that can be sent.
	MessagesPerSecond *int `pulumi:"messagesPerSecond"`
	// The ARN of an IAM Role used to submit events to Mobile Analytics' event ingestion service.
	RoleArn *string `pulumi:"roleArn"`
}

type EmailChannelState struct {
	// The application ID.
	ApplicationId pulumi.StringPtrInput
	// The ARN of the Amazon SES configuration set that you want to apply to messages that you send through the channel.
	ConfigurationSet pulumi.StringPtrInput
	// Whether the channel is enabled or disabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The email address used to send emails from.
	FromAddress pulumi.StringPtrInput
	// The ARN of an identity verified with SES.
	Identity pulumi.StringPtrInput
	// Messages per second that can be sent.
	MessagesPerSecond pulumi.IntPtrInput
	// The ARN of an IAM Role used to submit events to Mobile Analytics' event ingestion service.
	RoleArn pulumi.StringPtrInput
}

func (EmailChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*emailChannelState)(nil)).Elem()
}

type emailChannelArgs struct {
	// The application ID.
	ApplicationId string `pulumi:"applicationId"`
	// The ARN of the Amazon SES configuration set that you want to apply to messages that you send through the channel.
	ConfigurationSet *string `pulumi:"configurationSet"`
	// Whether the channel is enabled or disabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The email address used to send emails from.
	FromAddress string `pulumi:"fromAddress"`
	// The ARN of an identity verified with SES.
	Identity string `pulumi:"identity"`
	// The ARN of an IAM Role used to submit events to Mobile Analytics' event ingestion service.
	RoleArn string `pulumi:"roleArn"`
}

// The set of arguments for constructing a EmailChannel resource.
type EmailChannelArgs struct {
	// The application ID.
	ApplicationId pulumi.StringInput
	// The ARN of the Amazon SES configuration set that you want to apply to messages that you send through the channel.
	ConfigurationSet pulumi.StringPtrInput
	// Whether the channel is enabled or disabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The email address used to send emails from.
	FromAddress pulumi.StringInput
	// The ARN of an identity verified with SES.
	Identity pulumi.StringInput
	// The ARN of an IAM Role used to submit events to Mobile Analytics' event ingestion service.
	RoleArn pulumi.StringInput
}

func (EmailChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*emailChannelArgs)(nil)).Elem()
}

type EmailChannelInput interface {
	pulumi.Input

	ToEmailChannelOutput() EmailChannelOutput
	ToEmailChannelOutputWithContext(ctx context.Context) EmailChannelOutput
}

func (*EmailChannel) ElementType() reflect.Type {
	return reflect.TypeOf((*EmailChannel)(nil))
}

func (i *EmailChannel) ToEmailChannelOutput() EmailChannelOutput {
	return i.ToEmailChannelOutputWithContext(context.Background())
}

func (i *EmailChannel) ToEmailChannelOutputWithContext(ctx context.Context) EmailChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailChannelOutput)
}

func (i *EmailChannel) ToEmailChannelPtrOutput() EmailChannelPtrOutput {
	return i.ToEmailChannelPtrOutputWithContext(context.Background())
}

func (i *EmailChannel) ToEmailChannelPtrOutputWithContext(ctx context.Context) EmailChannelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailChannelPtrOutput)
}

type EmailChannelPtrInput interface {
	pulumi.Input

	ToEmailChannelPtrOutput() EmailChannelPtrOutput
	ToEmailChannelPtrOutputWithContext(ctx context.Context) EmailChannelPtrOutput
}

type emailChannelPtrType EmailChannelArgs

func (*emailChannelPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailChannel)(nil))
}

func (i *emailChannelPtrType) ToEmailChannelPtrOutput() EmailChannelPtrOutput {
	return i.ToEmailChannelPtrOutputWithContext(context.Background())
}

func (i *emailChannelPtrType) ToEmailChannelPtrOutputWithContext(ctx context.Context) EmailChannelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailChannelPtrOutput)
}

// EmailChannelArrayInput is an input type that accepts EmailChannelArray and EmailChannelArrayOutput values.
// You can construct a concrete instance of `EmailChannelArrayInput` via:
//
//          EmailChannelArray{ EmailChannelArgs{...} }
type EmailChannelArrayInput interface {
	pulumi.Input

	ToEmailChannelArrayOutput() EmailChannelArrayOutput
	ToEmailChannelArrayOutputWithContext(context.Context) EmailChannelArrayOutput
}

type EmailChannelArray []EmailChannelInput

func (EmailChannelArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*EmailChannel)(nil))
}

func (i EmailChannelArray) ToEmailChannelArrayOutput() EmailChannelArrayOutput {
	return i.ToEmailChannelArrayOutputWithContext(context.Background())
}

func (i EmailChannelArray) ToEmailChannelArrayOutputWithContext(ctx context.Context) EmailChannelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailChannelArrayOutput)
}

// EmailChannelMapInput is an input type that accepts EmailChannelMap and EmailChannelMapOutput values.
// You can construct a concrete instance of `EmailChannelMapInput` via:
//
//          EmailChannelMap{ "key": EmailChannelArgs{...} }
type EmailChannelMapInput interface {
	pulumi.Input

	ToEmailChannelMapOutput() EmailChannelMapOutput
	ToEmailChannelMapOutputWithContext(context.Context) EmailChannelMapOutput
}

type EmailChannelMap map[string]EmailChannelInput

func (EmailChannelMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*EmailChannel)(nil))
}

func (i EmailChannelMap) ToEmailChannelMapOutput() EmailChannelMapOutput {
	return i.ToEmailChannelMapOutputWithContext(context.Background())
}

func (i EmailChannelMap) ToEmailChannelMapOutputWithContext(ctx context.Context) EmailChannelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailChannelMapOutput)
}

type EmailChannelOutput struct {
	*pulumi.OutputState
}

func (EmailChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EmailChannel)(nil))
}

func (o EmailChannelOutput) ToEmailChannelOutput() EmailChannelOutput {
	return o
}

func (o EmailChannelOutput) ToEmailChannelOutputWithContext(ctx context.Context) EmailChannelOutput {
	return o
}

func (o EmailChannelOutput) ToEmailChannelPtrOutput() EmailChannelPtrOutput {
	return o.ToEmailChannelPtrOutputWithContext(context.Background())
}

func (o EmailChannelOutput) ToEmailChannelPtrOutputWithContext(ctx context.Context) EmailChannelPtrOutput {
	return o.ApplyT(func(v EmailChannel) *EmailChannel {
		return &v
	}).(EmailChannelPtrOutput)
}

type EmailChannelPtrOutput struct {
	*pulumi.OutputState
}

func (EmailChannelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailChannel)(nil))
}

func (o EmailChannelPtrOutput) ToEmailChannelPtrOutput() EmailChannelPtrOutput {
	return o
}

func (o EmailChannelPtrOutput) ToEmailChannelPtrOutputWithContext(ctx context.Context) EmailChannelPtrOutput {
	return o
}

type EmailChannelArrayOutput struct{ *pulumi.OutputState }

func (EmailChannelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EmailChannel)(nil))
}

func (o EmailChannelArrayOutput) ToEmailChannelArrayOutput() EmailChannelArrayOutput {
	return o
}

func (o EmailChannelArrayOutput) ToEmailChannelArrayOutputWithContext(ctx context.Context) EmailChannelArrayOutput {
	return o
}

func (o EmailChannelArrayOutput) Index(i pulumi.IntInput) EmailChannelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EmailChannel {
		return vs[0].([]EmailChannel)[vs[1].(int)]
	}).(EmailChannelOutput)
}

type EmailChannelMapOutput struct{ *pulumi.OutputState }

func (EmailChannelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]EmailChannel)(nil))
}

func (o EmailChannelMapOutput) ToEmailChannelMapOutput() EmailChannelMapOutput {
	return o
}

func (o EmailChannelMapOutput) ToEmailChannelMapOutputWithContext(ctx context.Context) EmailChannelMapOutput {
	return o
}

func (o EmailChannelMapOutput) MapIndex(k pulumi.StringInput) EmailChannelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) EmailChannel {
		return vs[0].(map[string]EmailChannel)[vs[1].(string)]
	}).(EmailChannelOutput)
}

func init() {
	pulumi.RegisterOutputType(EmailChannelOutput{})
	pulumi.RegisterOutputType(EmailChannelPtrOutput{})
	pulumi.RegisterOutputType(EmailChannelArrayOutput{})
	pulumi.RegisterOutputType(EmailChannelMapOutput{})
}
