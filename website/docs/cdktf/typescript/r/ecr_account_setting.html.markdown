---
subcategory: "ECR (Elastic Container Registry)"
layout: "aws"
page_title: "AWS: aws_ecr_account_setting"
description: |-
  Provides a resource to manage AWS ECR Basic Scan Type
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_ecr_account_settings

Provides a resource to manage AWS ECR Basic Scan Type

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { EcrAccountSetting } from "./.gen/providers/aws/";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new EcrAccountSetting(this, "foo", {
      name: "BASIC_SCAN_TYPE_VERSION",
      value: "CLAIR",
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `name` - (Required) The name of the ECR Scan Type. This should be `BASIC_SCAN_TYPE_VERSION`.
* `value` - (Required) The value of the ECR Scan Type. This can be `AWS_NATIVE` or `CLAIR`.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The ID of the ECR Scan Type (Same as the `name`)
* `name` - The Name of the ECR Scan Type
* `value` - The Value of the ECR Scan Type

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import ECR Scan Type using the `name`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { EcrAccountSetting } from "./.gen/providers/aws/";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    EcrAccountSetting.generateConfigForImport(
      this,
      "foo",
      "BASIC_SCAN_TYPE_VERSION"
    );
  }
}

```

Using `terraform import`, import EMR Security Configurations using the `name`. For example:

```console
% terraform import aws_ecr_account_setting.foo BASIC_SCAN_TYPE_VERSION
```

<!-- cache-key: cdktf-0.20.8 input-5255caeda1c78ccd9ae83c476c5a347f5714117c20331aa8b73ddef8c3ad0988 -->