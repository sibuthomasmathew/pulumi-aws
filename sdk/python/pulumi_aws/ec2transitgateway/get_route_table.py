# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = [
    'GetRouteTableResult',
    'AwaitableGetRouteTableResult',
    'get_route_table',
]

@pulumi.output_type
class GetRouteTableResult:
    """
    A collection of values returned by getRouteTable.
    """
    def __init__(__self__, arn=None, default_association_route_table=None, default_propagation_route_table=None, filters=None, id=None, tags=None, transit_gateway_id=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if default_association_route_table and not isinstance(default_association_route_table, bool):
            raise TypeError("Expected argument 'default_association_route_table' to be a bool")
        pulumi.set(__self__, "default_association_route_table", default_association_route_table)
        if default_propagation_route_table and not isinstance(default_propagation_route_table, bool):
            raise TypeError("Expected argument 'default_propagation_route_table' to be a bool")
        pulumi.set(__self__, "default_propagation_route_table", default_propagation_route_table)
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if transit_gateway_id and not isinstance(transit_gateway_id, str):
            raise TypeError("Expected argument 'transit_gateway_id' to be a str")
        pulumi.set(__self__, "transit_gateway_id", transit_gateway_id)

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        EC2 Transit Gateway Route Table Amazon Resource Name (ARN).
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="defaultAssociationRouteTable")
    def default_association_route_table(self) -> bool:
        """
        Boolean whether this is the default association route table for the EC2 Transit Gateway
        """
        return pulumi.get(self, "default_association_route_table")

    @property
    @pulumi.getter(name="defaultPropagationRouteTable")
    def default_propagation_route_table(self) -> bool:
        """
        Boolean whether this is the default propagation route table for the EC2 Transit Gateway
        """
        return pulumi.get(self, "default_propagation_route_table")

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetRouteTableFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        EC2 Transit Gateway Route Table identifier
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        Key-value tags for the EC2 Transit Gateway Route Table
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="transitGatewayId")
    def transit_gateway_id(self) -> str:
        """
        EC2 Transit Gateway identifier
        """
        return pulumi.get(self, "transit_gateway_id")


class AwaitableGetRouteTableResult(GetRouteTableResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRouteTableResult(
            arn=self.arn,
            default_association_route_table=self.default_association_route_table,
            default_propagation_route_table=self.default_propagation_route_table,
            filters=self.filters,
            id=self.id,
            tags=self.tags,
            transit_gateway_id=self.transit_gateway_id)


def get_route_table(filters: Optional[Sequence[pulumi.InputType['GetRouteTableFilterArgs']]] = None,
                    id: Optional[str] = None,
                    tags: Optional[Mapping[str, str]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRouteTableResult:
    """
    Get information on an EC2 Transit Gateway Route Table.

    ## Example Usage
    ### By Identifier

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.ec2transitgateway.get_route_table(id="tgw-rtb-12345678")
    ```


    :param Sequence[pulumi.InputType['GetRouteTableFilterArgs']] filters: One or more configuration blocks containing name-values filters. Detailed below.
    :param str id: Identifier of the EC2 Transit Gateway Route Table.
    :param Mapping[str, str] tags: Key-value tags for the EC2 Transit Gateway Route Table
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['id'] = id
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ec2transitgateway/getRouteTable:getRouteTable', __args__, opts=opts, typ=GetRouteTableResult).value

    return AwaitableGetRouteTableResult(
        arn=__ret__.arn,
        default_association_route_table=__ret__.default_association_route_table,
        default_propagation_route_table=__ret__.default_propagation_route_table,
        filters=__ret__.filters,
        id=__ret__.id,
        tags=__ret__.tags,
        transit_gateway_id=__ret__.transit_gateway_id)
