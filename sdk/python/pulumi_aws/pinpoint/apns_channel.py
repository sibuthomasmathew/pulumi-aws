# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['ApnsChannel']


class ApnsChannel(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 bundle_id: Optional[pulumi.Input[str]] = None,
                 certificate: Optional[pulumi.Input[str]] = None,
                 default_authentication_method: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 team_id: Optional[pulumi.Input[str]] = None,
                 token_key: Optional[pulumi.Input[str]] = None,
                 token_key_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a Pinpoint APNs Channel resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        app = aws.pinpoint.App("app")
        apns = aws.pinpoint.ApnsChannel("apns",
            application_id=app.application_id,
            certificate=(lambda path: open(path).read())("./certificate.pem"),
            private_key=(lambda path: open(path).read())("./private_key.key"))
        ```

        ## Import

        Pinpoint APNs Channel can be imported using the `application-id`, e.g.

        ```sh
         $ pulumi import aws:pinpoint/apnsChannel:ApnsChannel apns application-id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: The application ID.
        :param pulumi.Input[str] bundle_id: The ID assigned to your iOS app. To find this value, choose Certificates, IDs & Profiles, choose App IDs in the Identifiers section, and choose your app.
        :param pulumi.Input[str] certificate: The pem encoded TLS Certificate from Apple.
        :param pulumi.Input[str] default_authentication_method: The default authentication method used for APNs.
               __NOTE__: Amazon Pinpoint uses this default for every APNs push notification that you send using the console.
               You can override the default when you send a message programmatically using the Amazon Pinpoint API, the AWS CLI, or an AWS SDK.
               If your default authentication type fails, Amazon Pinpoint doesn't attempt to use the other authentication type.
        :param pulumi.Input[bool] enabled: Whether the channel is enabled or disabled. Defaults to `true`.
        :param pulumi.Input[str] private_key: The Certificate Private Key file (ie. `.key` file).
        :param pulumi.Input[str] team_id: The ID assigned to your Apple developer account team. This value is provided on the Membership page.
        :param pulumi.Input[str] token_key: The `.p8` file that you download from your Apple developer account when you create an authentication key.
        :param pulumi.Input[str] token_key_id: The ID assigned to your signing key. To find this value, choose Certificates, IDs & Profiles, and choose your key in the Keys section.
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

            if application_id is None and not opts.urn:
                raise TypeError("Missing required property 'application_id'")
            __props__['application_id'] = application_id
            __props__['bundle_id'] = bundle_id
            __props__['certificate'] = certificate
            __props__['default_authentication_method'] = default_authentication_method
            __props__['enabled'] = enabled
            __props__['private_key'] = private_key
            __props__['team_id'] = team_id
            __props__['token_key'] = token_key
            __props__['token_key_id'] = token_key_id
        super(ApnsChannel, __self__).__init__(
            'aws:pinpoint/apnsChannel:ApnsChannel',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            application_id: Optional[pulumi.Input[str]] = None,
            bundle_id: Optional[pulumi.Input[str]] = None,
            certificate: Optional[pulumi.Input[str]] = None,
            default_authentication_method: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            private_key: Optional[pulumi.Input[str]] = None,
            team_id: Optional[pulumi.Input[str]] = None,
            token_key: Optional[pulumi.Input[str]] = None,
            token_key_id: Optional[pulumi.Input[str]] = None) -> 'ApnsChannel':
        """
        Get an existing ApnsChannel resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: The application ID.
        :param pulumi.Input[str] bundle_id: The ID assigned to your iOS app. To find this value, choose Certificates, IDs & Profiles, choose App IDs in the Identifiers section, and choose your app.
        :param pulumi.Input[str] certificate: The pem encoded TLS Certificate from Apple.
        :param pulumi.Input[str] default_authentication_method: The default authentication method used for APNs.
               __NOTE__: Amazon Pinpoint uses this default for every APNs push notification that you send using the console.
               You can override the default when you send a message programmatically using the Amazon Pinpoint API, the AWS CLI, or an AWS SDK.
               If your default authentication type fails, Amazon Pinpoint doesn't attempt to use the other authentication type.
        :param pulumi.Input[bool] enabled: Whether the channel is enabled or disabled. Defaults to `true`.
        :param pulumi.Input[str] private_key: The Certificate Private Key file (ie. `.key` file).
        :param pulumi.Input[str] team_id: The ID assigned to your Apple developer account team. This value is provided on the Membership page.
        :param pulumi.Input[str] token_key: The `.p8` file that you download from your Apple developer account when you create an authentication key.
        :param pulumi.Input[str] token_key_id: The ID assigned to your signing key. To find this value, choose Certificates, IDs & Profiles, and choose your key in the Keys section.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["application_id"] = application_id
        __props__["bundle_id"] = bundle_id
        __props__["certificate"] = certificate
        __props__["default_authentication_method"] = default_authentication_method
        __props__["enabled"] = enabled
        __props__["private_key"] = private_key
        __props__["team_id"] = team_id
        __props__["token_key"] = token_key
        __props__["token_key_id"] = token_key_id
        return ApnsChannel(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> pulumi.Output[str]:
        """
        The application ID.
        """
        return pulumi.get(self, "application_id")

    @property
    @pulumi.getter(name="bundleId")
    def bundle_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID assigned to your iOS app. To find this value, choose Certificates, IDs & Profiles, choose App IDs in the Identifiers section, and choose your app.
        """
        return pulumi.get(self, "bundle_id")

    @property
    @pulumi.getter
    def certificate(self) -> pulumi.Output[Optional[str]]:
        """
        The pem encoded TLS Certificate from Apple.
        """
        return pulumi.get(self, "certificate")

    @property
    @pulumi.getter(name="defaultAuthenticationMethod")
    def default_authentication_method(self) -> pulumi.Output[Optional[str]]:
        """
        The default authentication method used for APNs.
        __NOTE__: Amazon Pinpoint uses this default for every APNs push notification that you send using the console.
        You can override the default when you send a message programmatically using the Amazon Pinpoint API, the AWS CLI, or an AWS SDK.
        If your default authentication type fails, Amazon Pinpoint doesn't attempt to use the other authentication type.
        """
        return pulumi.get(self, "default_authentication_method")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the channel is enabled or disabled. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> pulumi.Output[Optional[str]]:
        """
        The Certificate Private Key file (ie. `.key` file).
        """
        return pulumi.get(self, "private_key")

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID assigned to your Apple developer account team. This value is provided on the Membership page.
        """
        return pulumi.get(self, "team_id")

    @property
    @pulumi.getter(name="tokenKey")
    def token_key(self) -> pulumi.Output[Optional[str]]:
        """
        The `.p8` file that you download from your Apple developer account when you create an authentication key.
        """
        return pulumi.get(self, "token_key")

    @property
    @pulumi.getter(name="tokenKeyId")
    def token_key_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID assigned to your signing key. To find this value, choose Certificates, IDs & Profiles, and choose your key in the Keys section.
        """
        return pulumi.get(self, "token_key_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

