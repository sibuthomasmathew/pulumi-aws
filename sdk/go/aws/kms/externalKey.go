// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a KMS Customer Master Key that uses external key material. To instead manage a KMS Customer Master Key where AWS automatically generates and potentially rotates key material, see the `kms.Key` resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/kms"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := kms.NewExternalKey(ctx, "example", &kms.ExternalKeyArgs{
// 			Description: pulumi.String("KMS EXTERNAL for AMI encryption"),
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
// KMS External Keys can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import aws:kms/externalKey:ExternalKey a arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
// ```
type ExternalKey struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the key.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Duration in days after which the key is deleted after destruction of the resource. Must be between `7` and `30` days. Defaults to `30`.
	DeletionWindowInDays pulumi.IntPtrOutput `pulumi:"deletionWindowInDays"`
	// Description of the key.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies whether the key is enabled. Keys pending import can only be `false`. Imported keys default to `true` unless expired.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// Whether the key material expires. Empty when pending key material import, otherwise `KEY_MATERIAL_EXPIRES` or `KEY_MATERIAL_DOES_NOT_EXPIRE`.
	ExpirationModel pulumi.StringOutput `pulumi:"expirationModel"`
	// Base64 encoded 256-bit symmetric encryption key material to import. The CMK is permanently associated with this key material. The same key material can be reimported, but you cannot import different key material.
	KeyMaterialBase64 pulumi.StringPtrOutput `pulumi:"keyMaterialBase64"`
	// The state of the CMK.
	KeyState pulumi.StringOutput `pulumi:"keyState"`
	// The cryptographic operations for which you can use the CMK.
	KeyUsage pulumi.StringOutput `pulumi:"keyUsage"`
	// A key policy JSON document. If you do not provide a key policy, AWS KMS attaches a default key policy to the CMK.
	Policy pulumi.StringOutput `pulumi:"policy"`
	// A key-value map of tags to assign to the key.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Time at which the imported key material expires. When the key material expires, AWS KMS deletes the key material and the CMK becomes unusable. If not specified, key material does not expire. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
	ValidTo pulumi.StringPtrOutput `pulumi:"validTo"`
}

// NewExternalKey registers a new resource with the given unique name, arguments, and options.
func NewExternalKey(ctx *pulumi.Context,
	name string, args *ExternalKeyArgs, opts ...pulumi.ResourceOption) (*ExternalKey, error) {
	if args == nil {
		args = &ExternalKeyArgs{}
	}

	var resource ExternalKey
	err := ctx.RegisterResource("aws:kms/externalKey:ExternalKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExternalKey gets an existing ExternalKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExternalKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExternalKeyState, opts ...pulumi.ResourceOption) (*ExternalKey, error) {
	var resource ExternalKey
	err := ctx.ReadResource("aws:kms/externalKey:ExternalKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExternalKey resources.
type externalKeyState struct {
	// The Amazon Resource Name (ARN) of the key.
	Arn *string `pulumi:"arn"`
	// Duration in days after which the key is deleted after destruction of the resource. Must be between `7` and `30` days. Defaults to `30`.
	DeletionWindowInDays *int `pulumi:"deletionWindowInDays"`
	// Description of the key.
	Description *string `pulumi:"description"`
	// Specifies whether the key is enabled. Keys pending import can only be `false`. Imported keys default to `true` unless expired.
	Enabled *bool `pulumi:"enabled"`
	// Whether the key material expires. Empty when pending key material import, otherwise `KEY_MATERIAL_EXPIRES` or `KEY_MATERIAL_DOES_NOT_EXPIRE`.
	ExpirationModel *string `pulumi:"expirationModel"`
	// Base64 encoded 256-bit symmetric encryption key material to import. The CMK is permanently associated with this key material. The same key material can be reimported, but you cannot import different key material.
	KeyMaterialBase64 *string `pulumi:"keyMaterialBase64"`
	// The state of the CMK.
	KeyState *string `pulumi:"keyState"`
	// The cryptographic operations for which you can use the CMK.
	KeyUsage *string `pulumi:"keyUsage"`
	// A key policy JSON document. If you do not provide a key policy, AWS KMS attaches a default key policy to the CMK.
	Policy *string `pulumi:"policy"`
	// A key-value map of tags to assign to the key.
	Tags map[string]string `pulumi:"tags"`
	// Time at which the imported key material expires. When the key material expires, AWS KMS deletes the key material and the CMK becomes unusable. If not specified, key material does not expire. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
	ValidTo *string `pulumi:"validTo"`
}

type ExternalKeyState struct {
	// The Amazon Resource Name (ARN) of the key.
	Arn pulumi.StringPtrInput
	// Duration in days after which the key is deleted after destruction of the resource. Must be between `7` and `30` days. Defaults to `30`.
	DeletionWindowInDays pulumi.IntPtrInput
	// Description of the key.
	Description pulumi.StringPtrInput
	// Specifies whether the key is enabled. Keys pending import can only be `false`. Imported keys default to `true` unless expired.
	Enabled pulumi.BoolPtrInput
	// Whether the key material expires. Empty when pending key material import, otherwise `KEY_MATERIAL_EXPIRES` or `KEY_MATERIAL_DOES_NOT_EXPIRE`.
	ExpirationModel pulumi.StringPtrInput
	// Base64 encoded 256-bit symmetric encryption key material to import. The CMK is permanently associated with this key material. The same key material can be reimported, but you cannot import different key material.
	KeyMaterialBase64 pulumi.StringPtrInput
	// The state of the CMK.
	KeyState pulumi.StringPtrInput
	// The cryptographic operations for which you can use the CMK.
	KeyUsage pulumi.StringPtrInput
	// A key policy JSON document. If you do not provide a key policy, AWS KMS attaches a default key policy to the CMK.
	Policy pulumi.StringPtrInput
	// A key-value map of tags to assign to the key.
	Tags pulumi.StringMapInput
	// Time at which the imported key material expires. When the key material expires, AWS KMS deletes the key material and the CMK becomes unusable. If not specified, key material does not expire. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
	ValidTo pulumi.StringPtrInput
}

func (ExternalKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*externalKeyState)(nil)).Elem()
}

type externalKeyArgs struct {
	// Duration in days after which the key is deleted after destruction of the resource. Must be between `7` and `30` days. Defaults to `30`.
	DeletionWindowInDays *int `pulumi:"deletionWindowInDays"`
	// Description of the key.
	Description *string `pulumi:"description"`
	// Specifies whether the key is enabled. Keys pending import can only be `false`. Imported keys default to `true` unless expired.
	Enabled *bool `pulumi:"enabled"`
	// Base64 encoded 256-bit symmetric encryption key material to import. The CMK is permanently associated with this key material. The same key material can be reimported, but you cannot import different key material.
	KeyMaterialBase64 *string `pulumi:"keyMaterialBase64"`
	// A key policy JSON document. If you do not provide a key policy, AWS KMS attaches a default key policy to the CMK.
	Policy *string `pulumi:"policy"`
	// A key-value map of tags to assign to the key.
	Tags map[string]string `pulumi:"tags"`
	// Time at which the imported key material expires. When the key material expires, AWS KMS deletes the key material and the CMK becomes unusable. If not specified, key material does not expire. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
	ValidTo *string `pulumi:"validTo"`
}

// The set of arguments for constructing a ExternalKey resource.
type ExternalKeyArgs struct {
	// Duration in days after which the key is deleted after destruction of the resource. Must be between `7` and `30` days. Defaults to `30`.
	DeletionWindowInDays pulumi.IntPtrInput
	// Description of the key.
	Description pulumi.StringPtrInput
	// Specifies whether the key is enabled. Keys pending import can only be `false`. Imported keys default to `true` unless expired.
	Enabled pulumi.BoolPtrInput
	// Base64 encoded 256-bit symmetric encryption key material to import. The CMK is permanently associated with this key material. The same key material can be reimported, but you cannot import different key material.
	KeyMaterialBase64 pulumi.StringPtrInput
	// A key policy JSON document. If you do not provide a key policy, AWS KMS attaches a default key policy to the CMK.
	Policy pulumi.StringPtrInput
	// A key-value map of tags to assign to the key.
	Tags pulumi.StringMapInput
	// Time at which the imported key material expires. When the key material expires, AWS KMS deletes the key material and the CMK becomes unusable. If not specified, key material does not expire. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
	ValidTo pulumi.StringPtrInput
}

func (ExternalKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*externalKeyArgs)(nil)).Elem()
}

type ExternalKeyInput interface {
	pulumi.Input

	ToExternalKeyOutput() ExternalKeyOutput
	ToExternalKeyOutputWithContext(ctx context.Context) ExternalKeyOutput
}

func (*ExternalKey) ElementType() reflect.Type {
	return reflect.TypeOf((*ExternalKey)(nil))
}

func (i *ExternalKey) ToExternalKeyOutput() ExternalKeyOutput {
	return i.ToExternalKeyOutputWithContext(context.Background())
}

func (i *ExternalKey) ToExternalKeyOutputWithContext(ctx context.Context) ExternalKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalKeyOutput)
}

func (i *ExternalKey) ToExternalKeyPtrOutput() ExternalKeyPtrOutput {
	return i.ToExternalKeyPtrOutputWithContext(context.Background())
}

func (i *ExternalKey) ToExternalKeyPtrOutputWithContext(ctx context.Context) ExternalKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalKeyPtrOutput)
}

type ExternalKeyPtrInput interface {
	pulumi.Input

	ToExternalKeyPtrOutput() ExternalKeyPtrOutput
	ToExternalKeyPtrOutputWithContext(ctx context.Context) ExternalKeyPtrOutput
}

type externalKeyPtrType ExternalKeyArgs

func (*externalKeyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExternalKey)(nil))
}

func (i *externalKeyPtrType) ToExternalKeyPtrOutput() ExternalKeyPtrOutput {
	return i.ToExternalKeyPtrOutputWithContext(context.Background())
}

func (i *externalKeyPtrType) ToExternalKeyPtrOutputWithContext(ctx context.Context) ExternalKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalKeyPtrOutput)
}

// ExternalKeyArrayInput is an input type that accepts ExternalKeyArray and ExternalKeyArrayOutput values.
// You can construct a concrete instance of `ExternalKeyArrayInput` via:
//
//          ExternalKeyArray{ ExternalKeyArgs{...} }
type ExternalKeyArrayInput interface {
	pulumi.Input

	ToExternalKeyArrayOutput() ExternalKeyArrayOutput
	ToExternalKeyArrayOutputWithContext(context.Context) ExternalKeyArrayOutput
}

type ExternalKeyArray []ExternalKeyInput

func (ExternalKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ExternalKey)(nil))
}

func (i ExternalKeyArray) ToExternalKeyArrayOutput() ExternalKeyArrayOutput {
	return i.ToExternalKeyArrayOutputWithContext(context.Background())
}

func (i ExternalKeyArray) ToExternalKeyArrayOutputWithContext(ctx context.Context) ExternalKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalKeyArrayOutput)
}

// ExternalKeyMapInput is an input type that accepts ExternalKeyMap and ExternalKeyMapOutput values.
// You can construct a concrete instance of `ExternalKeyMapInput` via:
//
//          ExternalKeyMap{ "key": ExternalKeyArgs{...} }
type ExternalKeyMapInput interface {
	pulumi.Input

	ToExternalKeyMapOutput() ExternalKeyMapOutput
	ToExternalKeyMapOutputWithContext(context.Context) ExternalKeyMapOutput
}

type ExternalKeyMap map[string]ExternalKeyInput

func (ExternalKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ExternalKey)(nil))
}

func (i ExternalKeyMap) ToExternalKeyMapOutput() ExternalKeyMapOutput {
	return i.ToExternalKeyMapOutputWithContext(context.Background())
}

func (i ExternalKeyMap) ToExternalKeyMapOutputWithContext(ctx context.Context) ExternalKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalKeyMapOutput)
}

type ExternalKeyOutput struct {
	*pulumi.OutputState
}

func (ExternalKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExternalKey)(nil))
}

func (o ExternalKeyOutput) ToExternalKeyOutput() ExternalKeyOutput {
	return o
}

func (o ExternalKeyOutput) ToExternalKeyOutputWithContext(ctx context.Context) ExternalKeyOutput {
	return o
}

func (o ExternalKeyOutput) ToExternalKeyPtrOutput() ExternalKeyPtrOutput {
	return o.ToExternalKeyPtrOutputWithContext(context.Background())
}

func (o ExternalKeyOutput) ToExternalKeyPtrOutputWithContext(ctx context.Context) ExternalKeyPtrOutput {
	return o.ApplyT(func(v ExternalKey) *ExternalKey {
		return &v
	}).(ExternalKeyPtrOutput)
}

type ExternalKeyPtrOutput struct {
	*pulumi.OutputState
}

func (ExternalKeyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExternalKey)(nil))
}

func (o ExternalKeyPtrOutput) ToExternalKeyPtrOutput() ExternalKeyPtrOutput {
	return o
}

func (o ExternalKeyPtrOutput) ToExternalKeyPtrOutputWithContext(ctx context.Context) ExternalKeyPtrOutput {
	return o
}

type ExternalKeyArrayOutput struct{ *pulumi.OutputState }

func (ExternalKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExternalKey)(nil))
}

func (o ExternalKeyArrayOutput) ToExternalKeyArrayOutput() ExternalKeyArrayOutput {
	return o
}

func (o ExternalKeyArrayOutput) ToExternalKeyArrayOutputWithContext(ctx context.Context) ExternalKeyArrayOutput {
	return o
}

func (o ExternalKeyArrayOutput) Index(i pulumi.IntInput) ExternalKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExternalKey {
		return vs[0].([]ExternalKey)[vs[1].(int)]
	}).(ExternalKeyOutput)
}

type ExternalKeyMapOutput struct{ *pulumi.OutputState }

func (ExternalKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ExternalKey)(nil))
}

func (o ExternalKeyMapOutput) ToExternalKeyMapOutput() ExternalKeyMapOutput {
	return o
}

func (o ExternalKeyMapOutput) ToExternalKeyMapOutputWithContext(ctx context.Context) ExternalKeyMapOutput {
	return o
}

func (o ExternalKeyMapOutput) MapIndex(k pulumi.StringInput) ExternalKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ExternalKey {
		return vs[0].(map[string]ExternalKey)[vs[1].(string)]
	}).(ExternalKeyOutput)
}

func init() {
	pulumi.RegisterOutputType(ExternalKeyOutput{})
	pulumi.RegisterOutputType(ExternalKeyPtrOutput{})
	pulumi.RegisterOutputType(ExternalKeyArrayOutput{})
	pulumi.RegisterOutputType(ExternalKeyMapOutput{})
}
