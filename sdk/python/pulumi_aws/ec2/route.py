# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['Route']


class Route(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination_cidr_block: Optional[pulumi.Input[str]] = None,
                 destination_ipv6_cidr_block: Optional[pulumi.Input[str]] = None,
                 egress_only_gateway_id: Optional[pulumi.Input[str]] = None,
                 gateway_id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 local_gateway_id: Optional[pulumi.Input[str]] = None,
                 nat_gateway_id: Optional[pulumi.Input[str]] = None,
                 network_interface_id: Optional[pulumi.Input[str]] = None,
                 route_table_id: Optional[pulumi.Input[str]] = None,
                 transit_gateway_id: Optional[pulumi.Input[str]] = None,
                 vpc_endpoint_id: Optional[pulumi.Input[str]] = None,
                 vpc_peering_connection_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        ## Import

        Individual routes can be imported using `ROUTETABLEID_DESTINATION`. For example, import a route in route table `rtb-656C65616E6F72` with an IPv4 destination CIDR of `10.42.0.0/16` like thisconsole

        ```sh
         $ pulumi import aws:ec2/route:Route my_route rtb-656C65616E6F72_10.42.0.0/16
        ```

         Import a route in route table `rtb-656C65616E6F72` with an IPv6 destination CIDR of `2620:0:2d0:200::8/125` similarlyconsole

        ```sh
         $ pulumi import aws:ec2/route:Route my_route rtb-656C65616E6F72_2620:0:2d0:200::8/125
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] destination_cidr_block: The destination CIDR block.
        :param pulumi.Input[str] destination_ipv6_cidr_block: The destination IPv6 CIDR block.
        :param pulumi.Input[str] egress_only_gateway_id: Identifier of a VPC Egress Only Internet Gateway.
        :param pulumi.Input[str] gateway_id: Identifier of a VPC internet gateway or a virtual private gateway.
        :param pulumi.Input[str] instance_id: Identifier of an EC2 instance.
        :param pulumi.Input[str] local_gateway_id: Identifier of a Outpost local gateway.
        :param pulumi.Input[str] nat_gateway_id: Identifier of a VPC NAT gateway.
        :param pulumi.Input[str] network_interface_id: Identifier of an EC2 network interface.
        :param pulumi.Input[str] route_table_id: The ID of the routing table.
        :param pulumi.Input[str] transit_gateway_id: Identifier of an EC2 Transit Gateway.
        :param pulumi.Input[str] vpc_endpoint_id: Identifier of a VPC Endpoint.
        :param pulumi.Input[str] vpc_peering_connection_id: Identifier of a VPC peering connection.
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

            __props__['destination_cidr_block'] = destination_cidr_block
            __props__['destination_ipv6_cidr_block'] = destination_ipv6_cidr_block
            __props__['egress_only_gateway_id'] = egress_only_gateway_id
            __props__['gateway_id'] = gateway_id
            __props__['instance_id'] = instance_id
            __props__['local_gateway_id'] = local_gateway_id
            __props__['nat_gateway_id'] = nat_gateway_id
            __props__['network_interface_id'] = network_interface_id
            if route_table_id is None and not opts.urn:
                raise TypeError("Missing required property 'route_table_id'")
            __props__['route_table_id'] = route_table_id
            __props__['transit_gateway_id'] = transit_gateway_id
            __props__['vpc_endpoint_id'] = vpc_endpoint_id
            __props__['vpc_peering_connection_id'] = vpc_peering_connection_id
            __props__['destination_prefix_list_id'] = None
            __props__['instance_owner_id'] = None
            __props__['origin'] = None
            __props__['state'] = None
        super(Route, __self__).__init__(
            'aws:ec2/route:Route',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            destination_cidr_block: Optional[pulumi.Input[str]] = None,
            destination_ipv6_cidr_block: Optional[pulumi.Input[str]] = None,
            destination_prefix_list_id: Optional[pulumi.Input[str]] = None,
            egress_only_gateway_id: Optional[pulumi.Input[str]] = None,
            gateway_id: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            instance_owner_id: Optional[pulumi.Input[str]] = None,
            local_gateway_id: Optional[pulumi.Input[str]] = None,
            nat_gateway_id: Optional[pulumi.Input[str]] = None,
            network_interface_id: Optional[pulumi.Input[str]] = None,
            origin: Optional[pulumi.Input[str]] = None,
            route_table_id: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[str]] = None,
            transit_gateway_id: Optional[pulumi.Input[str]] = None,
            vpc_endpoint_id: Optional[pulumi.Input[str]] = None,
            vpc_peering_connection_id: Optional[pulumi.Input[str]] = None) -> 'Route':
        """
        Get an existing Route resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] destination_cidr_block: The destination CIDR block.
        :param pulumi.Input[str] destination_ipv6_cidr_block: The destination IPv6 CIDR block.
        :param pulumi.Input[str] egress_only_gateway_id: Identifier of a VPC Egress Only Internet Gateway.
        :param pulumi.Input[str] gateway_id: Identifier of a VPC internet gateway or a virtual private gateway.
        :param pulumi.Input[str] instance_id: Identifier of an EC2 instance.
        :param pulumi.Input[str] local_gateway_id: Identifier of a Outpost local gateway.
        :param pulumi.Input[str] nat_gateway_id: Identifier of a VPC NAT gateway.
        :param pulumi.Input[str] network_interface_id: Identifier of an EC2 network interface.
        :param pulumi.Input[str] route_table_id: The ID of the routing table.
        :param pulumi.Input[str] transit_gateway_id: Identifier of an EC2 Transit Gateway.
        :param pulumi.Input[str] vpc_endpoint_id: Identifier of a VPC Endpoint.
        :param pulumi.Input[str] vpc_peering_connection_id: Identifier of a VPC peering connection.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["destination_cidr_block"] = destination_cidr_block
        __props__["destination_ipv6_cidr_block"] = destination_ipv6_cidr_block
        __props__["destination_prefix_list_id"] = destination_prefix_list_id
        __props__["egress_only_gateway_id"] = egress_only_gateway_id
        __props__["gateway_id"] = gateway_id
        __props__["instance_id"] = instance_id
        __props__["instance_owner_id"] = instance_owner_id
        __props__["local_gateway_id"] = local_gateway_id
        __props__["nat_gateway_id"] = nat_gateway_id
        __props__["network_interface_id"] = network_interface_id
        __props__["origin"] = origin
        __props__["route_table_id"] = route_table_id
        __props__["state"] = state
        __props__["transit_gateway_id"] = transit_gateway_id
        __props__["vpc_endpoint_id"] = vpc_endpoint_id
        __props__["vpc_peering_connection_id"] = vpc_peering_connection_id
        return Route(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="destinationCidrBlock")
    def destination_cidr_block(self) -> pulumi.Output[Optional[str]]:
        """
        The destination CIDR block.
        """
        return pulumi.get(self, "destination_cidr_block")

    @property
    @pulumi.getter(name="destinationIpv6CidrBlock")
    def destination_ipv6_cidr_block(self) -> pulumi.Output[Optional[str]]:
        """
        The destination IPv6 CIDR block.
        """
        return pulumi.get(self, "destination_ipv6_cidr_block")

    @property
    @pulumi.getter(name="destinationPrefixListId")
    def destination_prefix_list_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "destination_prefix_list_id")

    @property
    @pulumi.getter(name="egressOnlyGatewayId")
    def egress_only_gateway_id(self) -> pulumi.Output[Optional[str]]:
        """
        Identifier of a VPC Egress Only Internet Gateway.
        """
        return pulumi.get(self, "egress_only_gateway_id")

    @property
    @pulumi.getter(name="gatewayId")
    def gateway_id(self) -> pulumi.Output[Optional[str]]:
        """
        Identifier of a VPC internet gateway or a virtual private gateway.
        """
        return pulumi.get(self, "gateway_id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Identifier of an EC2 instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="instanceOwnerId")
    def instance_owner_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "instance_owner_id")

    @property
    @pulumi.getter(name="localGatewayId")
    def local_gateway_id(self) -> pulumi.Output[Optional[str]]:
        """
        Identifier of a Outpost local gateway.
        """
        return pulumi.get(self, "local_gateway_id")

    @property
    @pulumi.getter(name="natGatewayId")
    def nat_gateway_id(self) -> pulumi.Output[Optional[str]]:
        """
        Identifier of a VPC NAT gateway.
        """
        return pulumi.get(self, "nat_gateway_id")

    @property
    @pulumi.getter(name="networkInterfaceId")
    def network_interface_id(self) -> pulumi.Output[str]:
        """
        Identifier of an EC2 network interface.
        """
        return pulumi.get(self, "network_interface_id")

    @property
    @pulumi.getter
    def origin(self) -> pulumi.Output[str]:
        return pulumi.get(self, "origin")

    @property
    @pulumi.getter(name="routeTableId")
    def route_table_id(self) -> pulumi.Output[str]:
        """
        The ID of the routing table.
        """
        return pulumi.get(self, "route_table_id")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="transitGatewayId")
    def transit_gateway_id(self) -> pulumi.Output[Optional[str]]:
        """
        Identifier of an EC2 Transit Gateway.
        """
        return pulumi.get(self, "transit_gateway_id")

    @property
    @pulumi.getter(name="vpcEndpointId")
    def vpc_endpoint_id(self) -> pulumi.Output[Optional[str]]:
        """
        Identifier of a VPC Endpoint.
        """
        return pulumi.get(self, "vpc_endpoint_id")

    @property
    @pulumi.getter(name="vpcPeeringConnectionId")
    def vpc_peering_connection_id(self) -> pulumi.Output[Optional[str]]:
        """
        Identifier of a VPC peering connection.
        """
        return pulumi.get(self, "vpc_peering_connection_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

