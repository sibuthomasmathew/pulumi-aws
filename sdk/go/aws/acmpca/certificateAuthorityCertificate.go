// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package acmpca

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Associates a certificate with an AWS Certificate Manager Private Certificate Authority (ACM PCA Certificate Authority). An ACM PCA Certificate Authority is unable to issue certificates until it has a certificate associated with it. A root level ACM PCA Certificate Authority is able to self-sign its own root certificate.
//
// ## Example Usage
// ### Self-Signed Root Certificate Authority Certificate
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/acmpca"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleCertificateAuthority, err := acmpca.NewCertificateAuthority(ctx, "exampleCertificateAuthority", &acmpca.CertificateAuthorityArgs{
// 			Type: pulumi.String("ROOT"),
// 			CertificateAuthorityConfiguration: &acmpca.CertificateAuthorityCertificateAuthorityConfigurationArgs{
// 				KeyAlgorithm:     pulumi.String("RSA_4096"),
// 				SigningAlgorithm: pulumi.String("SHA512WITHRSA"),
// 				Subject: &acmpca.CertificateAuthorityCertificateAuthorityConfigurationSubjectArgs{
// 					CommonName: pulumi.String("example.com"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		current, err := aws.GetPartition(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleCertificate, err := acmpca.NewCertificate(ctx, "exampleCertificate", &acmpca.CertificateArgs{
// 			CertificateAuthorityArn:   exampleCertificateAuthority.Arn,
// 			CertificateSigningRequest: exampleCertificateAuthority.CertificateSigningRequest,
// 			SigningAlgorithm:          pulumi.String("SHA512WITHRSA"),
// 			TemplateArn:               pulumi.String(fmt.Sprintf("%v%v%v", "arn:", current.Partition, ":acm-pca:::template/RootCACertificate/V1")),
// 			Validity: &acmpca.CertificateValidityArgs{
// 				Type:  pulumi.String("YEARS"),
// 				Value: pulumi.String("1"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = acmpca.NewCertificateAuthorityCertificate(ctx, "exampleCertificateAuthorityCertificate", &acmpca.CertificateAuthorityCertificateArgs{
// 			CertificateAuthorityArn: exampleCertificateAuthority.Arn,
// 			Certificate:             exampleCertificate.Certificate,
// 			CertificateChain:        exampleCertificate.CertificateChain,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Certificate for Subordinate Certificate Authority
//
// Note that the certificate for the subordinate certificate authority must be issued by the root certificate authority using a signing request from the subordinate certificate authority.
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/acmpca"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		subordinateCertificateAuthority, err := acmpca.NewCertificateAuthority(ctx, "subordinateCertificateAuthority", &acmpca.CertificateAuthorityArgs{
// 			Type: pulumi.String("SUBORDINATE"),
// 			CertificateAuthorityConfiguration: &acmpca.CertificateAuthorityCertificateAuthorityConfigurationArgs{
// 				KeyAlgorithm:     pulumi.String("RSA_2048"),
// 				SigningAlgorithm: pulumi.String("SHA512WITHRSA"),
// 				Subject: &acmpca.CertificateAuthorityCertificateAuthorityConfigurationSubjectArgs{
// 					CommonName: pulumi.String("sub.example.com"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		rootCertificateAuthority, err := acmpca.NewCertificateAuthority(ctx, "rootCertificateAuthority", nil)
// 		if err != nil {
// 			return err
// 		}
// 		current, err := aws.GetPartition(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		subordinateCertificate, err := acmpca.NewCertificate(ctx, "subordinateCertificate", &acmpca.CertificateArgs{
// 			CertificateAuthorityArn:   rootCertificateAuthority.Arn,
// 			CertificateSigningRequest: subordinateCertificateAuthority.CertificateSigningRequest,
// 			SigningAlgorithm:          pulumi.String("SHA512WITHRSA"),
// 			TemplateArn:               pulumi.String(fmt.Sprintf("%v%v%v", "arn:", current.Partition, ":acm-pca:::template/SubordinateCACertificate_PathLen0/V1")),
// 			Validity: &acmpca.CertificateValidityArgs{
// 				Type:  pulumi.String("YEARS"),
// 				Value: pulumi.String("1"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = acmpca.NewCertificateAuthorityCertificate(ctx, "subordinateCertificateAuthorityCertificate", &acmpca.CertificateAuthorityCertificateArgs{
// 			CertificateAuthorityArn: subordinateCertificateAuthority.Arn,
// 			Certificate:             subordinateCertificate.Certificate,
// 			CertificateChain:        subordinateCertificate.CertificateChain,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = acmpca.NewCertificateAuthorityCertificate(ctx, "rootCertificateAuthorityCertificate", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = acmpca.NewCertificate(ctx, "rootCertificate", nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type CertificateAuthorityCertificate struct {
	pulumi.CustomResourceState

	// The PEM-encoded certificate for the Certificate Authority.
	Certificate pulumi.StringOutput `pulumi:"certificate"`
	// Amazon Resource Name (ARN) of the Certificate Authority.
	CertificateAuthorityArn pulumi.StringOutput `pulumi:"certificateAuthorityArn"`
	// The PEM-encoded certificate chain that includes any intermediate certificates and chains up to root CA. Required for subordinate Certificate Authorities. Not allowed for root Certificate Authorities.
	CertificateChain pulumi.StringPtrOutput `pulumi:"certificateChain"`
}

// NewCertificateAuthorityCertificate registers a new resource with the given unique name, arguments, and options.
func NewCertificateAuthorityCertificate(ctx *pulumi.Context,
	name string, args *CertificateAuthorityCertificateArgs, opts ...pulumi.ResourceOption) (*CertificateAuthorityCertificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Certificate == nil {
		return nil, errors.New("invalid value for required argument 'Certificate'")
	}
	if args.CertificateAuthorityArn == nil {
		return nil, errors.New("invalid value for required argument 'CertificateAuthorityArn'")
	}
	var resource CertificateAuthorityCertificate
	err := ctx.RegisterResource("aws:acmpca/certificateAuthorityCertificate:CertificateAuthorityCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificateAuthorityCertificate gets an existing CertificateAuthorityCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificateAuthorityCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateAuthorityCertificateState, opts ...pulumi.ResourceOption) (*CertificateAuthorityCertificate, error) {
	var resource CertificateAuthorityCertificate
	err := ctx.ReadResource("aws:acmpca/certificateAuthorityCertificate:CertificateAuthorityCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CertificateAuthorityCertificate resources.
type certificateAuthorityCertificateState struct {
	// The PEM-encoded certificate for the Certificate Authority.
	Certificate *string `pulumi:"certificate"`
	// Amazon Resource Name (ARN) of the Certificate Authority.
	CertificateAuthorityArn *string `pulumi:"certificateAuthorityArn"`
	// The PEM-encoded certificate chain that includes any intermediate certificates and chains up to root CA. Required for subordinate Certificate Authorities. Not allowed for root Certificate Authorities.
	CertificateChain *string `pulumi:"certificateChain"`
}

type CertificateAuthorityCertificateState struct {
	// The PEM-encoded certificate for the Certificate Authority.
	Certificate pulumi.StringPtrInput
	// Amazon Resource Name (ARN) of the Certificate Authority.
	CertificateAuthorityArn pulumi.StringPtrInput
	// The PEM-encoded certificate chain that includes any intermediate certificates and chains up to root CA. Required for subordinate Certificate Authorities. Not allowed for root Certificate Authorities.
	CertificateChain pulumi.StringPtrInput
}

func (CertificateAuthorityCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateAuthorityCertificateState)(nil)).Elem()
}

type certificateAuthorityCertificateArgs struct {
	// The PEM-encoded certificate for the Certificate Authority.
	Certificate string `pulumi:"certificate"`
	// Amazon Resource Name (ARN) of the Certificate Authority.
	CertificateAuthorityArn string `pulumi:"certificateAuthorityArn"`
	// The PEM-encoded certificate chain that includes any intermediate certificates and chains up to root CA. Required for subordinate Certificate Authorities. Not allowed for root Certificate Authorities.
	CertificateChain *string `pulumi:"certificateChain"`
}

// The set of arguments for constructing a CertificateAuthorityCertificate resource.
type CertificateAuthorityCertificateArgs struct {
	// The PEM-encoded certificate for the Certificate Authority.
	Certificate pulumi.StringInput
	// Amazon Resource Name (ARN) of the Certificate Authority.
	CertificateAuthorityArn pulumi.StringInput
	// The PEM-encoded certificate chain that includes any intermediate certificates and chains up to root CA. Required for subordinate Certificate Authorities. Not allowed for root Certificate Authorities.
	CertificateChain pulumi.StringPtrInput
}

func (CertificateAuthorityCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateAuthorityCertificateArgs)(nil)).Elem()
}

type CertificateAuthorityCertificateInput interface {
	pulumi.Input

	ToCertificateAuthorityCertificateOutput() CertificateAuthorityCertificateOutput
	ToCertificateAuthorityCertificateOutputWithContext(ctx context.Context) CertificateAuthorityCertificateOutput
}

func (*CertificateAuthorityCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateAuthorityCertificate)(nil))
}

func (i *CertificateAuthorityCertificate) ToCertificateAuthorityCertificateOutput() CertificateAuthorityCertificateOutput {
	return i.ToCertificateAuthorityCertificateOutputWithContext(context.Background())
}

func (i *CertificateAuthorityCertificate) ToCertificateAuthorityCertificateOutputWithContext(ctx context.Context) CertificateAuthorityCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateAuthorityCertificateOutput)
}

func (i *CertificateAuthorityCertificate) ToCertificateAuthorityCertificatePtrOutput() CertificateAuthorityCertificatePtrOutput {
	return i.ToCertificateAuthorityCertificatePtrOutputWithContext(context.Background())
}

func (i *CertificateAuthorityCertificate) ToCertificateAuthorityCertificatePtrOutputWithContext(ctx context.Context) CertificateAuthorityCertificatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateAuthorityCertificatePtrOutput)
}

type CertificateAuthorityCertificatePtrInput interface {
	pulumi.Input

	ToCertificateAuthorityCertificatePtrOutput() CertificateAuthorityCertificatePtrOutput
	ToCertificateAuthorityCertificatePtrOutputWithContext(ctx context.Context) CertificateAuthorityCertificatePtrOutput
}

type certificateAuthorityCertificatePtrType CertificateAuthorityCertificateArgs

func (*certificateAuthorityCertificatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateAuthorityCertificate)(nil))
}

func (i *certificateAuthorityCertificatePtrType) ToCertificateAuthorityCertificatePtrOutput() CertificateAuthorityCertificatePtrOutput {
	return i.ToCertificateAuthorityCertificatePtrOutputWithContext(context.Background())
}

func (i *certificateAuthorityCertificatePtrType) ToCertificateAuthorityCertificatePtrOutputWithContext(ctx context.Context) CertificateAuthorityCertificatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateAuthorityCertificatePtrOutput)
}

// CertificateAuthorityCertificateArrayInput is an input type that accepts CertificateAuthorityCertificateArray and CertificateAuthorityCertificateArrayOutput values.
// You can construct a concrete instance of `CertificateAuthorityCertificateArrayInput` via:
//
//          CertificateAuthorityCertificateArray{ CertificateAuthorityCertificateArgs{...} }
type CertificateAuthorityCertificateArrayInput interface {
	pulumi.Input

	ToCertificateAuthorityCertificateArrayOutput() CertificateAuthorityCertificateArrayOutput
	ToCertificateAuthorityCertificateArrayOutputWithContext(context.Context) CertificateAuthorityCertificateArrayOutput
}

type CertificateAuthorityCertificateArray []CertificateAuthorityCertificateInput

func (CertificateAuthorityCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*CertificateAuthorityCertificate)(nil))
}

func (i CertificateAuthorityCertificateArray) ToCertificateAuthorityCertificateArrayOutput() CertificateAuthorityCertificateArrayOutput {
	return i.ToCertificateAuthorityCertificateArrayOutputWithContext(context.Background())
}

func (i CertificateAuthorityCertificateArray) ToCertificateAuthorityCertificateArrayOutputWithContext(ctx context.Context) CertificateAuthorityCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateAuthorityCertificateArrayOutput)
}

// CertificateAuthorityCertificateMapInput is an input type that accepts CertificateAuthorityCertificateMap and CertificateAuthorityCertificateMapOutput values.
// You can construct a concrete instance of `CertificateAuthorityCertificateMapInput` via:
//
//          CertificateAuthorityCertificateMap{ "key": CertificateAuthorityCertificateArgs{...} }
type CertificateAuthorityCertificateMapInput interface {
	pulumi.Input

	ToCertificateAuthorityCertificateMapOutput() CertificateAuthorityCertificateMapOutput
	ToCertificateAuthorityCertificateMapOutputWithContext(context.Context) CertificateAuthorityCertificateMapOutput
}

type CertificateAuthorityCertificateMap map[string]CertificateAuthorityCertificateInput

func (CertificateAuthorityCertificateMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*CertificateAuthorityCertificate)(nil))
}

func (i CertificateAuthorityCertificateMap) ToCertificateAuthorityCertificateMapOutput() CertificateAuthorityCertificateMapOutput {
	return i.ToCertificateAuthorityCertificateMapOutputWithContext(context.Background())
}

func (i CertificateAuthorityCertificateMap) ToCertificateAuthorityCertificateMapOutputWithContext(ctx context.Context) CertificateAuthorityCertificateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateAuthorityCertificateMapOutput)
}

type CertificateAuthorityCertificateOutput struct {
	*pulumi.OutputState
}

func (CertificateAuthorityCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateAuthorityCertificate)(nil))
}

func (o CertificateAuthorityCertificateOutput) ToCertificateAuthorityCertificateOutput() CertificateAuthorityCertificateOutput {
	return o
}

func (o CertificateAuthorityCertificateOutput) ToCertificateAuthorityCertificateOutputWithContext(ctx context.Context) CertificateAuthorityCertificateOutput {
	return o
}

func (o CertificateAuthorityCertificateOutput) ToCertificateAuthorityCertificatePtrOutput() CertificateAuthorityCertificatePtrOutput {
	return o.ToCertificateAuthorityCertificatePtrOutputWithContext(context.Background())
}

func (o CertificateAuthorityCertificateOutput) ToCertificateAuthorityCertificatePtrOutputWithContext(ctx context.Context) CertificateAuthorityCertificatePtrOutput {
	return o.ApplyT(func(v CertificateAuthorityCertificate) *CertificateAuthorityCertificate {
		return &v
	}).(CertificateAuthorityCertificatePtrOutput)
}

type CertificateAuthorityCertificatePtrOutput struct {
	*pulumi.OutputState
}

func (CertificateAuthorityCertificatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateAuthorityCertificate)(nil))
}

func (o CertificateAuthorityCertificatePtrOutput) ToCertificateAuthorityCertificatePtrOutput() CertificateAuthorityCertificatePtrOutput {
	return o
}

func (o CertificateAuthorityCertificatePtrOutput) ToCertificateAuthorityCertificatePtrOutputWithContext(ctx context.Context) CertificateAuthorityCertificatePtrOutput {
	return o
}

type CertificateAuthorityCertificateArrayOutput struct{ *pulumi.OutputState }

func (CertificateAuthorityCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CertificateAuthorityCertificate)(nil))
}

func (o CertificateAuthorityCertificateArrayOutput) ToCertificateAuthorityCertificateArrayOutput() CertificateAuthorityCertificateArrayOutput {
	return o
}

func (o CertificateAuthorityCertificateArrayOutput) ToCertificateAuthorityCertificateArrayOutputWithContext(ctx context.Context) CertificateAuthorityCertificateArrayOutput {
	return o
}

func (o CertificateAuthorityCertificateArrayOutput) Index(i pulumi.IntInput) CertificateAuthorityCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CertificateAuthorityCertificate {
		return vs[0].([]CertificateAuthorityCertificate)[vs[1].(int)]
	}).(CertificateAuthorityCertificateOutput)
}

type CertificateAuthorityCertificateMapOutput struct{ *pulumi.OutputState }

func (CertificateAuthorityCertificateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CertificateAuthorityCertificate)(nil))
}

func (o CertificateAuthorityCertificateMapOutput) ToCertificateAuthorityCertificateMapOutput() CertificateAuthorityCertificateMapOutput {
	return o
}

func (o CertificateAuthorityCertificateMapOutput) ToCertificateAuthorityCertificateMapOutputWithContext(ctx context.Context) CertificateAuthorityCertificateMapOutput {
	return o
}

func (o CertificateAuthorityCertificateMapOutput) MapIndex(k pulumi.StringInput) CertificateAuthorityCertificateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CertificateAuthorityCertificate {
		return vs[0].(map[string]CertificateAuthorityCertificate)[vs[1].(string)]
	}).(CertificateAuthorityCertificateOutput)
}

func init() {
	pulumi.RegisterOutputType(CertificateAuthorityCertificateOutput{})
	pulumi.RegisterOutputType(CertificateAuthorityCertificatePtrOutput{})
	pulumi.RegisterOutputType(CertificateAuthorityCertificateArrayOutput{})
	pulumi.RegisterOutputType(CertificateAuthorityCertificateMapOutput{})
}
