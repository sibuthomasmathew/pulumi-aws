// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Mq
{
    public static class GetBroker
    {
        /// <summary>
        /// Provides information about a MQ Broker.
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
        ///         var config = new Config();
        ///         var brokerId = config.Get("brokerId") ?? "";
        ///         var brokerName = config.Get("brokerName") ?? "";
        ///         var byId = Output.Create(Aws.Mq.GetBroker.InvokeAsync(new Aws.Mq.GetBrokerArgs
        ///         {
        ///             BrokerId = brokerId,
        ///         }));
        ///         var byName = Output.Create(Aws.Mq.GetBroker.InvokeAsync(new Aws.Mq.GetBrokerArgs
        ///         {
        ///             BrokerName = brokerName,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetBrokerResult> InvokeAsync(GetBrokerArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBrokerResult>("aws:mq/getBroker:getBroker", args ?? new GetBrokerArgs(), options.WithVersion());
    }


    public sealed class GetBrokerArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique id of the mq broker.
        /// </summary>
        [Input("brokerId")]
        public string? BrokerId { get; set; }

        /// <summary>
        /// The unique name of the mq broker.
        /// </summary>
        [Input("brokerName")]
        public string? BrokerName { get; set; }

        [Input("logs")]
        public Inputs.GetBrokerLogsArgs? Logs { get; set; }

        [Input("tags")]
        private Dictionary<string, string>? _tags;
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetBrokerArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBrokerResult
    {
        public readonly string Arn;
        public readonly string AuthenticationStrategy;
        public readonly bool AutoMinorVersionUpgrade;
        public readonly string BrokerId;
        public readonly string BrokerName;
        public readonly Outputs.GetBrokerConfigurationResult Configuration;
        public readonly string DeploymentMode;
        public readonly ImmutableArray<Outputs.GetBrokerEncryptionOptionResult> EncryptionOptions;
        public readonly string EngineType;
        public readonly string EngineVersion;
        public readonly string HostInstanceType;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetBrokerInstanceResult> Instances;
        public readonly ImmutableArray<Outputs.GetBrokerLdapServerMetadataResult> LdapServerMetadatas;
        public readonly Outputs.GetBrokerLogsResult? Logs;
        public readonly Outputs.GetBrokerMaintenanceWindowStartTimeResult MaintenanceWindowStartTime;
        public readonly bool PubliclyAccessible;
        public readonly ImmutableArray<string> SecurityGroups;
        public readonly string StorageType;
        public readonly ImmutableArray<string> SubnetIds;
        public readonly ImmutableDictionary<string, string> Tags;
        public readonly ImmutableArray<Outputs.GetBrokerUserResult> Users;

        [OutputConstructor]
        private GetBrokerResult(
            string arn,

            string authenticationStrategy,

            bool autoMinorVersionUpgrade,

            string brokerId,

            string brokerName,

            Outputs.GetBrokerConfigurationResult configuration,

            string deploymentMode,

            ImmutableArray<Outputs.GetBrokerEncryptionOptionResult> encryptionOptions,

            string engineType,

            string engineVersion,

            string hostInstanceType,

            string id,

            ImmutableArray<Outputs.GetBrokerInstanceResult> instances,

            ImmutableArray<Outputs.GetBrokerLdapServerMetadataResult> ldapServerMetadatas,

            Outputs.GetBrokerLogsResult? logs,

            Outputs.GetBrokerMaintenanceWindowStartTimeResult maintenanceWindowStartTime,

            bool publiclyAccessible,

            ImmutableArray<string> securityGroups,

            string storageType,

            ImmutableArray<string> subnetIds,

            ImmutableDictionary<string, string> tags,

            ImmutableArray<Outputs.GetBrokerUserResult> users)
        {
            Arn = arn;
            AuthenticationStrategy = authenticationStrategy;
            AutoMinorVersionUpgrade = autoMinorVersionUpgrade;
            BrokerId = brokerId;
            BrokerName = brokerName;
            Configuration = configuration;
            DeploymentMode = deploymentMode;
            EncryptionOptions = encryptionOptions;
            EngineType = engineType;
            EngineVersion = engineVersion;
            HostInstanceType = hostInstanceType;
            Id = id;
            Instances = instances;
            LdapServerMetadatas = ldapServerMetadatas;
            Logs = logs;
            MaintenanceWindowStartTime = maintenanceWindowStartTime;
            PubliclyAccessible = publiclyAccessible;
            SecurityGroups = securityGroups;
            StorageType = storageType;
            SubnetIds = subnetIds;
            Tags = tags;
            Users = users;
        }
    }
}
