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

__all__ = ['SecretRotation']


class SecretRotation(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 rotation_lambda_arn: Optional[pulumi.Input[str]] = None,
                 rotation_rules: Optional[pulumi.Input[pulumi.InputType['SecretRotationRotationRulesArgs']]] = None,
                 secret_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a resource to manage AWS Secrets Manager secret rotation. To manage a secret, see the `secretsmanager.Secret` resource. To manage a secret value, see the `secretsmanager.SecretVersion` resource.

        ## Example Usage
        ### Basic

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.secretsmanager.SecretRotation("example",
            secret_id=aws_secretsmanager_secret["example"]["id"],
            rotation_lambda_arn=aws_lambda_function["example"]["arn"],
            rotation_rules=aws.secretsmanager.SecretRotationRotationRulesArgs(
                automatically_after_days=30,
            ))
        ```
        ### Rotation Configuration

        To enable automatic secret rotation, the Secrets Manager service requires usage of a Lambda function. The [Rotate Secrets section in the Secrets Manager User Guide](https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotating-secrets.html) provides additional information about deploying a prebuilt Lambda functions for supported credential rotation (e.g. RDS) or deploying a custom Lambda function.

        > **NOTE:** Configuring rotation causes the secret to rotate once as soon as you enable rotation. Before you do this, you must ensure that all of your applications that use the credentials stored in the secret are updated to retrieve the secret from AWS Secrets Manager. The old credentials might no longer be usable after the initial rotation and any applications that you fail to update will break as soon as the old credentials are no longer valid.

        > **NOTE:** If you cancel a rotation that is in progress (by removing the `rotation` configuration), it can leave the VersionStage labels in an unexpected state. Depending on what step of the rotation was in progress, you might need to remove the staging label AWSPENDING from the partially created version, specified by the SecretVersionId response value. You should also evaluate the partially rotated new version to see if it should be deleted, which you can do by removing all staging labels from the new version's VersionStage field.

        ## Import

        `aws_secretsmanager_secret_rotation` can be imported by using the secret Amazon Resource Name (ARN), e.g.

        ```sh
         $ pulumi import aws:secretsmanager/secretRotation:SecretRotation example arn:aws:secretsmanager:us-east-1:123456789012:secret:example-123456
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] rotation_lambda_arn: Specifies the ARN of the Lambda function that can rotate the secret.
        :param pulumi.Input[pulumi.InputType['SecretRotationRotationRulesArgs']] rotation_rules: A structure that defines the rotation configuration for this secret. Defined below.
        :param pulumi.Input[str] secret_id: Specifies the secret to which you want to add a new version. You can specify either the Amazon Resource Name (ARN) or the friendly name of the secret. The secret must already exist.
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

            if rotation_lambda_arn is None and not opts.urn:
                raise TypeError("Missing required property 'rotation_lambda_arn'")
            __props__['rotation_lambda_arn'] = rotation_lambda_arn
            if rotation_rules is None and not opts.urn:
                raise TypeError("Missing required property 'rotation_rules'")
            __props__['rotation_rules'] = rotation_rules
            if secret_id is None and not opts.urn:
                raise TypeError("Missing required property 'secret_id'")
            __props__['secret_id'] = secret_id
            __props__['tags'] = tags
            __props__['rotation_enabled'] = None
        super(SecretRotation, __self__).__init__(
            'aws:secretsmanager/secretRotation:SecretRotation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            rotation_enabled: Optional[pulumi.Input[bool]] = None,
            rotation_lambda_arn: Optional[pulumi.Input[str]] = None,
            rotation_rules: Optional[pulumi.Input[pulumi.InputType['SecretRotationRotationRulesArgs']]] = None,
            secret_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'SecretRotation':
        """
        Get an existing SecretRotation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] rotation_enabled: Specifies whether automatic rotation is enabled for this secret.
        :param pulumi.Input[str] rotation_lambda_arn: Specifies the ARN of the Lambda function that can rotate the secret.
        :param pulumi.Input[pulumi.InputType['SecretRotationRotationRulesArgs']] rotation_rules: A structure that defines the rotation configuration for this secret. Defined below.
        :param pulumi.Input[str] secret_id: Specifies the secret to which you want to add a new version. You can specify either the Amazon Resource Name (ARN) or the friendly name of the secret. The secret must already exist.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["rotation_enabled"] = rotation_enabled
        __props__["rotation_lambda_arn"] = rotation_lambda_arn
        __props__["rotation_rules"] = rotation_rules
        __props__["secret_id"] = secret_id
        __props__["tags"] = tags
        return SecretRotation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="rotationEnabled")
    def rotation_enabled(self) -> pulumi.Output[bool]:
        """
        Specifies whether automatic rotation is enabled for this secret.
        """
        return pulumi.get(self, "rotation_enabled")

    @property
    @pulumi.getter(name="rotationLambdaArn")
    def rotation_lambda_arn(self) -> pulumi.Output[str]:
        """
        Specifies the ARN of the Lambda function that can rotate the secret.
        """
        return pulumi.get(self, "rotation_lambda_arn")

    @property
    @pulumi.getter(name="rotationRules")
    def rotation_rules(self) -> pulumi.Output['outputs.SecretRotationRotationRules']:
        """
        A structure that defines the rotation configuration for this secret. Defined below.
        """
        return pulumi.get(self, "rotation_rules")

    @property
    @pulumi.getter(name="secretId")
    def secret_id(self) -> pulumi.Output[str]:
        """
        Specifies the secret to which you want to add a new version. You can specify either the Amazon Resource Name (ARN) or the friendly name of the secret. The secret must already exist.
        """
        return pulumi.get(self, "secret_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

