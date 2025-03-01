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

__all__ = ['SigningProfile']


class SigningProfile(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 platform_id: Optional[pulumi.Input[str]] = None,
                 signature_validity_period: Optional[pulumi.Input[pulumi.InputType['SigningProfileSignatureValidityPeriodArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates a Signer Signing Profile. A signing profile contains information about the code signing configuration parameters that can be used by a given code signing user.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        test_sp = aws.signer.SigningProfile("testSp", platform_id="AWSLambda-SHA384-ECDSA")
        prod_sp = aws.signer.SigningProfile("prodSp",
            name_prefix="prod_sp_",
            platform_id="AWSLambda-SHA384-ECDSA",
            signature_validity_period=aws.signer.SigningProfileSignatureValidityPeriodArgs(
                type="YEARS",
                value=5,
            ),
            tags={
                "tag1": "value1",
                "tag2": "value2",
            })
        ```

        ## Import

        Signer signing profiles can be imported using the `name`, e.g.

        ```sh
         $ pulumi import aws:signer/signingProfile:SigningProfile test_signer_signing_profile test_sp_DdW3Mk1foYL88fajut4mTVFGpuwfd4ACO6ANL0D1uIj7lrn8adK
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: A unique signing profile name. By default generated by the provider. Signing profile names are immutable and cannot be reused after canceled.
        :param pulumi.Input[str] name_prefix: A signing profile name prefix. The provider will generate a unique suffix. Conflicts with `name`.
        :param pulumi.Input[str] platform_id: The ID of the platform that is used by the target signing profile.
        :param pulumi.Input[pulumi.InputType['SigningProfileSignatureValidityPeriodArgs']] signature_validity_period: The validity period for a signing job.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A list of tags associated with the signing profile.
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

            __props__['name'] = name
            __props__['name_prefix'] = name_prefix
            if platform_id is None and not opts.urn:
                raise TypeError("Missing required property 'platform_id'")
            __props__['platform_id'] = platform_id
            __props__['signature_validity_period'] = signature_validity_period
            __props__['tags'] = tags
            __props__['arn'] = None
            __props__['platform_display_name'] = None
            __props__['revocation_records'] = None
            __props__['status'] = None
            __props__['version'] = None
            __props__['version_arn'] = None
        super(SigningProfile, __self__).__init__(
            'aws:signer/signingProfile:SigningProfile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            name_prefix: Optional[pulumi.Input[str]] = None,
            platform_display_name: Optional[pulumi.Input[str]] = None,
            platform_id: Optional[pulumi.Input[str]] = None,
            revocation_records: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SigningProfileRevocationRecordArgs']]]]] = None,
            signature_validity_period: Optional[pulumi.Input[pulumi.InputType['SigningProfileSignatureValidityPeriodArgs']]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            version: Optional[pulumi.Input[str]] = None,
            version_arn: Optional[pulumi.Input[str]] = None) -> 'SigningProfile':
        """
        Get an existing SigningProfile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) for the signing profile.
        :param pulumi.Input[str] name: A unique signing profile name. By default generated by the provider. Signing profile names are immutable and cannot be reused after canceled.
        :param pulumi.Input[str] name_prefix: A signing profile name prefix. The provider will generate a unique suffix. Conflicts with `name`.
        :param pulumi.Input[str] platform_display_name: A human-readable name for the signing platform associated with the signing profile.
        :param pulumi.Input[str] platform_id: The ID of the platform that is used by the target signing profile.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SigningProfileRevocationRecordArgs']]]] revocation_records: Revocation information for a signing profile.
        :param pulumi.Input[pulumi.InputType['SigningProfileSignatureValidityPeriodArgs']] signature_validity_period: The validity period for a signing job.
        :param pulumi.Input[str] status: The status of the target signing profile.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A list of tags associated with the signing profile.
        :param pulumi.Input[str] version: The current version of the signing profile.
        :param pulumi.Input[str] version_arn: The signing profile ARN, including the profile version.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["name"] = name
        __props__["name_prefix"] = name_prefix
        __props__["platform_display_name"] = platform_display_name
        __props__["platform_id"] = platform_id
        __props__["revocation_records"] = revocation_records
        __props__["signature_validity_period"] = signature_validity_period
        __props__["status"] = status
        __props__["tags"] = tags
        __props__["version"] = version
        __props__["version_arn"] = version_arn
        return SigningProfile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) for the signing profile.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A unique signing profile name. By default generated by the provider. Signing profile names are immutable and cannot be reused after canceled.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> pulumi.Output[Optional[str]]:
        """
        A signing profile name prefix. The provider will generate a unique suffix. Conflicts with `name`.
        """
        return pulumi.get(self, "name_prefix")

    @property
    @pulumi.getter(name="platformDisplayName")
    def platform_display_name(self) -> pulumi.Output[str]:
        """
        A human-readable name for the signing platform associated with the signing profile.
        """
        return pulumi.get(self, "platform_display_name")

    @property
    @pulumi.getter(name="platformId")
    def platform_id(self) -> pulumi.Output[str]:
        """
        The ID of the platform that is used by the target signing profile.
        """
        return pulumi.get(self, "platform_id")

    @property
    @pulumi.getter(name="revocationRecords")
    def revocation_records(self) -> pulumi.Output[Sequence['outputs.SigningProfileRevocationRecord']]:
        """
        Revocation information for a signing profile.
        """
        return pulumi.get(self, "revocation_records")

    @property
    @pulumi.getter(name="signatureValidityPeriod")
    def signature_validity_period(self) -> pulumi.Output['outputs.SigningProfileSignatureValidityPeriod']:
        """
        The validity period for a signing job.
        """
        return pulumi.get(self, "signature_validity_period")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the target signing profile.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A list of tags associated with the signing profile.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[str]:
        """
        The current version of the signing profile.
        """
        return pulumi.get(self, "version")

    @property
    @pulumi.getter(name="versionArn")
    def version_arn(self) -> pulumi.Output[str]:
        """
        The signing profile ARN, including the profile version.
        """
        return pulumi.get(self, "version_arn")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

