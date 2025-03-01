# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['Vpc']


class Vpc(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 assign_generated_ipv6_cidr_block: Optional[pulumi.Input[bool]] = None,
                 cidr_block: Optional[pulumi.Input[str]] = None,
                 enable_classiclink: Optional[pulumi.Input[bool]] = None,
                 enable_classiclink_dns_support: Optional[pulumi.Input[bool]] = None,
                 enable_dns_hostnames: Optional[pulumi.Input[bool]] = None,
                 enable_dns_support: Optional[pulumi.Input[bool]] = None,
                 instance_tenancy: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a VPC resource.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_aws as aws

        main = aws.ec2.Vpc("main", cidr_block="10.0.0.0/16")
        ```

        Basic usage with tags:

        ```python
        import pulumi
        import pulumi_aws as aws

        main = aws.ec2.Vpc("main",
            cidr_block="10.0.0.0/16",
            instance_tenancy="default",
            tags={
                "Name": "main",
            })
        ```

        ## Import

        VPCs can be imported using the `vpc id`, e.g.

        ```sh
         $ pulumi import aws:ec2/vpc:Vpc test_vpc vpc-a01106c2
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] assign_generated_ipv6_cidr_block: Requests an Amazon-provided IPv6 CIDR
               block with a /56 prefix length for the VPC. You cannot specify the range of IP addresses, or
               the size of the CIDR block. Default is `false`.
        :param pulumi.Input[str] cidr_block: The CIDR block for the VPC.
        :param pulumi.Input[bool] enable_classiclink: A boolean flag to enable/disable ClassicLink
               for the VPC. Only valid in regions and accounts that support EC2 Classic.
               See the [ClassicLink documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) for more information. Defaults false.
        :param pulumi.Input[bool] enable_classiclink_dns_support: A boolean flag to enable/disable ClassicLink DNS Support for the VPC.
               Only valid in regions and accounts that support EC2 Classic.
        :param pulumi.Input[bool] enable_dns_hostnames: A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
        :param pulumi.Input[bool] enable_dns_support: A boolean flag to enable/disable DNS support in the VPC. Defaults true.
        :param pulumi.Input[str] instance_tenancy: A tenancy option for instances launched into the VPC. Default is `default`, which
               makes your instances shared on the host. Using either of the other options (`dedicated` or `host`) costs at least $2/hr.
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

            __props__['assign_generated_ipv6_cidr_block'] = assign_generated_ipv6_cidr_block
            if cidr_block is None and not opts.urn:
                raise TypeError("Missing required property 'cidr_block'")
            __props__['cidr_block'] = cidr_block
            __props__['enable_classiclink'] = enable_classiclink
            __props__['enable_classiclink_dns_support'] = enable_classiclink_dns_support
            __props__['enable_dns_hostnames'] = enable_dns_hostnames
            __props__['enable_dns_support'] = enable_dns_support
            __props__['instance_tenancy'] = instance_tenancy
            __props__['tags'] = tags
            __props__['tags_all'] = tags_all
            __props__['arn'] = None
            __props__['default_network_acl_id'] = None
            __props__['default_route_table_id'] = None
            __props__['default_security_group_id'] = None
            __props__['dhcp_options_id'] = None
            __props__['ipv6_association_id'] = None
            __props__['ipv6_cidr_block'] = None
            __props__['main_route_table_id'] = None
            __props__['owner_id'] = None
        super(Vpc, __self__).__init__(
            'aws:ec2/vpc:Vpc',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            assign_generated_ipv6_cidr_block: Optional[pulumi.Input[bool]] = None,
            cidr_block: Optional[pulumi.Input[str]] = None,
            default_network_acl_id: Optional[pulumi.Input[str]] = None,
            default_route_table_id: Optional[pulumi.Input[str]] = None,
            default_security_group_id: Optional[pulumi.Input[str]] = None,
            dhcp_options_id: Optional[pulumi.Input[str]] = None,
            enable_classiclink: Optional[pulumi.Input[bool]] = None,
            enable_classiclink_dns_support: Optional[pulumi.Input[bool]] = None,
            enable_dns_hostnames: Optional[pulumi.Input[bool]] = None,
            enable_dns_support: Optional[pulumi.Input[bool]] = None,
            instance_tenancy: Optional[pulumi.Input[str]] = None,
            ipv6_association_id: Optional[pulumi.Input[str]] = None,
            ipv6_cidr_block: Optional[pulumi.Input[str]] = None,
            main_route_table_id: Optional[pulumi.Input[str]] = None,
            owner_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Vpc':
        """
        Get an existing Vpc resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of VPC
        :param pulumi.Input[bool] assign_generated_ipv6_cidr_block: Requests an Amazon-provided IPv6 CIDR
               block with a /56 prefix length for the VPC. You cannot specify the range of IP addresses, or
               the size of the CIDR block. Default is `false`.
        :param pulumi.Input[str] cidr_block: The CIDR block for the VPC.
        :param pulumi.Input[str] default_network_acl_id: The ID of the network ACL created by default on VPC creation
        :param pulumi.Input[str] default_route_table_id: The ID of the route table created by default on VPC creation
        :param pulumi.Input[str] default_security_group_id: The ID of the security group created by default on VPC creation
        :param pulumi.Input[bool] enable_classiclink: A boolean flag to enable/disable ClassicLink
               for the VPC. Only valid in regions and accounts that support EC2 Classic.
               See the [ClassicLink documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) for more information. Defaults false.
        :param pulumi.Input[bool] enable_classiclink_dns_support: A boolean flag to enable/disable ClassicLink DNS Support for the VPC.
               Only valid in regions and accounts that support EC2 Classic.
        :param pulumi.Input[bool] enable_dns_hostnames: A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
        :param pulumi.Input[bool] enable_dns_support: A boolean flag to enable/disable DNS support in the VPC. Defaults true.
        :param pulumi.Input[str] instance_tenancy: A tenancy option for instances launched into the VPC. Default is `default`, which
               makes your instances shared on the host. Using either of the other options (`dedicated` or `host`) costs at least $2/hr.
        :param pulumi.Input[str] ipv6_association_id: The association ID for the IPv6 CIDR block.
        :param pulumi.Input[str] ipv6_cidr_block: The IPv6 CIDR block.
        :param pulumi.Input[str] main_route_table_id: The ID of the main route table associated with
               this VPC. Note that you can change a VPC's main route table by using an
               `ec2.MainRouteTableAssociation`.
        :param pulumi.Input[str] owner_id: The ID of the AWS account that owns the VPC.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["assign_generated_ipv6_cidr_block"] = assign_generated_ipv6_cidr_block
        __props__["cidr_block"] = cidr_block
        __props__["default_network_acl_id"] = default_network_acl_id
        __props__["default_route_table_id"] = default_route_table_id
        __props__["default_security_group_id"] = default_security_group_id
        __props__["dhcp_options_id"] = dhcp_options_id
        __props__["enable_classiclink"] = enable_classiclink
        __props__["enable_classiclink_dns_support"] = enable_classiclink_dns_support
        __props__["enable_dns_hostnames"] = enable_dns_hostnames
        __props__["enable_dns_support"] = enable_dns_support
        __props__["instance_tenancy"] = instance_tenancy
        __props__["ipv6_association_id"] = ipv6_association_id
        __props__["ipv6_cidr_block"] = ipv6_cidr_block
        __props__["main_route_table_id"] = main_route_table_id
        __props__["owner_id"] = owner_id
        __props__["tags"] = tags
        __props__["tags_all"] = tags_all
        return Vpc(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of VPC
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="assignGeneratedIpv6CidrBlock")
    def assign_generated_ipv6_cidr_block(self) -> pulumi.Output[Optional[bool]]:
        """
        Requests an Amazon-provided IPv6 CIDR
        block with a /56 prefix length for the VPC. You cannot specify the range of IP addresses, or
        the size of the CIDR block. Default is `false`.
        """
        return pulumi.get(self, "assign_generated_ipv6_cidr_block")

    @property
    @pulumi.getter(name="cidrBlock")
    def cidr_block(self) -> pulumi.Output[str]:
        """
        The CIDR block for the VPC.
        """
        return pulumi.get(self, "cidr_block")

    @property
    @pulumi.getter(name="defaultNetworkAclId")
    def default_network_acl_id(self) -> pulumi.Output[str]:
        """
        The ID of the network ACL created by default on VPC creation
        """
        return pulumi.get(self, "default_network_acl_id")

    @property
    @pulumi.getter(name="defaultRouteTableId")
    def default_route_table_id(self) -> pulumi.Output[str]:
        """
        The ID of the route table created by default on VPC creation
        """
        return pulumi.get(self, "default_route_table_id")

    @property
    @pulumi.getter(name="defaultSecurityGroupId")
    def default_security_group_id(self) -> pulumi.Output[str]:
        """
        The ID of the security group created by default on VPC creation
        """
        return pulumi.get(self, "default_security_group_id")

    @property
    @pulumi.getter(name="dhcpOptionsId")
    def dhcp_options_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "dhcp_options_id")

    @property
    @pulumi.getter(name="enableClassiclink")
    def enable_classiclink(self) -> pulumi.Output[bool]:
        """
        A boolean flag to enable/disable ClassicLink
        for the VPC. Only valid in regions and accounts that support EC2 Classic.
        See the [ClassicLink documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) for more information. Defaults false.
        """
        return pulumi.get(self, "enable_classiclink")

    @property
    @pulumi.getter(name="enableClassiclinkDnsSupport")
    def enable_classiclink_dns_support(self) -> pulumi.Output[bool]:
        """
        A boolean flag to enable/disable ClassicLink DNS Support for the VPC.
        Only valid in regions and accounts that support EC2 Classic.
        """
        return pulumi.get(self, "enable_classiclink_dns_support")

    @property
    @pulumi.getter(name="enableDnsHostnames")
    def enable_dns_hostnames(self) -> pulumi.Output[bool]:
        """
        A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
        """
        return pulumi.get(self, "enable_dns_hostnames")

    @property
    @pulumi.getter(name="enableDnsSupport")
    def enable_dns_support(self) -> pulumi.Output[Optional[bool]]:
        """
        A boolean flag to enable/disable DNS support in the VPC. Defaults true.
        """
        return pulumi.get(self, "enable_dns_support")

    @property
    @pulumi.getter(name="instanceTenancy")
    def instance_tenancy(self) -> pulumi.Output[Optional[str]]:
        """
        A tenancy option for instances launched into the VPC. Default is `default`, which
        makes your instances shared on the host. Using either of the other options (`dedicated` or `host`) costs at least $2/hr.
        """
        return pulumi.get(self, "instance_tenancy")

    @property
    @pulumi.getter(name="ipv6AssociationId")
    def ipv6_association_id(self) -> pulumi.Output[str]:
        """
        The association ID for the IPv6 CIDR block.
        """
        return pulumi.get(self, "ipv6_association_id")

    @property
    @pulumi.getter(name="ipv6CidrBlock")
    def ipv6_cidr_block(self) -> pulumi.Output[str]:
        """
        The IPv6 CIDR block.
        """
        return pulumi.get(self, "ipv6_cidr_block")

    @property
    @pulumi.getter(name="mainRouteTableId")
    def main_route_table_id(self) -> pulumi.Output[str]:
        """
        The ID of the main route table associated with
        this VPC. Note that you can change a VPC's main route table by using an
        `ec2.MainRouteTableAssociation`.
        """
        return pulumi.get(self, "main_route_table_id")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> pulumi.Output[str]:
        """
        The ID of the AWS account that owns the VPC.
        """
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        return pulumi.get(self, "tags_all")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

