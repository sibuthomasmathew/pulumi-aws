// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Iam
{
    /// <summary>
    /// Provides an IAM Server Certificate resource to upload Server Certificates.
    /// Certs uploaded to IAM can easily work with other AWS services such as:
    /// 
    /// - AWS Elastic Beanstalk
    /// - Elastic Load Balancing
    /// - CloudFront
    /// - AWS OpsWorks
    /// 
    /// For information about server certificates in IAM, see [Managing Server
    /// Certificates][2] in AWS Documentation.
    /// 
    /// ## Example Usage
    /// 
    /// **Using certs on file:**
    /// 
    /// ```csharp
    /// using System.IO;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var testCert = new Aws.Iam.ServerCertificate("testCert", new Aws.Iam.ServerCertificateArgs
    ///         {
    ///             CertificateBody = File.ReadAllText("self-ca-cert.pem"),
    ///             PrivateKey = File.ReadAllText("test-key.pem"),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// **Example with cert in-line:**
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var testCertAlt = new Aws.Iam.ServerCertificate("testCertAlt", new Aws.Iam.ServerCertificateArgs
    ///         {
    ///             CertificateBody = @"-----BEGIN CERTIFICATE-----
    /// [......] # cert contents
    /// -----END CERTIFICATE-----
    /// 
    /// ",
    ///             PrivateKey = @"-----BEGIN RSA PRIVATE KEY-----
    /// [......] # cert contents
    /// -----END RSA PRIVATE KEY-----
    /// 
    /// ",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// **Use in combination with an AWS ELB resource:**
    /// 
    /// Some properties of an IAM Server Certificates cannot be updated while they are
    /// in use. In order for this provider to effectively manage a Certificate in this situation, it is
    /// recommended you utilize the `name_prefix` attribute and enable the
    /// `create_before_destroy` [lifecycle block][lifecycle]. This will allow this provider
    /// to create a new, updated `aws.iam.ServerCertificate` resource and replace it in
    /// dependant resources before attempting to destroy the old version.
    /// 
    /// ```csharp
    /// using System.IO;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var testCert = new Aws.Iam.ServerCertificate("testCert", new Aws.Iam.ServerCertificateArgs
    ///         {
    ///             NamePrefix = "example-cert",
    ///             CertificateBody = File.ReadAllText("self-ca-cert.pem"),
    ///             PrivateKey = File.ReadAllText("test-key.pem"),
    ///         });
    ///         var ourapp = new Aws.Elb.LoadBalancer("ourapp", new Aws.Elb.LoadBalancerArgs
    ///         {
    ///             AvailabilityZones = 
    ///             {
    ///                 "us-west-2a",
    ///             },
    ///             CrossZoneLoadBalancing = true,
    ///             Listeners = 
    ///             {
    ///                 new Aws.Elb.Inputs.LoadBalancerListenerArgs
    ///                 {
    ///                     InstancePort = 8000,
    ///                     InstanceProtocol = "http",
    ///                     LbPort = 443,
    ///                     LbProtocol = "https",
    ///                     SslCertificateId = testCert.Arn,
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// IAM Server Certificates can be imported using the `name`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:iam/serverCertificate:ServerCertificate certificate example.com-certificate-until-2018
    /// ```
    /// 
    ///  [1]https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html [2]https://docs.aws.amazon.com/IAM/latest/UserGuide/ManagingServerCerts.html [lifecycle]/docs/configuration/resources.html
    /// </summary>
    [AwsResourceType("aws:iam/serverCertificate:ServerCertificate")]
    public partial class ServerCertificate : Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) specifying the server certificate.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The contents of the public key certificate in
        /// PEM-encoded format.
        /// </summary>
        [Output("certificateBody")]
        public Output<string> CertificateBody { get; private set; } = null!;

        /// <summary>
        /// The contents of the certificate chain.
        /// This is typically a concatenation of the PEM-encoded public key certificates
        /// of the chain.
        /// </summary>
        [Output("certificateChain")]
        public Output<string?> CertificateChain { get; private set; } = null!;

        /// <summary>
        /// Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) on which the certificate is set to expire.
        /// </summary>
        [Output("expiration")]
        public Output<string> Expiration { get; private set; } = null!;

        /// <summary>
        /// The name of the Server Certificate. Do not include the
        /// path in this value. If omitted, this provider will assign a random, unique name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Creates a unique name beginning with the specified
        /// prefix. Conflicts with `name`.
        /// </summary>
        [Output("namePrefix")]
        public Output<string?> NamePrefix { get; private set; } = null!;

        /// <summary>
        /// The IAM path for the server certificate.  If it is not
        /// included, it defaults to a slash (/). If this certificate is for use with
        /// AWS CloudFront, the path must be in format `/cloudfront/your_path_here`.
        /// See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more details on IAM Paths.
        /// </summary>
        [Output("path")]
        public Output<string?> Path { get; private set; } = null!;

        /// <summary>
        /// The contents of the private key in PEM-encoded format.
        /// </summary>
        [Output("privateKey")]
        public Output<string> PrivateKey { get; private set; } = null!;

        /// <summary>
        /// Map of resource tags for the server certificate.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) when the server certificate was uploaded.
        /// </summary>
        [Output("uploadDate")]
        public Output<string> UploadDate { get; private set; } = null!;


        /// <summary>
        /// Create a ServerCertificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServerCertificate(string name, ServerCertificateArgs args, CustomResourceOptions? options = null)
            : base("aws:iam/serverCertificate:ServerCertificate", name, args ?? new ServerCertificateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServerCertificate(string name, Input<string> id, ServerCertificateState? state = null, CustomResourceOptions? options = null)
            : base("aws:iam/serverCertificate:ServerCertificate", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ServerCertificate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServerCertificate Get(string name, Input<string> id, ServerCertificateState? state = null, CustomResourceOptions? options = null)
        {
            return new ServerCertificate(name, id, state, options);
        }
    }

    public sealed class ServerCertificateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The contents of the public key certificate in
        /// PEM-encoded format.
        /// </summary>
        [Input("certificateBody", required: true)]
        public Input<string> CertificateBody { get; set; } = null!;

        /// <summary>
        /// The contents of the certificate chain.
        /// This is typically a concatenation of the PEM-encoded public key certificates
        /// of the chain.
        /// </summary>
        [Input("certificateChain")]
        public Input<string>? CertificateChain { get; set; }

        /// <summary>
        /// The name of the Server Certificate. Do not include the
        /// path in this value. If omitted, this provider will assign a random, unique name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Creates a unique name beginning with the specified
        /// prefix. Conflicts with `name`.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// The IAM path for the server certificate.  If it is not
        /// included, it defaults to a slash (/). If this certificate is for use with
        /// AWS CloudFront, the path must be in format `/cloudfront/your_path_here`.
        /// See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more details on IAM Paths.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// The contents of the private key in PEM-encoded format.
        /// </summary>
        [Input("privateKey", required: true)]
        public Input<string> PrivateKey { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of resource tags for the server certificate.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ServerCertificateArgs()
        {
        }
    }

    public sealed class ServerCertificateState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) specifying the server certificate.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The contents of the public key certificate in
        /// PEM-encoded format.
        /// </summary>
        [Input("certificateBody")]
        public Input<string>? CertificateBody { get; set; }

        /// <summary>
        /// The contents of the certificate chain.
        /// This is typically a concatenation of the PEM-encoded public key certificates
        /// of the chain.
        /// </summary>
        [Input("certificateChain")]
        public Input<string>? CertificateChain { get; set; }

        /// <summary>
        /// Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) on which the certificate is set to expire.
        /// </summary>
        [Input("expiration")]
        public Input<string>? Expiration { get; set; }

        /// <summary>
        /// The name of the Server Certificate. Do not include the
        /// path in this value. If omitted, this provider will assign a random, unique name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Creates a unique name beginning with the specified
        /// prefix. Conflicts with `name`.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// The IAM path for the server certificate.  If it is not
        /// included, it defaults to a slash (/). If this certificate is for use with
        /// AWS CloudFront, the path must be in format `/cloudfront/your_path_here`.
        /// See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more details on IAM Paths.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// The contents of the private key in PEM-encoded format.
        /// </summary>
        [Input("privateKey")]
        public Input<string>? PrivateKey { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of resource tags for the server certificate.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) when the server certificate was uploaded.
        /// </summary>
        [Input("uploadDate")]
        public Input<string>? UploadDate { get; set; }

        public ServerCertificateState()
        {
        }
    }
}
