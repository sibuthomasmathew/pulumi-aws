# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['Cluster']


class Cluster(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 apply_immediately: Optional[pulumi.Input[bool]] = None,
                 availability_zones: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 backup_retention_period: Optional[pulumi.Input[int]] = None,
                 cluster_identifier: Optional[pulumi.Input[str]] = None,
                 cluster_identifier_prefix: Optional[pulumi.Input[str]] = None,
                 cluster_members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 db_cluster_parameter_group_name: Optional[pulumi.Input[str]] = None,
                 db_subnet_group_name: Optional[pulumi.Input[str]] = None,
                 deletion_protection: Optional[pulumi.Input[bool]] = None,
                 enabled_cloudwatch_logs_exports: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 engine: Optional[pulumi.Input[str]] = None,
                 engine_version: Optional[pulumi.Input[str]] = None,
                 final_snapshot_identifier: Optional[pulumi.Input[str]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 master_password: Optional[pulumi.Input[str]] = None,
                 master_username: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 preferred_backup_window: Optional[pulumi.Input[str]] = None,
                 preferred_maintenance_window: Optional[pulumi.Input[str]] = None,
                 skip_final_snapshot: Optional[pulumi.Input[bool]] = None,
                 snapshot_identifier: Optional[pulumi.Input[str]] = None,
                 storage_encrypted: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vpc_security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a DocDB Cluster.

        Changes to a DocDB Cluster can occur when you manually change a
        parameter, such as `port`, and are reflected in the next maintenance
        window. Because of this, this provider may report a difference in its planning
        phase because a modification has not yet taken place. You can use the
        `apply_immediately` flag to instruct the service to apply the change immediately
        (see documentation below).

        > **Note:** using `apply_immediately` can result in a brief downtime as the server reboots.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        docdb = aws.docdb.Cluster("docdb",
            backup_retention_period=5,
            cluster_identifier="my-docdb-cluster",
            engine="docdb",
            master_password="mustbeeightchars",
            master_username="foo",
            preferred_backup_window="07:00-09:00",
            skip_final_snapshot=True)
        ```

        ## Import

        DocDB Clusters can be imported using the `cluster_identifier`, e.g.

        ```sh
         $ pulumi import aws:docdb/cluster:Cluster docdb_cluster docdb-prod-cluster
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] apply_immediately: Specifies whether any cluster modifications
               are applied immediately, or during the next maintenance window. Default is
               `false`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] availability_zones: A list of EC2 Availability Zones that
               instances in the DB cluster can be created in.
        :param pulumi.Input[int] backup_retention_period: The days to retain backups for. Default `1`
        :param pulumi.Input[str] cluster_identifier: The cluster identifier. If omitted, this provider will assign a random, unique identifier.
        :param pulumi.Input[str] cluster_identifier_prefix: Creates a unique cluster identifier beginning with the specified prefix. Conflicts with `cluster_identifer`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] cluster_members: List of DocDB Instances that are a part of this cluster
        :param pulumi.Input[str] db_cluster_parameter_group_name: A cluster parameter group to associate with the cluster.
        :param pulumi.Input[str] db_subnet_group_name: A DB subnet group to associate with this DB instance.
        :param pulumi.Input[bool] deletion_protection: A value that indicates whether the DB cluster has deletion protection enabled. The database can't be deleted when deletion protection is enabled. By default, deletion protection is disabled.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] enabled_cloudwatch_logs_exports: List of log types to export to cloudwatch. If omitted, no logs will be exported.
               The following log types are supported: `audit`, `profiler`.
        :param pulumi.Input[str] engine: The name of the database engine to be used for this DB cluster. Defaults to `docdb`. Valid Values: `docdb`
        :param pulumi.Input[str] engine_version: The database engine version. Updating this argument results in an outage.
        :param pulumi.Input[str] final_snapshot_identifier: The name of your final DB snapshot
               when this DB cluster is deleted. If omitted, no final snapshot will be
               made.
        :param pulumi.Input[str] kms_key_id: The ARN for the KMS encryption key. When specifying `kms_key_id`, `storage_encrypted` needs to be set to true.
        :param pulumi.Input[str] master_password: Password for the master DB user. Note that this may
               show up in logs, and it will be stored in the state file. Please refer to the DocDB Naming Constraints.
        :param pulumi.Input[str] master_username: Username for the master DB user.
        :param pulumi.Input[int] port: The port on which the DB accepts connections
        :param pulumi.Input[str] preferred_backup_window: The daily time range during which automated backups are created if automated backups are enabled using the BackupRetentionPeriod parameter.Time in UTC
               Default: A 30-minute window selected at random from an 8-hour block of time per region. e.g. 04:00-09:00
        :param pulumi.Input[str] preferred_maintenance_window: The weekly time range during which system maintenance can occur, in (UTC) e.g. wed:04:00-wed:04:30
        :param pulumi.Input[bool] skip_final_snapshot: Determines whether a final DB snapshot is created before the DB cluster is deleted. If true is specified, no DB snapshot is created. If false is specified, a DB snapshot is created before the DB cluster is deleted, using the value from `final_snapshot_identifier`. Default is `false`.
        :param pulumi.Input[str] snapshot_identifier: Specifies whether or not to create this cluster from a snapshot. You can use either the name or ARN when specifying a DB cluster snapshot, or the ARN when specifying a DB snapshot.
        :param pulumi.Input[bool] storage_encrypted: Specifies whether the DB cluster is encrypted. The default is `false`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the DB cluster.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vpc_security_group_ids: List of VPC security groups to associate
               with the Cluster
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['apply_immediately'] = apply_immediately
            __props__['availability_zones'] = availability_zones
            __props__['backup_retention_period'] = backup_retention_period
            __props__['cluster_identifier'] = cluster_identifier
            __props__['cluster_identifier_prefix'] = cluster_identifier_prefix
            __props__['cluster_members'] = cluster_members
            __props__['db_cluster_parameter_group_name'] = db_cluster_parameter_group_name
            __props__['db_subnet_group_name'] = db_subnet_group_name
            __props__['deletion_protection'] = deletion_protection
            __props__['enabled_cloudwatch_logs_exports'] = enabled_cloudwatch_logs_exports
            __props__['engine'] = engine
            __props__['engine_version'] = engine_version
            __props__['final_snapshot_identifier'] = final_snapshot_identifier
            __props__['kms_key_id'] = kms_key_id
            __props__['master_password'] = master_password
            __props__['master_username'] = master_username
            __props__['port'] = port
            __props__['preferred_backup_window'] = preferred_backup_window
            __props__['preferred_maintenance_window'] = preferred_maintenance_window
            __props__['skip_final_snapshot'] = skip_final_snapshot
            __props__['snapshot_identifier'] = snapshot_identifier
            __props__['storage_encrypted'] = storage_encrypted
            __props__['tags'] = tags
            __props__['vpc_security_group_ids'] = vpc_security_group_ids
            __props__['arn'] = None
            __props__['cluster_resource_id'] = None
            __props__['endpoint'] = None
            __props__['hosted_zone_id'] = None
            __props__['reader_endpoint'] = None
        super(Cluster, __self__).__init__(
            'aws:docdb/cluster:Cluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            apply_immediately: Optional[pulumi.Input[bool]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            availability_zones: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            backup_retention_period: Optional[pulumi.Input[int]] = None,
            cluster_identifier: Optional[pulumi.Input[str]] = None,
            cluster_identifier_prefix: Optional[pulumi.Input[str]] = None,
            cluster_members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            cluster_resource_id: Optional[pulumi.Input[str]] = None,
            db_cluster_parameter_group_name: Optional[pulumi.Input[str]] = None,
            db_subnet_group_name: Optional[pulumi.Input[str]] = None,
            deletion_protection: Optional[pulumi.Input[bool]] = None,
            enabled_cloudwatch_logs_exports: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            endpoint: Optional[pulumi.Input[str]] = None,
            engine: Optional[pulumi.Input[str]] = None,
            engine_version: Optional[pulumi.Input[str]] = None,
            final_snapshot_identifier: Optional[pulumi.Input[str]] = None,
            hosted_zone_id: Optional[pulumi.Input[str]] = None,
            kms_key_id: Optional[pulumi.Input[str]] = None,
            master_password: Optional[pulumi.Input[str]] = None,
            master_username: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[int]] = None,
            preferred_backup_window: Optional[pulumi.Input[str]] = None,
            preferred_maintenance_window: Optional[pulumi.Input[str]] = None,
            reader_endpoint: Optional[pulumi.Input[str]] = None,
            skip_final_snapshot: Optional[pulumi.Input[bool]] = None,
            snapshot_identifier: Optional[pulumi.Input[str]] = None,
            storage_encrypted: Optional[pulumi.Input[bool]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            vpc_security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'Cluster':
        """
        Get an existing Cluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] apply_immediately: Specifies whether any cluster modifications
               are applied immediately, or during the next maintenance window. Default is
               `false`.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of cluster
        :param pulumi.Input[Sequence[pulumi.Input[str]]] availability_zones: A list of EC2 Availability Zones that
               instances in the DB cluster can be created in.
        :param pulumi.Input[int] backup_retention_period: The days to retain backups for. Default `1`
        :param pulumi.Input[str] cluster_identifier: The cluster identifier. If omitted, this provider will assign a random, unique identifier.
        :param pulumi.Input[str] cluster_identifier_prefix: Creates a unique cluster identifier beginning with the specified prefix. Conflicts with `cluster_identifer`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] cluster_members: List of DocDB Instances that are a part of this cluster
        :param pulumi.Input[str] cluster_resource_id: The DocDB Cluster Resource ID
        :param pulumi.Input[str] db_cluster_parameter_group_name: A cluster parameter group to associate with the cluster.
        :param pulumi.Input[str] db_subnet_group_name: A DB subnet group to associate with this DB instance.
        :param pulumi.Input[bool] deletion_protection: A value that indicates whether the DB cluster has deletion protection enabled. The database can't be deleted when deletion protection is enabled. By default, deletion protection is disabled.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] enabled_cloudwatch_logs_exports: List of log types to export to cloudwatch. If omitted, no logs will be exported.
               The following log types are supported: `audit`, `profiler`.
        :param pulumi.Input[str] endpoint: The DNS address of the DocDB instance
        :param pulumi.Input[str] engine: The name of the database engine to be used for this DB cluster. Defaults to `docdb`. Valid Values: `docdb`
        :param pulumi.Input[str] engine_version: The database engine version. Updating this argument results in an outage.
        :param pulumi.Input[str] final_snapshot_identifier: The name of your final DB snapshot
               when this DB cluster is deleted. If omitted, no final snapshot will be
               made.
        :param pulumi.Input[str] hosted_zone_id: The Route53 Hosted Zone ID of the endpoint
        :param pulumi.Input[str] kms_key_id: The ARN for the KMS encryption key. When specifying `kms_key_id`, `storage_encrypted` needs to be set to true.
        :param pulumi.Input[str] master_password: Password for the master DB user. Note that this may
               show up in logs, and it will be stored in the state file. Please refer to the DocDB Naming Constraints.
        :param pulumi.Input[str] master_username: Username for the master DB user.
        :param pulumi.Input[int] port: The port on which the DB accepts connections
        :param pulumi.Input[str] preferred_backup_window: The daily time range during which automated backups are created if automated backups are enabled using the BackupRetentionPeriod parameter.Time in UTC
               Default: A 30-minute window selected at random from an 8-hour block of time per region. e.g. 04:00-09:00
        :param pulumi.Input[str] preferred_maintenance_window: The weekly time range during which system maintenance can occur, in (UTC) e.g. wed:04:00-wed:04:30
        :param pulumi.Input[str] reader_endpoint: A read-only endpoint for the DocDB cluster, automatically load-balanced across replicas
        :param pulumi.Input[bool] skip_final_snapshot: Determines whether a final DB snapshot is created before the DB cluster is deleted. If true is specified, no DB snapshot is created. If false is specified, a DB snapshot is created before the DB cluster is deleted, using the value from `final_snapshot_identifier`. Default is `false`.
        :param pulumi.Input[str] snapshot_identifier: Specifies whether or not to create this cluster from a snapshot. You can use either the name or ARN when specifying a DB cluster snapshot, or the ARN when specifying a DB snapshot.
        :param pulumi.Input[bool] storage_encrypted: Specifies whether the DB cluster is encrypted. The default is `false`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the DB cluster.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vpc_security_group_ids: List of VPC security groups to associate
               with the Cluster
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["apply_immediately"] = apply_immediately
        __props__["arn"] = arn
        __props__["availability_zones"] = availability_zones
        __props__["backup_retention_period"] = backup_retention_period
        __props__["cluster_identifier"] = cluster_identifier
        __props__["cluster_identifier_prefix"] = cluster_identifier_prefix
        __props__["cluster_members"] = cluster_members
        __props__["cluster_resource_id"] = cluster_resource_id
        __props__["db_cluster_parameter_group_name"] = db_cluster_parameter_group_name
        __props__["db_subnet_group_name"] = db_subnet_group_name
        __props__["deletion_protection"] = deletion_protection
        __props__["enabled_cloudwatch_logs_exports"] = enabled_cloudwatch_logs_exports
        __props__["endpoint"] = endpoint
        __props__["engine"] = engine
        __props__["engine_version"] = engine_version
        __props__["final_snapshot_identifier"] = final_snapshot_identifier
        __props__["hosted_zone_id"] = hosted_zone_id
        __props__["kms_key_id"] = kms_key_id
        __props__["master_password"] = master_password
        __props__["master_username"] = master_username
        __props__["port"] = port
        __props__["preferred_backup_window"] = preferred_backup_window
        __props__["preferred_maintenance_window"] = preferred_maintenance_window
        __props__["reader_endpoint"] = reader_endpoint
        __props__["skip_final_snapshot"] = skip_final_snapshot
        __props__["snapshot_identifier"] = snapshot_identifier
        __props__["storage_encrypted"] = storage_encrypted
        __props__["tags"] = tags
        __props__["vpc_security_group_ids"] = vpc_security_group_ids
        return Cluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applyImmediately")
    def apply_immediately(self) -> pulumi.Output[bool]:
        """
        Specifies whether any cluster modifications
        are applied immediately, or during the next maintenance window. Default is
        `false`.
        """
        return pulumi.get(self, "apply_immediately")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of cluster
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="availabilityZones")
    def availability_zones(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of EC2 Availability Zones that
        instances in the DB cluster can be created in.
        """
        return pulumi.get(self, "availability_zones")

    @property
    @pulumi.getter(name="backupRetentionPeriod")
    def backup_retention_period(self) -> pulumi.Output[Optional[int]]:
        """
        The days to retain backups for. Default `1`
        """
        return pulumi.get(self, "backup_retention_period")

    @property
    @pulumi.getter(name="clusterIdentifier")
    def cluster_identifier(self) -> pulumi.Output[str]:
        """
        The cluster identifier. If omitted, this provider will assign a random, unique identifier.
        """
        return pulumi.get(self, "cluster_identifier")

    @property
    @pulumi.getter(name="clusterIdentifierPrefix")
    def cluster_identifier_prefix(self) -> pulumi.Output[str]:
        """
        Creates a unique cluster identifier beginning with the specified prefix. Conflicts with `cluster_identifer`.
        """
        return pulumi.get(self, "cluster_identifier_prefix")

    @property
    @pulumi.getter(name="clusterMembers")
    def cluster_members(self) -> pulumi.Output[Sequence[str]]:
        """
        List of DocDB Instances that are a part of this cluster
        """
        return pulumi.get(self, "cluster_members")

    @property
    @pulumi.getter(name="clusterResourceId")
    def cluster_resource_id(self) -> pulumi.Output[str]:
        """
        The DocDB Cluster Resource ID
        """
        return pulumi.get(self, "cluster_resource_id")

    @property
    @pulumi.getter(name="dbClusterParameterGroupName")
    def db_cluster_parameter_group_name(self) -> pulumi.Output[str]:
        """
        A cluster parameter group to associate with the cluster.
        """
        return pulumi.get(self, "db_cluster_parameter_group_name")

    @property
    @pulumi.getter(name="dbSubnetGroupName")
    def db_subnet_group_name(self) -> pulumi.Output[str]:
        """
        A DB subnet group to associate with this DB instance.
        """
        return pulumi.get(self, "db_subnet_group_name")

    @property
    @pulumi.getter(name="deletionProtection")
    def deletion_protection(self) -> pulumi.Output[Optional[bool]]:
        """
        A value that indicates whether the DB cluster has deletion protection enabled. The database can't be deleted when deletion protection is enabled. By default, deletion protection is disabled.
        """
        return pulumi.get(self, "deletion_protection")

    @property
    @pulumi.getter(name="enabledCloudwatchLogsExports")
    def enabled_cloudwatch_logs_exports(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of log types to export to cloudwatch. If omitted, no logs will be exported.
        The following log types are supported: `audit`, `profiler`.
        """
        return pulumi.get(self, "enabled_cloudwatch_logs_exports")

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Output[str]:
        """
        The DNS address of the DocDB instance
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter
    def engine(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the database engine to be used for this DB cluster. Defaults to `docdb`. Valid Values: `docdb`
        """
        return pulumi.get(self, "engine")

    @property
    @pulumi.getter(name="engineVersion")
    def engine_version(self) -> pulumi.Output[str]:
        """
        The database engine version. Updating this argument results in an outage.
        """
        return pulumi.get(self, "engine_version")

    @property
    @pulumi.getter(name="finalSnapshotIdentifier")
    def final_snapshot_identifier(self) -> pulumi.Output[Optional[str]]:
        """
        The name of your final DB snapshot
        when this DB cluster is deleted. If omitted, no final snapshot will be
        made.
        """
        return pulumi.get(self, "final_snapshot_identifier")

    @property
    @pulumi.getter(name="hostedZoneId")
    def hosted_zone_id(self) -> pulumi.Output[str]:
        """
        The Route53 Hosted Zone ID of the endpoint
        """
        return pulumi.get(self, "hosted_zone_id")

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> pulumi.Output[str]:
        """
        The ARN for the KMS encryption key. When specifying `kms_key_id`, `storage_encrypted` needs to be set to true.
        """
        return pulumi.get(self, "kms_key_id")

    @property
    @pulumi.getter(name="masterPassword")
    def master_password(self) -> pulumi.Output[Optional[str]]:
        """
        Password for the master DB user. Note that this may
        show up in logs, and it will be stored in the state file. Please refer to the DocDB Naming Constraints.
        """
        return pulumi.get(self, "master_password")

    @property
    @pulumi.getter(name="masterUsername")
    def master_username(self) -> pulumi.Output[str]:
        """
        Username for the master DB user.
        """
        return pulumi.get(self, "master_username")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[Optional[int]]:
        """
        The port on which the DB accepts connections
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="preferredBackupWindow")
    def preferred_backup_window(self) -> pulumi.Output[str]:
        """
        The daily time range during which automated backups are created if automated backups are enabled using the BackupRetentionPeriod parameter.Time in UTC
        Default: A 30-minute window selected at random from an 8-hour block of time per region. e.g. 04:00-09:00
        """
        return pulumi.get(self, "preferred_backup_window")

    @property
    @pulumi.getter(name="preferredMaintenanceWindow")
    def preferred_maintenance_window(self) -> pulumi.Output[str]:
        """
        The weekly time range during which system maintenance can occur, in (UTC) e.g. wed:04:00-wed:04:30
        """
        return pulumi.get(self, "preferred_maintenance_window")

    @property
    @pulumi.getter(name="readerEndpoint")
    def reader_endpoint(self) -> pulumi.Output[str]:
        """
        A read-only endpoint for the DocDB cluster, automatically load-balanced across replicas
        """
        return pulumi.get(self, "reader_endpoint")

    @property
    @pulumi.getter(name="skipFinalSnapshot")
    def skip_final_snapshot(self) -> pulumi.Output[Optional[bool]]:
        """
        Determines whether a final DB snapshot is created before the DB cluster is deleted. If true is specified, no DB snapshot is created. If false is specified, a DB snapshot is created before the DB cluster is deleted, using the value from `final_snapshot_identifier`. Default is `false`.
        """
        return pulumi.get(self, "skip_final_snapshot")

    @property
    @pulumi.getter(name="snapshotIdentifier")
    def snapshot_identifier(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies whether or not to create this cluster from a snapshot. You can use either the name or ARN when specifying a DB cluster snapshot, or the ARN when specifying a DB snapshot.
        """
        return pulumi.get(self, "snapshot_identifier")

    @property
    @pulumi.getter(name="storageEncrypted")
    def storage_encrypted(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether the DB cluster is encrypted. The default is `false`.
        """
        return pulumi.get(self, "storage_encrypted")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the DB cluster.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcSecurityGroupIds")
    def vpc_security_group_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        List of VPC security groups to associate
        with the Cluster
        """
        return pulumi.get(self, "vpc_security_group_ids")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

