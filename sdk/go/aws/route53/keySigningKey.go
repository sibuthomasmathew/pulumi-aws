// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Route 53 Key Signing Key. To manage Domain Name System Security Extensions (DNSSEC) for a Hosted Zone, see the `route53.HostedZoneDnsSec` resource. For more information about managing DNSSEC in Route 53, see the [Route 53 Developer Guide](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-configuring-dnssec.html).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"encoding/json"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/kms"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/route53"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		tmpJSON0, err := json.Marshal(map[string]interface{}{
// 			"Statement": []interface{}{
// 				map[string]interface{}{
// 					"Action": []string{
// 						"kms:DescribeKey",
// 						"kms:GetPublicKey",
// 						"kms:Sign",
// 					},
// 					"Effect": "Allow",
// 					"Principal": map[string]interface{}{
// 						"Service": "api-service.dnssec.route53.aws.internal",
// 					},
// 					"Sid": "Route 53 DNSSEC Permissions",
// 				},
// 				map[string]interface{}{
// 					"Action": "kms:*",
// 					"Effect": "Allow",
// 					"Principal": map[string]interface{}{
// 						"AWS": "*",
// 					},
// 					"Resource": "*",
// 					"Sid":      "IAM User Permissions",
// 				},
// 			},
// 			"Version": "2012-10-17",
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		json0 := string(tmpJSON0)
// 		_, err := kms.NewKey(ctx, "exampleKey", &kms.KeyArgs{
// 			CustomerMasterKeySpec: pulumi.String("ECC_NIST_P256"),
// 			DeletionWindowInDays:  pulumi.Int(7),
// 			KeyUsage:              pulumi.String("SIGN_VERIFY"),
// 			Policy:                pulumi.String(json0),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = route53.NewZone(ctx, "exampleZone", nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleKeySigningKey, err := route53.NewKeySigningKey(ctx, "exampleKeySigningKey", &route53.KeySigningKeyArgs{
// 			HostedZoneId:            pulumi.Any(aws_route53_zone.Test.Id),
// 			KeyManagementServiceArn: pulumi.Any(aws_kms_key.Test.Arn),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = route53.NewHostedZoneDnsSec(ctx, "exampleHostedZoneDnsSec", &route53.HostedZoneDnsSecArgs{
// 			HostedZoneId: exampleKeySigningKey.HostedZoneId,
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
// `aws_route53_key_signing_key` resources can be imported by using the Route 53 Hosted Zone identifier and KMS Key identifier, separated by a comma (`,`), e.g.
//
// ```sh
//  $ pulumi import aws:route53/keySigningKey:KeySigningKey example Z1D633PJN98FT9,example
// ```
type KeySigningKey struct {
	pulumi.CustomResourceState

	// A string used to represent the delegation signer digest algorithm. This value must follow the guidelines provided by [RFC-8624 Section 3.3](https://tools.ietf.org/html/rfc8624#section-3.3).
	DigestAlgorithmMnemonic pulumi.StringOutput `pulumi:"digestAlgorithmMnemonic"`
	// An integer used to represent the delegation signer digest algorithm. This value must follow the guidelines provided by [RFC-8624 Section 3.3](https://tools.ietf.org/html/rfc8624#section-3.3).
	DigestAlgorithmType pulumi.IntOutput `pulumi:"digestAlgorithmType"`
	// A cryptographic digest of a DNSKEY resource record (RR). DNSKEY records are used to publish the public key that resolvers can use to verify DNSSEC signatures that are used to secure certain kinds of information provided by the DNS system.
	DigestValue pulumi.StringOutput `pulumi:"digestValue"`
	// A string that represents a DNSKEY record.
	DnskeyRecord pulumi.StringOutput `pulumi:"dnskeyRecord"`
	// A string that represents a delegation signer (DS) record.
	DsRecord pulumi.StringOutput `pulumi:"dsRecord"`
	// An integer that specifies how the key is used. For key-signing key (KSK), this value is always 257.
	Flag pulumi.IntOutput `pulumi:"flag"`
	// Identifier of the Route 53 Hosted Zone.
	HostedZoneId pulumi.StringOutput `pulumi:"hostedZoneId"`
	// Amazon Resource Name (ARN) of the Key Management Service (KMS) Key. This must be unique for each key-signing key (KSK) in a single hosted zone. This key must be in the `us-east-1` Region and meet certain requirements, which are described in the [Route 53 Developer Guide](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-configuring-dnssec-cmk-requirements.html) and [Route 53 API Reference](https://docs.aws.amazon.com/Route53/latest/APIReference/API_CreateKeySigningKey.html).
	KeyManagementServiceArn pulumi.StringOutput `pulumi:"keyManagementServiceArn"`
	// An integer used to identify the DNSSEC record for the domain name. The process used to calculate the value is described in [RFC-4034 Appendix B](https://tools.ietf.org/rfc/rfc4034.txt).
	KeyTag pulumi.IntOutput `pulumi:"keyTag"`
	// Name of the key-signing key (KSK). Must be unique for each key-singing key in the same hosted zone.
	Name pulumi.StringOutput `pulumi:"name"`
	// The public key, represented as a Base64 encoding, as required by [RFC-4034 Page 5](https://tools.ietf.org/rfc/rfc4034.txt).
	PublicKey pulumi.StringOutput `pulumi:"publicKey"`
	// A string used to represent the signing algorithm. This value must follow the guidelines provided by [RFC-8624 Section 3.1](https://tools.ietf.org/html/rfc8624#section-3.1).
	SigningAlgorithmMnemonic pulumi.StringOutput `pulumi:"signingAlgorithmMnemonic"`
	// An integer used to represent the signing algorithm. This value must follow the guidelines provided by [RFC-8624 Section 3.1](https://tools.ietf.org/html/rfc8624#section-3.1).
	SigningAlgorithmType pulumi.IntOutput `pulumi:"signingAlgorithmType"`
	// Status of the key-signing key (KSK). Valid values: `ACTIVE`, `INACTIVE`. Defaults to `ACTIVE`.
	Status pulumi.StringPtrOutput `pulumi:"status"`
}

// NewKeySigningKey registers a new resource with the given unique name, arguments, and options.
func NewKeySigningKey(ctx *pulumi.Context,
	name string, args *KeySigningKeyArgs, opts ...pulumi.ResourceOption) (*KeySigningKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HostedZoneId == nil {
		return nil, errors.New("invalid value for required argument 'HostedZoneId'")
	}
	if args.KeyManagementServiceArn == nil {
		return nil, errors.New("invalid value for required argument 'KeyManagementServiceArn'")
	}
	var resource KeySigningKey
	err := ctx.RegisterResource("aws:route53/keySigningKey:KeySigningKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKeySigningKey gets an existing KeySigningKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKeySigningKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeySigningKeyState, opts ...pulumi.ResourceOption) (*KeySigningKey, error) {
	var resource KeySigningKey
	err := ctx.ReadResource("aws:route53/keySigningKey:KeySigningKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KeySigningKey resources.
type keySigningKeyState struct {
	// A string used to represent the delegation signer digest algorithm. This value must follow the guidelines provided by [RFC-8624 Section 3.3](https://tools.ietf.org/html/rfc8624#section-3.3).
	DigestAlgorithmMnemonic *string `pulumi:"digestAlgorithmMnemonic"`
	// An integer used to represent the delegation signer digest algorithm. This value must follow the guidelines provided by [RFC-8624 Section 3.3](https://tools.ietf.org/html/rfc8624#section-3.3).
	DigestAlgorithmType *int `pulumi:"digestAlgorithmType"`
	// A cryptographic digest of a DNSKEY resource record (RR). DNSKEY records are used to publish the public key that resolvers can use to verify DNSSEC signatures that are used to secure certain kinds of information provided by the DNS system.
	DigestValue *string `pulumi:"digestValue"`
	// A string that represents a DNSKEY record.
	DnskeyRecord *string `pulumi:"dnskeyRecord"`
	// A string that represents a delegation signer (DS) record.
	DsRecord *string `pulumi:"dsRecord"`
	// An integer that specifies how the key is used. For key-signing key (KSK), this value is always 257.
	Flag *int `pulumi:"flag"`
	// Identifier of the Route 53 Hosted Zone.
	HostedZoneId *string `pulumi:"hostedZoneId"`
	// Amazon Resource Name (ARN) of the Key Management Service (KMS) Key. This must be unique for each key-signing key (KSK) in a single hosted zone. This key must be in the `us-east-1` Region and meet certain requirements, which are described in the [Route 53 Developer Guide](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-configuring-dnssec-cmk-requirements.html) and [Route 53 API Reference](https://docs.aws.amazon.com/Route53/latest/APIReference/API_CreateKeySigningKey.html).
	KeyManagementServiceArn *string `pulumi:"keyManagementServiceArn"`
	// An integer used to identify the DNSSEC record for the domain name. The process used to calculate the value is described in [RFC-4034 Appendix B](https://tools.ietf.org/rfc/rfc4034.txt).
	KeyTag *int `pulumi:"keyTag"`
	// Name of the key-signing key (KSK). Must be unique for each key-singing key in the same hosted zone.
	Name *string `pulumi:"name"`
	// The public key, represented as a Base64 encoding, as required by [RFC-4034 Page 5](https://tools.ietf.org/rfc/rfc4034.txt).
	PublicKey *string `pulumi:"publicKey"`
	// A string used to represent the signing algorithm. This value must follow the guidelines provided by [RFC-8624 Section 3.1](https://tools.ietf.org/html/rfc8624#section-3.1).
	SigningAlgorithmMnemonic *string `pulumi:"signingAlgorithmMnemonic"`
	// An integer used to represent the signing algorithm. This value must follow the guidelines provided by [RFC-8624 Section 3.1](https://tools.ietf.org/html/rfc8624#section-3.1).
	SigningAlgorithmType *int `pulumi:"signingAlgorithmType"`
	// Status of the key-signing key (KSK). Valid values: `ACTIVE`, `INACTIVE`. Defaults to `ACTIVE`.
	Status *string `pulumi:"status"`
}

type KeySigningKeyState struct {
	// A string used to represent the delegation signer digest algorithm. This value must follow the guidelines provided by [RFC-8624 Section 3.3](https://tools.ietf.org/html/rfc8624#section-3.3).
	DigestAlgorithmMnemonic pulumi.StringPtrInput
	// An integer used to represent the delegation signer digest algorithm. This value must follow the guidelines provided by [RFC-8624 Section 3.3](https://tools.ietf.org/html/rfc8624#section-3.3).
	DigestAlgorithmType pulumi.IntPtrInput
	// A cryptographic digest of a DNSKEY resource record (RR). DNSKEY records are used to publish the public key that resolvers can use to verify DNSSEC signatures that are used to secure certain kinds of information provided by the DNS system.
	DigestValue pulumi.StringPtrInput
	// A string that represents a DNSKEY record.
	DnskeyRecord pulumi.StringPtrInput
	// A string that represents a delegation signer (DS) record.
	DsRecord pulumi.StringPtrInput
	// An integer that specifies how the key is used. For key-signing key (KSK), this value is always 257.
	Flag pulumi.IntPtrInput
	// Identifier of the Route 53 Hosted Zone.
	HostedZoneId pulumi.StringPtrInput
	// Amazon Resource Name (ARN) of the Key Management Service (KMS) Key. This must be unique for each key-signing key (KSK) in a single hosted zone. This key must be in the `us-east-1` Region and meet certain requirements, which are described in the [Route 53 Developer Guide](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-configuring-dnssec-cmk-requirements.html) and [Route 53 API Reference](https://docs.aws.amazon.com/Route53/latest/APIReference/API_CreateKeySigningKey.html).
	KeyManagementServiceArn pulumi.StringPtrInput
	// An integer used to identify the DNSSEC record for the domain name. The process used to calculate the value is described in [RFC-4034 Appendix B](https://tools.ietf.org/rfc/rfc4034.txt).
	KeyTag pulumi.IntPtrInput
	// Name of the key-signing key (KSK). Must be unique for each key-singing key in the same hosted zone.
	Name pulumi.StringPtrInput
	// The public key, represented as a Base64 encoding, as required by [RFC-4034 Page 5](https://tools.ietf.org/rfc/rfc4034.txt).
	PublicKey pulumi.StringPtrInput
	// A string used to represent the signing algorithm. This value must follow the guidelines provided by [RFC-8624 Section 3.1](https://tools.ietf.org/html/rfc8624#section-3.1).
	SigningAlgorithmMnemonic pulumi.StringPtrInput
	// An integer used to represent the signing algorithm. This value must follow the guidelines provided by [RFC-8624 Section 3.1](https://tools.ietf.org/html/rfc8624#section-3.1).
	SigningAlgorithmType pulumi.IntPtrInput
	// Status of the key-signing key (KSK). Valid values: `ACTIVE`, `INACTIVE`. Defaults to `ACTIVE`.
	Status pulumi.StringPtrInput
}

func (KeySigningKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*keySigningKeyState)(nil)).Elem()
}

type keySigningKeyArgs struct {
	// Identifier of the Route 53 Hosted Zone.
	HostedZoneId string `pulumi:"hostedZoneId"`
	// Amazon Resource Name (ARN) of the Key Management Service (KMS) Key. This must be unique for each key-signing key (KSK) in a single hosted zone. This key must be in the `us-east-1` Region and meet certain requirements, which are described in the [Route 53 Developer Guide](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-configuring-dnssec-cmk-requirements.html) and [Route 53 API Reference](https://docs.aws.amazon.com/Route53/latest/APIReference/API_CreateKeySigningKey.html).
	KeyManagementServiceArn string `pulumi:"keyManagementServiceArn"`
	// Name of the key-signing key (KSK). Must be unique for each key-singing key in the same hosted zone.
	Name *string `pulumi:"name"`
	// Status of the key-signing key (KSK). Valid values: `ACTIVE`, `INACTIVE`. Defaults to `ACTIVE`.
	Status *string `pulumi:"status"`
}

// The set of arguments for constructing a KeySigningKey resource.
type KeySigningKeyArgs struct {
	// Identifier of the Route 53 Hosted Zone.
	HostedZoneId pulumi.StringInput
	// Amazon Resource Name (ARN) of the Key Management Service (KMS) Key. This must be unique for each key-signing key (KSK) in a single hosted zone. This key must be in the `us-east-1` Region and meet certain requirements, which are described in the [Route 53 Developer Guide](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-configuring-dnssec-cmk-requirements.html) and [Route 53 API Reference](https://docs.aws.amazon.com/Route53/latest/APIReference/API_CreateKeySigningKey.html).
	KeyManagementServiceArn pulumi.StringInput
	// Name of the key-signing key (KSK). Must be unique for each key-singing key in the same hosted zone.
	Name pulumi.StringPtrInput
	// Status of the key-signing key (KSK). Valid values: `ACTIVE`, `INACTIVE`. Defaults to `ACTIVE`.
	Status pulumi.StringPtrInput
}

func (KeySigningKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keySigningKeyArgs)(nil)).Elem()
}

type KeySigningKeyInput interface {
	pulumi.Input

	ToKeySigningKeyOutput() KeySigningKeyOutput
	ToKeySigningKeyOutputWithContext(ctx context.Context) KeySigningKeyOutput
}

func (*KeySigningKey) ElementType() reflect.Type {
	return reflect.TypeOf((*KeySigningKey)(nil))
}

func (i *KeySigningKey) ToKeySigningKeyOutput() KeySigningKeyOutput {
	return i.ToKeySigningKeyOutputWithContext(context.Background())
}

func (i *KeySigningKey) ToKeySigningKeyOutputWithContext(ctx context.Context) KeySigningKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeySigningKeyOutput)
}

func (i *KeySigningKey) ToKeySigningKeyPtrOutput() KeySigningKeyPtrOutput {
	return i.ToKeySigningKeyPtrOutputWithContext(context.Background())
}

func (i *KeySigningKey) ToKeySigningKeyPtrOutputWithContext(ctx context.Context) KeySigningKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeySigningKeyPtrOutput)
}

type KeySigningKeyPtrInput interface {
	pulumi.Input

	ToKeySigningKeyPtrOutput() KeySigningKeyPtrOutput
	ToKeySigningKeyPtrOutputWithContext(ctx context.Context) KeySigningKeyPtrOutput
}

type keySigningKeyPtrType KeySigningKeyArgs

func (*keySigningKeyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeySigningKey)(nil))
}

func (i *keySigningKeyPtrType) ToKeySigningKeyPtrOutput() KeySigningKeyPtrOutput {
	return i.ToKeySigningKeyPtrOutputWithContext(context.Background())
}

func (i *keySigningKeyPtrType) ToKeySigningKeyPtrOutputWithContext(ctx context.Context) KeySigningKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeySigningKeyPtrOutput)
}

// KeySigningKeyArrayInput is an input type that accepts KeySigningKeyArray and KeySigningKeyArrayOutput values.
// You can construct a concrete instance of `KeySigningKeyArrayInput` via:
//
//          KeySigningKeyArray{ KeySigningKeyArgs{...} }
type KeySigningKeyArrayInput interface {
	pulumi.Input

	ToKeySigningKeyArrayOutput() KeySigningKeyArrayOutput
	ToKeySigningKeyArrayOutputWithContext(context.Context) KeySigningKeyArrayOutput
}

type KeySigningKeyArray []KeySigningKeyInput

func (KeySigningKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*KeySigningKey)(nil))
}

func (i KeySigningKeyArray) ToKeySigningKeyArrayOutput() KeySigningKeyArrayOutput {
	return i.ToKeySigningKeyArrayOutputWithContext(context.Background())
}

func (i KeySigningKeyArray) ToKeySigningKeyArrayOutputWithContext(ctx context.Context) KeySigningKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeySigningKeyArrayOutput)
}

// KeySigningKeyMapInput is an input type that accepts KeySigningKeyMap and KeySigningKeyMapOutput values.
// You can construct a concrete instance of `KeySigningKeyMapInput` via:
//
//          KeySigningKeyMap{ "key": KeySigningKeyArgs{...} }
type KeySigningKeyMapInput interface {
	pulumi.Input

	ToKeySigningKeyMapOutput() KeySigningKeyMapOutput
	ToKeySigningKeyMapOutputWithContext(context.Context) KeySigningKeyMapOutput
}

type KeySigningKeyMap map[string]KeySigningKeyInput

func (KeySigningKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*KeySigningKey)(nil))
}

func (i KeySigningKeyMap) ToKeySigningKeyMapOutput() KeySigningKeyMapOutput {
	return i.ToKeySigningKeyMapOutputWithContext(context.Background())
}

func (i KeySigningKeyMap) ToKeySigningKeyMapOutputWithContext(ctx context.Context) KeySigningKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeySigningKeyMapOutput)
}

type KeySigningKeyOutput struct {
	*pulumi.OutputState
}

func (KeySigningKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeySigningKey)(nil))
}

func (o KeySigningKeyOutput) ToKeySigningKeyOutput() KeySigningKeyOutput {
	return o
}

func (o KeySigningKeyOutput) ToKeySigningKeyOutputWithContext(ctx context.Context) KeySigningKeyOutput {
	return o
}

func (o KeySigningKeyOutput) ToKeySigningKeyPtrOutput() KeySigningKeyPtrOutput {
	return o.ToKeySigningKeyPtrOutputWithContext(context.Background())
}

func (o KeySigningKeyOutput) ToKeySigningKeyPtrOutputWithContext(ctx context.Context) KeySigningKeyPtrOutput {
	return o.ApplyT(func(v KeySigningKey) *KeySigningKey {
		return &v
	}).(KeySigningKeyPtrOutput)
}

type KeySigningKeyPtrOutput struct {
	*pulumi.OutputState
}

func (KeySigningKeyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeySigningKey)(nil))
}

func (o KeySigningKeyPtrOutput) ToKeySigningKeyPtrOutput() KeySigningKeyPtrOutput {
	return o
}

func (o KeySigningKeyPtrOutput) ToKeySigningKeyPtrOutputWithContext(ctx context.Context) KeySigningKeyPtrOutput {
	return o
}

type KeySigningKeyArrayOutput struct{ *pulumi.OutputState }

func (KeySigningKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeySigningKey)(nil))
}

func (o KeySigningKeyArrayOutput) ToKeySigningKeyArrayOutput() KeySigningKeyArrayOutput {
	return o
}

func (o KeySigningKeyArrayOutput) ToKeySigningKeyArrayOutputWithContext(ctx context.Context) KeySigningKeyArrayOutput {
	return o
}

func (o KeySigningKeyArrayOutput) Index(i pulumi.IntInput) KeySigningKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KeySigningKey {
		return vs[0].([]KeySigningKey)[vs[1].(int)]
	}).(KeySigningKeyOutput)
}

type KeySigningKeyMapOutput struct{ *pulumi.OutputState }

func (KeySigningKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]KeySigningKey)(nil))
}

func (o KeySigningKeyMapOutput) ToKeySigningKeyMapOutput() KeySigningKeyMapOutput {
	return o
}

func (o KeySigningKeyMapOutput) ToKeySigningKeyMapOutputWithContext(ctx context.Context) KeySigningKeyMapOutput {
	return o
}

func (o KeySigningKeyMapOutput) MapIndex(k pulumi.StringInput) KeySigningKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) KeySigningKey {
		return vs[0].(map[string]KeySigningKey)[vs[1].(string)]
	}).(KeySigningKeyOutput)
}

func init() {
	pulumi.RegisterOutputType(KeySigningKeyOutput{})
	pulumi.RegisterOutputType(KeySigningKeyPtrOutput{})
	pulumi.RegisterOutputType(KeySigningKeyArrayOutput{})
	pulumi.RegisterOutputType(KeySigningKeyMapOutput{})
}
