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

__all__ = ['EndpointGroup']


class EndpointGroup(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 endpoint_configurations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointGroupEndpointConfigurationArgs']]]]] = None,
                 endpoint_group_region: Optional[pulumi.Input[str]] = None,
                 health_check_interval_seconds: Optional[pulumi.Input[int]] = None,
                 health_check_path: Optional[pulumi.Input[str]] = None,
                 health_check_port: Optional[pulumi.Input[int]] = None,
                 health_check_protocol: Optional[pulumi.Input[str]] = None,
                 listener_arn: Optional[pulumi.Input[str]] = None,
                 port_overrides: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointGroupPortOverrideArgs']]]]] = None,
                 threshold_count: Optional[pulumi.Input[int]] = None,
                 traffic_dial_percentage: Optional[pulumi.Input[float]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a Global Accelerator endpoint group.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.globalaccelerator.EndpointGroup("example",
            listener_arn=aws_globalaccelerator_listener["example"]["id"],
            endpoint_configurations=[aws.globalaccelerator.EndpointGroupEndpointConfigurationArgs(
                endpoint_id=aws_lb["example"]["arn"],
                weight=100,
            )])
        ```

        ## Import

        Global Accelerator endpoint groups can be imported using the `id`, e.g.

        ```sh
         $ pulumi import aws:globalaccelerator/endpointGroup:EndpointGroup example arn:aws:globalaccelerator::111111111111:accelerator/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/listener/xxxxxxx/endpoint-group/xxxxxxxx
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointGroupEndpointConfigurationArgs']]]] endpoint_configurations: The list of endpoint objects. Fields documented below.
        :param pulumi.Input[str] endpoint_group_region: The name of the AWS Region where the endpoint group is located.
        :param pulumi.Input[int] health_check_interval_seconds: The time—10 seconds or 30 seconds—between each health check for an endpoint. The default value is 30.
        :param pulumi.Input[str] health_check_path: If the protocol is HTTP/S, then this specifies the path that is the destination for health check targets. The default value is slash (`/`). the provider will only perform drift detection of its value when present in a configuration.
        :param pulumi.Input[int] health_check_port: The port that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default port is the listener port that this endpoint group is associated with. If listener port is a list of ports, Global Accelerator uses the first port in the list.
               the provider will only perform drift detection of its value when present in a configuration.
        :param pulumi.Input[str] health_check_protocol: The protocol that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default value is TCP.
        :param pulumi.Input[str] listener_arn: The Amazon Resource Name (ARN) of the listener.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointGroupPortOverrideArgs']]]] port_overrides: Override specific listener ports used to route traffic to endpoints that are part of this endpoint group. Fields documented below.
        :param pulumi.Input[int] threshold_count: The number of consecutive health checks required to set the state of a healthy endpoint to unhealthy, or to set an unhealthy endpoint to healthy. The default value is 3.
        :param pulumi.Input[float] traffic_dial_percentage: The percentage of traffic to send to an AWS Region. Additional traffic is distributed to other endpoint groups for this listener. The default value is 100.
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

            __props__['endpoint_configurations'] = endpoint_configurations
            __props__['endpoint_group_region'] = endpoint_group_region
            __props__['health_check_interval_seconds'] = health_check_interval_seconds
            __props__['health_check_path'] = health_check_path
            __props__['health_check_port'] = health_check_port
            __props__['health_check_protocol'] = health_check_protocol
            if listener_arn is None and not opts.urn:
                raise TypeError("Missing required property 'listener_arn'")
            __props__['listener_arn'] = listener_arn
            __props__['port_overrides'] = port_overrides
            __props__['threshold_count'] = threshold_count
            __props__['traffic_dial_percentage'] = traffic_dial_percentage
            __props__['arn'] = None
        super(EndpointGroup, __self__).__init__(
            'aws:globalaccelerator/endpointGroup:EndpointGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            endpoint_configurations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointGroupEndpointConfigurationArgs']]]]] = None,
            endpoint_group_region: Optional[pulumi.Input[str]] = None,
            health_check_interval_seconds: Optional[pulumi.Input[int]] = None,
            health_check_path: Optional[pulumi.Input[str]] = None,
            health_check_port: Optional[pulumi.Input[int]] = None,
            health_check_protocol: Optional[pulumi.Input[str]] = None,
            listener_arn: Optional[pulumi.Input[str]] = None,
            port_overrides: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointGroupPortOverrideArgs']]]]] = None,
            threshold_count: Optional[pulumi.Input[int]] = None,
            traffic_dial_percentage: Optional[pulumi.Input[float]] = None) -> 'EndpointGroup':
        """
        Get an existing EndpointGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) of the endpoint group.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointGroupEndpointConfigurationArgs']]]] endpoint_configurations: The list of endpoint objects. Fields documented below.
        :param pulumi.Input[str] endpoint_group_region: The name of the AWS Region where the endpoint group is located.
        :param pulumi.Input[int] health_check_interval_seconds: The time—10 seconds or 30 seconds—between each health check for an endpoint. The default value is 30.
        :param pulumi.Input[str] health_check_path: If the protocol is HTTP/S, then this specifies the path that is the destination for health check targets. The default value is slash (`/`). the provider will only perform drift detection of its value when present in a configuration.
        :param pulumi.Input[int] health_check_port: The port that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default port is the listener port that this endpoint group is associated with. If listener port is a list of ports, Global Accelerator uses the first port in the list.
               the provider will only perform drift detection of its value when present in a configuration.
        :param pulumi.Input[str] health_check_protocol: The protocol that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default value is TCP.
        :param pulumi.Input[str] listener_arn: The Amazon Resource Name (ARN) of the listener.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointGroupPortOverrideArgs']]]] port_overrides: Override specific listener ports used to route traffic to endpoints that are part of this endpoint group. Fields documented below.
        :param pulumi.Input[int] threshold_count: The number of consecutive health checks required to set the state of a healthy endpoint to unhealthy, or to set an unhealthy endpoint to healthy. The default value is 3.
        :param pulumi.Input[float] traffic_dial_percentage: The percentage of traffic to send to an AWS Region. Additional traffic is distributed to other endpoint groups for this listener. The default value is 100.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["endpoint_configurations"] = endpoint_configurations
        __props__["endpoint_group_region"] = endpoint_group_region
        __props__["health_check_interval_seconds"] = health_check_interval_seconds
        __props__["health_check_path"] = health_check_path
        __props__["health_check_port"] = health_check_port
        __props__["health_check_protocol"] = health_check_protocol
        __props__["listener_arn"] = listener_arn
        __props__["port_overrides"] = port_overrides
        __props__["threshold_count"] = threshold_count
        __props__["traffic_dial_percentage"] = traffic_dial_percentage
        return EndpointGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of the endpoint group.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="endpointConfigurations")
    def endpoint_configurations(self) -> pulumi.Output[Optional[Sequence['outputs.EndpointGroupEndpointConfiguration']]]:
        """
        The list of endpoint objects. Fields documented below.
        """
        return pulumi.get(self, "endpoint_configurations")

    @property
    @pulumi.getter(name="endpointGroupRegion")
    def endpoint_group_region(self) -> pulumi.Output[str]:
        """
        The name of the AWS Region where the endpoint group is located.
        """
        return pulumi.get(self, "endpoint_group_region")

    @property
    @pulumi.getter(name="healthCheckIntervalSeconds")
    def health_check_interval_seconds(self) -> pulumi.Output[Optional[int]]:
        """
        The time—10 seconds or 30 seconds—between each health check for an endpoint. The default value is 30.
        """
        return pulumi.get(self, "health_check_interval_seconds")

    @property
    @pulumi.getter(name="healthCheckPath")
    def health_check_path(self) -> pulumi.Output[str]:
        """
        If the protocol is HTTP/S, then this specifies the path that is the destination for health check targets. The default value is slash (`/`). the provider will only perform drift detection of its value when present in a configuration.
        """
        return pulumi.get(self, "health_check_path")

    @property
    @pulumi.getter(name="healthCheckPort")
    def health_check_port(self) -> pulumi.Output[int]:
        """
        The port that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default port is the listener port that this endpoint group is associated with. If listener port is a list of ports, Global Accelerator uses the first port in the list.
        the provider will only perform drift detection of its value when present in a configuration.
        """
        return pulumi.get(self, "health_check_port")

    @property
    @pulumi.getter(name="healthCheckProtocol")
    def health_check_protocol(self) -> pulumi.Output[Optional[str]]:
        """
        The protocol that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default value is TCP.
        """
        return pulumi.get(self, "health_check_protocol")

    @property
    @pulumi.getter(name="listenerArn")
    def listener_arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of the listener.
        """
        return pulumi.get(self, "listener_arn")

    @property
    @pulumi.getter(name="portOverrides")
    def port_overrides(self) -> pulumi.Output[Optional[Sequence['outputs.EndpointGroupPortOverride']]]:
        """
        Override specific listener ports used to route traffic to endpoints that are part of this endpoint group. Fields documented below.
        """
        return pulumi.get(self, "port_overrides")

    @property
    @pulumi.getter(name="thresholdCount")
    def threshold_count(self) -> pulumi.Output[Optional[int]]:
        """
        The number of consecutive health checks required to set the state of a healthy endpoint to unhealthy, or to set an unhealthy endpoint to healthy. The default value is 3.
        """
        return pulumi.get(self, "threshold_count")

    @property
    @pulumi.getter(name="trafficDialPercentage")
    def traffic_dial_percentage(self) -> pulumi.Output[Optional[float]]:
        """
        The percentage of traffic to send to an AWS Region. Additional traffic is distributed to other endpoint groups for this listener. The default value is 100.
        """
        return pulumi.get(self, "traffic_dial_percentage")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

