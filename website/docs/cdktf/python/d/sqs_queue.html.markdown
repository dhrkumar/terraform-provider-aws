---
subcategory: "SQS (Simple Queue)"
layout: "aws"
page_title: "AWS: aws_sqs_queue"
description: |-
  Get information on an Amazon Simple Queue Service (SQS) Queue
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_sqs_queue

Use this data source to get the ARN and URL of queue in AWS Simple Queue Service (SQS).
By using this data source, you can reference SQS queues without having to hardcode
the ARNs as input.

~> **NOTE:** To use this data source, you must have the `sqs:GetQueueAttributes` and `sqs:GetQueueURL` permissions.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_sqs_queue import DataAwsSqsQueue
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataAwsSqsQueue(self, "example",
            name="queue"
        )
```

## Argument Reference

This data source supports the following arguments:

* `name` - (Required) Name of the queue to match.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `arn` - ARN of the queue.
* `url` - URL of the queue.
* `tags` - Map of tags for the resource.

<!-- cache-key: cdktf-0.20.8 input-24d5a97ed7fbead815006896ee90ce8d0a4e26e3a321d0c96d945d459d82dbe3 -->