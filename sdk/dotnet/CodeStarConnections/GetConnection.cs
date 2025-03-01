// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CodeStarConnections
{
    public static class GetConnection
    {
        /// <summary>
        /// Provides details about CodeStar Connection.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Aws.CodeStarConnections.GetConnection.InvokeAsync(new Aws.CodeStarConnections.GetConnectionArgs
        ///         {
        ///             Arn = aws_codestarconnections_connection.Example.Arn,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetConnectionResult> InvokeAsync(GetConnectionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetConnectionResult>("aws:codestarconnections/getConnection:getConnection", args ?? new GetConnectionArgs(), options.WithVersion());
    }


    public sealed class GetConnectionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The CodeStar Connection ARN.
        /// </summary>
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        [Input("tags")]
        private Dictionary<string, string>? _tags;

        /// <summary>
        /// Map of key-value resource tags to associate with the resource.
        /// </summary>
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetConnectionArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetConnectionResult
    {
        public readonly string Arn;
        /// <summary>
        /// The CodeStar Connection status. Possible values are `PENDING`, `AVAILABLE` and `ERROR`.
        /// </summary>
        public readonly string ConnectionStatus;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the CodeStar Connection. The name is unique in the calling AWS account.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The name of the external provider where your third-party code repository is configured. Possible values are `Bitbucket`, `GitHub`, or `GitHubEnterpriseServer`.
        /// </summary>
        public readonly string ProviderType;
        /// <summary>
        /// Map of key-value resource tags to associate with the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetConnectionResult(
            string arn,

            string connectionStatus,

            string id,

            string name,

            string providerType,

            ImmutableDictionary<string, string> tags)
        {
            Arn = arn;
            ConnectionStatus = connectionStatus;
            Id = id;
            Name = name;
            ProviderType = providerType;
            Tags = tags;
        }
    }
}
