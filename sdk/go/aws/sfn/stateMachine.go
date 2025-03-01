// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sfn

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Step Function State Machine resource
//
// ## Example Usage
// ### Basic (Standard Workflow)
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/sfn"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := sfn.NewStateMachine(ctx, "sfnStateMachine", &sfn.StateMachineArgs{
// 			RoleArn:    pulumi.Any(aws_iam_role.Iam_for_sfn.Arn),
// 			Definition: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Comment\": \"A Hello World example of the Amazon States Language using an AWS Lambda Function\",\n", "  \"StartAt\": \"HelloWorld\",\n", "  \"States\": {\n", "    \"HelloWorld\": {\n", "      \"Type\": \"Task\",\n", "      \"Resource\": \"", aws_lambda_function.Lambda.Arn, "\",\n", "      \"End\": true\n", "    }\n", "  }\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Basic (Express Workflow)
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/sfn"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := sfn.NewStateMachine(ctx, "sfnStateMachine", &sfn.StateMachineArgs{
// 			RoleArn:    pulumi.Any(aws_iam_role.Iam_for_sfn.Arn),
// 			Type:       pulumi.String("EXPRESS"),
// 			Definition: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Comment\": \"A Hello World example of the Amazon States Language using an AWS Lambda Function\",\n", "  \"StartAt\": \"HelloWorld\",\n", "  \"States\": {\n", "    \"HelloWorld\": {\n", "      \"Type\": \"Task\",\n", "      \"Resource\": \"", aws_lambda_function.Lambda.Arn, "\",\n", "      \"End\": true\n", "    }\n", "  }\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Logging
//
// > *NOTE:* See the [AWS Step Functions Developer Guide](https://docs.aws.amazon.com/step-functions/latest/dg/welcome.html) for more information about enabling Step Function logging.
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/sfn"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := sfn.NewStateMachine(ctx, "sfnStateMachine", &sfn.StateMachineArgs{
// 			RoleArn:    pulumi.Any(aws_iam_role.Iam_for_sfn.Arn),
// 			Definition: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Comment\": \"A Hello World example of the Amazon States Language using an AWS Lambda Function\",\n", "  \"StartAt\": \"HelloWorld\",\n", "  \"States\": {\n", "    \"HelloWorld\": {\n", "      \"Type\": \"Task\",\n", "      \"Resource\": \"", aws_lambda_function.Lambda.Arn, "\",\n", "      \"End\": true\n", "    }\n", "  }\n", "}\n")),
// 			LoggingConfiguration: &sfn.StateMachineLoggingConfigurationArgs{
// 				LogDestination:       pulumi.String(fmt.Sprintf("%v%v", aws_cloudwatch_log_group.Log_group_for_sfn.Arn, ":*")),
// 				IncludeExecutionData: pulumi.Bool(true),
// 				Level:                pulumi.String("ERROR"),
// 			},
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
// State Machines can be imported using the `arn`, e.g.
//
// ```sh
//  $ pulumi import aws:sfn/stateMachine:StateMachine foo arn:aws:states:eu-west-1:123456789098:stateMachine:bar
// ```
type StateMachine struct {
	pulumi.CustomResourceState

	// The ARN of the state machine.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The date the state machine was created.
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// The Amazon States Language definition of the state machine.
	Definition pulumi.StringOutput `pulumi:"definition"`
	// Defines what execution history events are logged and where they are logged. The `loggingConfiguration` parameter is only valid when `type` is set to `EXPRESS`. Defaults to `OFF`. For more information see [Logging Express Workflows](https://docs.aws.amazon.com/step-functions/latest/dg/cw-logs.html) and [Log Levels](https://docs.aws.amazon.com/step-functions/latest/dg/cloudwatch-log-level.html) in the AWS Step Functions User Guide.
	LoggingConfiguration StateMachineLoggingConfigurationOutput `pulumi:"loggingConfiguration"`
	// The name of the state machine.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Amazon Resource Name (ARN) of the IAM role to use for this state machine.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
	// The current status of the state machine. Either "ACTIVE" or "DELETING".
	Status pulumi.StringOutput `pulumi:"status"`
	// Key-value map of resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Determines whether a Standard or Express state machine is created. The default is STANDARD. You cannot update the type of a state machine once it has been created. Valid Values: STANDARD | EXPRESS
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewStateMachine registers a new resource with the given unique name, arguments, and options.
func NewStateMachine(ctx *pulumi.Context,
	name string, args *StateMachineArgs, opts ...pulumi.ResourceOption) (*StateMachine, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Definition == nil {
		return nil, errors.New("invalid value for required argument 'Definition'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	var resource StateMachine
	err := ctx.RegisterResource("aws:sfn/stateMachine:StateMachine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStateMachine gets an existing StateMachine resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStateMachine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StateMachineState, opts ...pulumi.ResourceOption) (*StateMachine, error) {
	var resource StateMachine
	err := ctx.ReadResource("aws:sfn/stateMachine:StateMachine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StateMachine resources.
type stateMachineState struct {
	// The ARN of the state machine.
	Arn *string `pulumi:"arn"`
	// The date the state machine was created.
	CreationDate *string `pulumi:"creationDate"`
	// The Amazon States Language definition of the state machine.
	Definition *string `pulumi:"definition"`
	// Defines what execution history events are logged and where they are logged. The `loggingConfiguration` parameter is only valid when `type` is set to `EXPRESS`. Defaults to `OFF`. For more information see [Logging Express Workflows](https://docs.aws.amazon.com/step-functions/latest/dg/cw-logs.html) and [Log Levels](https://docs.aws.amazon.com/step-functions/latest/dg/cloudwatch-log-level.html) in the AWS Step Functions User Guide.
	LoggingConfiguration *StateMachineLoggingConfiguration `pulumi:"loggingConfiguration"`
	// The name of the state machine.
	Name *string `pulumi:"name"`
	// The Amazon Resource Name (ARN) of the IAM role to use for this state machine.
	RoleArn *string `pulumi:"roleArn"`
	// The current status of the state machine. Either "ACTIVE" or "DELETING".
	Status *string `pulumi:"status"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
	// Determines whether a Standard or Express state machine is created. The default is STANDARD. You cannot update the type of a state machine once it has been created. Valid Values: STANDARD | EXPRESS
	Type *string `pulumi:"type"`
}

type StateMachineState struct {
	// The ARN of the state machine.
	Arn pulumi.StringPtrInput
	// The date the state machine was created.
	CreationDate pulumi.StringPtrInput
	// The Amazon States Language definition of the state machine.
	Definition pulumi.StringPtrInput
	// Defines what execution history events are logged and where they are logged. The `loggingConfiguration` parameter is only valid when `type` is set to `EXPRESS`. Defaults to `OFF`. For more information see [Logging Express Workflows](https://docs.aws.amazon.com/step-functions/latest/dg/cw-logs.html) and [Log Levels](https://docs.aws.amazon.com/step-functions/latest/dg/cloudwatch-log-level.html) in the AWS Step Functions User Guide.
	LoggingConfiguration StateMachineLoggingConfigurationPtrInput
	// The name of the state machine.
	Name pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the IAM role to use for this state machine.
	RoleArn pulumi.StringPtrInput
	// The current status of the state machine. Either "ACTIVE" or "DELETING".
	Status pulumi.StringPtrInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
	// Determines whether a Standard or Express state machine is created. The default is STANDARD. You cannot update the type of a state machine once it has been created. Valid Values: STANDARD | EXPRESS
	Type pulumi.StringPtrInput
}

func (StateMachineState) ElementType() reflect.Type {
	return reflect.TypeOf((*stateMachineState)(nil)).Elem()
}

type stateMachineArgs struct {
	// The Amazon States Language definition of the state machine.
	Definition string `pulumi:"definition"`
	// Defines what execution history events are logged and where they are logged. The `loggingConfiguration` parameter is only valid when `type` is set to `EXPRESS`. Defaults to `OFF`. For more information see [Logging Express Workflows](https://docs.aws.amazon.com/step-functions/latest/dg/cw-logs.html) and [Log Levels](https://docs.aws.amazon.com/step-functions/latest/dg/cloudwatch-log-level.html) in the AWS Step Functions User Guide.
	LoggingConfiguration *StateMachineLoggingConfiguration `pulumi:"loggingConfiguration"`
	// The name of the state machine.
	Name *string `pulumi:"name"`
	// The Amazon Resource Name (ARN) of the IAM role to use for this state machine.
	RoleArn string `pulumi:"roleArn"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
	// Determines whether a Standard or Express state machine is created. The default is STANDARD. You cannot update the type of a state machine once it has been created. Valid Values: STANDARD | EXPRESS
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a StateMachine resource.
type StateMachineArgs struct {
	// The Amazon States Language definition of the state machine.
	Definition pulumi.StringInput
	// Defines what execution history events are logged and where they are logged. The `loggingConfiguration` parameter is only valid when `type` is set to `EXPRESS`. Defaults to `OFF`. For more information see [Logging Express Workflows](https://docs.aws.amazon.com/step-functions/latest/dg/cw-logs.html) and [Log Levels](https://docs.aws.amazon.com/step-functions/latest/dg/cloudwatch-log-level.html) in the AWS Step Functions User Guide.
	LoggingConfiguration StateMachineLoggingConfigurationPtrInput
	// The name of the state machine.
	Name pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the IAM role to use for this state machine.
	RoleArn pulumi.StringInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
	// Determines whether a Standard or Express state machine is created. The default is STANDARD. You cannot update the type of a state machine once it has been created. Valid Values: STANDARD | EXPRESS
	Type pulumi.StringPtrInput
}

func (StateMachineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*stateMachineArgs)(nil)).Elem()
}

type StateMachineInput interface {
	pulumi.Input

	ToStateMachineOutput() StateMachineOutput
	ToStateMachineOutputWithContext(ctx context.Context) StateMachineOutput
}

func (*StateMachine) ElementType() reflect.Type {
	return reflect.TypeOf((*StateMachine)(nil))
}

func (i *StateMachine) ToStateMachineOutput() StateMachineOutput {
	return i.ToStateMachineOutputWithContext(context.Background())
}

func (i *StateMachine) ToStateMachineOutputWithContext(ctx context.Context) StateMachineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StateMachineOutput)
}

func (i *StateMachine) ToStateMachinePtrOutput() StateMachinePtrOutput {
	return i.ToStateMachinePtrOutputWithContext(context.Background())
}

func (i *StateMachine) ToStateMachinePtrOutputWithContext(ctx context.Context) StateMachinePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StateMachinePtrOutput)
}

type StateMachinePtrInput interface {
	pulumi.Input

	ToStateMachinePtrOutput() StateMachinePtrOutput
	ToStateMachinePtrOutputWithContext(ctx context.Context) StateMachinePtrOutput
}

type stateMachinePtrType StateMachineArgs

func (*stateMachinePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StateMachine)(nil))
}

func (i *stateMachinePtrType) ToStateMachinePtrOutput() StateMachinePtrOutput {
	return i.ToStateMachinePtrOutputWithContext(context.Background())
}

func (i *stateMachinePtrType) ToStateMachinePtrOutputWithContext(ctx context.Context) StateMachinePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StateMachinePtrOutput)
}

// StateMachineArrayInput is an input type that accepts StateMachineArray and StateMachineArrayOutput values.
// You can construct a concrete instance of `StateMachineArrayInput` via:
//
//          StateMachineArray{ StateMachineArgs{...} }
type StateMachineArrayInput interface {
	pulumi.Input

	ToStateMachineArrayOutput() StateMachineArrayOutput
	ToStateMachineArrayOutputWithContext(context.Context) StateMachineArrayOutput
}

type StateMachineArray []StateMachineInput

func (StateMachineArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*StateMachine)(nil))
}

func (i StateMachineArray) ToStateMachineArrayOutput() StateMachineArrayOutput {
	return i.ToStateMachineArrayOutputWithContext(context.Background())
}

func (i StateMachineArray) ToStateMachineArrayOutputWithContext(ctx context.Context) StateMachineArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StateMachineArrayOutput)
}

// StateMachineMapInput is an input type that accepts StateMachineMap and StateMachineMapOutput values.
// You can construct a concrete instance of `StateMachineMapInput` via:
//
//          StateMachineMap{ "key": StateMachineArgs{...} }
type StateMachineMapInput interface {
	pulumi.Input

	ToStateMachineMapOutput() StateMachineMapOutput
	ToStateMachineMapOutputWithContext(context.Context) StateMachineMapOutput
}

type StateMachineMap map[string]StateMachineInput

func (StateMachineMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*StateMachine)(nil))
}

func (i StateMachineMap) ToStateMachineMapOutput() StateMachineMapOutput {
	return i.ToStateMachineMapOutputWithContext(context.Background())
}

func (i StateMachineMap) ToStateMachineMapOutputWithContext(ctx context.Context) StateMachineMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StateMachineMapOutput)
}

type StateMachineOutput struct {
	*pulumi.OutputState
}

func (StateMachineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StateMachine)(nil))
}

func (o StateMachineOutput) ToStateMachineOutput() StateMachineOutput {
	return o
}

func (o StateMachineOutput) ToStateMachineOutputWithContext(ctx context.Context) StateMachineOutput {
	return o
}

func (o StateMachineOutput) ToStateMachinePtrOutput() StateMachinePtrOutput {
	return o.ToStateMachinePtrOutputWithContext(context.Background())
}

func (o StateMachineOutput) ToStateMachinePtrOutputWithContext(ctx context.Context) StateMachinePtrOutput {
	return o.ApplyT(func(v StateMachine) *StateMachine {
		return &v
	}).(StateMachinePtrOutput)
}

type StateMachinePtrOutput struct {
	*pulumi.OutputState
}

func (StateMachinePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StateMachine)(nil))
}

func (o StateMachinePtrOutput) ToStateMachinePtrOutput() StateMachinePtrOutput {
	return o
}

func (o StateMachinePtrOutput) ToStateMachinePtrOutputWithContext(ctx context.Context) StateMachinePtrOutput {
	return o
}

type StateMachineArrayOutput struct{ *pulumi.OutputState }

func (StateMachineArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StateMachine)(nil))
}

func (o StateMachineArrayOutput) ToStateMachineArrayOutput() StateMachineArrayOutput {
	return o
}

func (o StateMachineArrayOutput) ToStateMachineArrayOutputWithContext(ctx context.Context) StateMachineArrayOutput {
	return o
}

func (o StateMachineArrayOutput) Index(i pulumi.IntInput) StateMachineOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StateMachine {
		return vs[0].([]StateMachine)[vs[1].(int)]
	}).(StateMachineOutput)
}

type StateMachineMapOutput struct{ *pulumi.OutputState }

func (StateMachineMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]StateMachine)(nil))
}

func (o StateMachineMapOutput) ToStateMachineMapOutput() StateMachineMapOutput {
	return o
}

func (o StateMachineMapOutput) ToStateMachineMapOutputWithContext(ctx context.Context) StateMachineMapOutput {
	return o
}

func (o StateMachineMapOutput) MapIndex(k pulumi.StringInput) StateMachineOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) StateMachine {
		return vs[0].(map[string]StateMachine)[vs[1].(string)]
	}).(StateMachineOutput)
}

func init() {
	pulumi.RegisterOutputType(StateMachineOutput{})
	pulumi.RegisterOutputType(StateMachinePtrOutput{})
	pulumi.RegisterOutputType(StateMachineArrayOutput{})
	pulumi.RegisterOutputType(StateMachineMapOutput{})
}
