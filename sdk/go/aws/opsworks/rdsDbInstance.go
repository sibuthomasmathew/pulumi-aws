// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package opsworks

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an OpsWorks RDS DB Instance resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/opsworks"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := opsworks.NewRdsDbInstance(ctx, "myInstance", &opsworks.RdsDbInstanceArgs{
// 			StackId:          pulumi.Any(aws_opsworks_stack.My_stack.Id),
// 			RdsDbInstanceArn: pulumi.Any(aws_db_instance.My_instance.Arn),
// 			DbUser:           pulumi.String("someUser"),
// 			DbPassword:       pulumi.String("somePass"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type RdsDbInstance struct {
	pulumi.CustomResourceState

	// A db password
	DbPassword pulumi.StringOutput `pulumi:"dbPassword"`
	// A db username
	DbUser pulumi.StringOutput `pulumi:"dbUser"`
	// The db instance to register for this stack. Changing this will force a new resource.
	RdsDbInstanceArn pulumi.StringOutput `pulumi:"rdsDbInstanceArn"`
	// The stack to register a db instance for. Changing this will force a new resource.
	StackId pulumi.StringOutput `pulumi:"stackId"`
}

// NewRdsDbInstance registers a new resource with the given unique name, arguments, and options.
func NewRdsDbInstance(ctx *pulumi.Context,
	name string, args *RdsDbInstanceArgs, opts ...pulumi.ResourceOption) (*RdsDbInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DbPassword == nil {
		return nil, errors.New("invalid value for required argument 'DbPassword'")
	}
	if args.DbUser == nil {
		return nil, errors.New("invalid value for required argument 'DbUser'")
	}
	if args.RdsDbInstanceArn == nil {
		return nil, errors.New("invalid value for required argument 'RdsDbInstanceArn'")
	}
	if args.StackId == nil {
		return nil, errors.New("invalid value for required argument 'StackId'")
	}
	var resource RdsDbInstance
	err := ctx.RegisterResource("aws:opsworks/rdsDbInstance:RdsDbInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRdsDbInstance gets an existing RdsDbInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRdsDbInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RdsDbInstanceState, opts ...pulumi.ResourceOption) (*RdsDbInstance, error) {
	var resource RdsDbInstance
	err := ctx.ReadResource("aws:opsworks/rdsDbInstance:RdsDbInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RdsDbInstance resources.
type rdsDbInstanceState struct {
	// A db password
	DbPassword *string `pulumi:"dbPassword"`
	// A db username
	DbUser *string `pulumi:"dbUser"`
	// The db instance to register for this stack. Changing this will force a new resource.
	RdsDbInstanceArn *string `pulumi:"rdsDbInstanceArn"`
	// The stack to register a db instance for. Changing this will force a new resource.
	StackId *string `pulumi:"stackId"`
}

type RdsDbInstanceState struct {
	// A db password
	DbPassword pulumi.StringPtrInput
	// A db username
	DbUser pulumi.StringPtrInput
	// The db instance to register for this stack. Changing this will force a new resource.
	RdsDbInstanceArn pulumi.StringPtrInput
	// The stack to register a db instance for. Changing this will force a new resource.
	StackId pulumi.StringPtrInput
}

func (RdsDbInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*rdsDbInstanceState)(nil)).Elem()
}

type rdsDbInstanceArgs struct {
	// A db password
	DbPassword string `pulumi:"dbPassword"`
	// A db username
	DbUser string `pulumi:"dbUser"`
	// The db instance to register for this stack. Changing this will force a new resource.
	RdsDbInstanceArn string `pulumi:"rdsDbInstanceArn"`
	// The stack to register a db instance for. Changing this will force a new resource.
	StackId string `pulumi:"stackId"`
}

// The set of arguments for constructing a RdsDbInstance resource.
type RdsDbInstanceArgs struct {
	// A db password
	DbPassword pulumi.StringInput
	// A db username
	DbUser pulumi.StringInput
	// The db instance to register for this stack. Changing this will force a new resource.
	RdsDbInstanceArn pulumi.StringInput
	// The stack to register a db instance for. Changing this will force a new resource.
	StackId pulumi.StringInput
}

func (RdsDbInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rdsDbInstanceArgs)(nil)).Elem()
}

type RdsDbInstanceInput interface {
	pulumi.Input

	ToRdsDbInstanceOutput() RdsDbInstanceOutput
	ToRdsDbInstanceOutputWithContext(ctx context.Context) RdsDbInstanceOutput
}

func (*RdsDbInstance) ElementType() reflect.Type {
	return reflect.TypeOf((*RdsDbInstance)(nil))
}

func (i *RdsDbInstance) ToRdsDbInstanceOutput() RdsDbInstanceOutput {
	return i.ToRdsDbInstanceOutputWithContext(context.Background())
}

func (i *RdsDbInstance) ToRdsDbInstanceOutputWithContext(ctx context.Context) RdsDbInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdsDbInstanceOutput)
}

func (i *RdsDbInstance) ToRdsDbInstancePtrOutput() RdsDbInstancePtrOutput {
	return i.ToRdsDbInstancePtrOutputWithContext(context.Background())
}

func (i *RdsDbInstance) ToRdsDbInstancePtrOutputWithContext(ctx context.Context) RdsDbInstancePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdsDbInstancePtrOutput)
}

type RdsDbInstancePtrInput interface {
	pulumi.Input

	ToRdsDbInstancePtrOutput() RdsDbInstancePtrOutput
	ToRdsDbInstancePtrOutputWithContext(ctx context.Context) RdsDbInstancePtrOutput
}

type rdsDbInstancePtrType RdsDbInstanceArgs

func (*rdsDbInstancePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RdsDbInstance)(nil))
}

func (i *rdsDbInstancePtrType) ToRdsDbInstancePtrOutput() RdsDbInstancePtrOutput {
	return i.ToRdsDbInstancePtrOutputWithContext(context.Background())
}

func (i *rdsDbInstancePtrType) ToRdsDbInstancePtrOutputWithContext(ctx context.Context) RdsDbInstancePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdsDbInstancePtrOutput)
}

// RdsDbInstanceArrayInput is an input type that accepts RdsDbInstanceArray and RdsDbInstanceArrayOutput values.
// You can construct a concrete instance of `RdsDbInstanceArrayInput` via:
//
//          RdsDbInstanceArray{ RdsDbInstanceArgs{...} }
type RdsDbInstanceArrayInput interface {
	pulumi.Input

	ToRdsDbInstanceArrayOutput() RdsDbInstanceArrayOutput
	ToRdsDbInstanceArrayOutputWithContext(context.Context) RdsDbInstanceArrayOutput
}

type RdsDbInstanceArray []RdsDbInstanceInput

func (RdsDbInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*RdsDbInstance)(nil))
}

func (i RdsDbInstanceArray) ToRdsDbInstanceArrayOutput() RdsDbInstanceArrayOutput {
	return i.ToRdsDbInstanceArrayOutputWithContext(context.Background())
}

func (i RdsDbInstanceArray) ToRdsDbInstanceArrayOutputWithContext(ctx context.Context) RdsDbInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdsDbInstanceArrayOutput)
}

// RdsDbInstanceMapInput is an input type that accepts RdsDbInstanceMap and RdsDbInstanceMapOutput values.
// You can construct a concrete instance of `RdsDbInstanceMapInput` via:
//
//          RdsDbInstanceMap{ "key": RdsDbInstanceArgs{...} }
type RdsDbInstanceMapInput interface {
	pulumi.Input

	ToRdsDbInstanceMapOutput() RdsDbInstanceMapOutput
	ToRdsDbInstanceMapOutputWithContext(context.Context) RdsDbInstanceMapOutput
}

type RdsDbInstanceMap map[string]RdsDbInstanceInput

func (RdsDbInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*RdsDbInstance)(nil))
}

func (i RdsDbInstanceMap) ToRdsDbInstanceMapOutput() RdsDbInstanceMapOutput {
	return i.ToRdsDbInstanceMapOutputWithContext(context.Background())
}

func (i RdsDbInstanceMap) ToRdsDbInstanceMapOutputWithContext(ctx context.Context) RdsDbInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdsDbInstanceMapOutput)
}

type RdsDbInstanceOutput struct {
	*pulumi.OutputState
}

func (RdsDbInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RdsDbInstance)(nil))
}

func (o RdsDbInstanceOutput) ToRdsDbInstanceOutput() RdsDbInstanceOutput {
	return o
}

func (o RdsDbInstanceOutput) ToRdsDbInstanceOutputWithContext(ctx context.Context) RdsDbInstanceOutput {
	return o
}

func (o RdsDbInstanceOutput) ToRdsDbInstancePtrOutput() RdsDbInstancePtrOutput {
	return o.ToRdsDbInstancePtrOutputWithContext(context.Background())
}

func (o RdsDbInstanceOutput) ToRdsDbInstancePtrOutputWithContext(ctx context.Context) RdsDbInstancePtrOutput {
	return o.ApplyT(func(v RdsDbInstance) *RdsDbInstance {
		return &v
	}).(RdsDbInstancePtrOutput)
}

type RdsDbInstancePtrOutput struct {
	*pulumi.OutputState
}

func (RdsDbInstancePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RdsDbInstance)(nil))
}

func (o RdsDbInstancePtrOutput) ToRdsDbInstancePtrOutput() RdsDbInstancePtrOutput {
	return o
}

func (o RdsDbInstancePtrOutput) ToRdsDbInstancePtrOutputWithContext(ctx context.Context) RdsDbInstancePtrOutput {
	return o
}

type RdsDbInstanceArrayOutput struct{ *pulumi.OutputState }

func (RdsDbInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RdsDbInstance)(nil))
}

func (o RdsDbInstanceArrayOutput) ToRdsDbInstanceArrayOutput() RdsDbInstanceArrayOutput {
	return o
}

func (o RdsDbInstanceArrayOutput) ToRdsDbInstanceArrayOutputWithContext(ctx context.Context) RdsDbInstanceArrayOutput {
	return o
}

func (o RdsDbInstanceArrayOutput) Index(i pulumi.IntInput) RdsDbInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RdsDbInstance {
		return vs[0].([]RdsDbInstance)[vs[1].(int)]
	}).(RdsDbInstanceOutput)
}

type RdsDbInstanceMapOutput struct{ *pulumi.OutputState }

func (RdsDbInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RdsDbInstance)(nil))
}

func (o RdsDbInstanceMapOutput) ToRdsDbInstanceMapOutput() RdsDbInstanceMapOutput {
	return o
}

func (o RdsDbInstanceMapOutput) ToRdsDbInstanceMapOutputWithContext(ctx context.Context) RdsDbInstanceMapOutput {
	return o
}

func (o RdsDbInstanceMapOutput) MapIndex(k pulumi.StringInput) RdsDbInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RdsDbInstance {
		return vs[0].(map[string]RdsDbInstance)[vs[1].(string)]
	}).(RdsDbInstanceOutput)
}

func init() {
	pulumi.RegisterOutputType(RdsDbInstanceOutput{})
	pulumi.RegisterOutputType(RdsDbInstancePtrOutput{})
	pulumi.RegisterOutputType(RdsDbInstanceArrayOutput{})
	pulumi.RegisterOutputType(RdsDbInstanceMapOutput{})
}
