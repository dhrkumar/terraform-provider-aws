---
subcategory: "ECR (Elastic Container Registry)"
layout: "aws"
page_title: "AWS: aws_ecr_repositories"
description: |-
  Terraform data source for providing information on AWS ECR (Elastic Container Registry) Repositories.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_ecr_repositories

Terraform data source for providing information on AWS ECR (Elastic Container Registry) Repositories.

## Example Usage

### Basic Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_ecr_repositories import DataAwsEcrRepositories
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataAwsEcrRepositories(self, "example")
```

## Argument Reference

This data source does not support any arguments.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `id` - AWS Region.
* `names` - A list if AWS Elastic Container Registries for the region.

<!-- cache-key: cdktf-0.20.8 input-eb5beff7f44054362324d03e57bc31451600ce4c220d3e421ee95c6c6547873f -->