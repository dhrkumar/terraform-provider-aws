---
subcategory: "EC2 (Elastic Compute Cloud)"
layout: "aws"
page_title: "AWS: aws_ec2_capacity_block_reservation"
description: |-
  Provides an EC2 Capacity Block Reservation. This allows you to purchase capacity block for your Amazon EC2 instances in a specific Availability Zone for machine learning (ML) Workloads.
---

# Resource: aws_ec2_capacity_block_reservation

Provides an EC2 Capacity Block Reservation. This allows you to purchase capacity block for your Amazon EC2 instances in a specific Availability Zone for machine learning (ML) Workloads.

~> **NOTE:** Once created, a reservation is valid for the `duration` of the provided `capacity_block_offering_id` and cannot be deleted. Performing a `destroy` will only remove the resource from state. For more information see [EC2 Capacity Block Reservation Documentation](https://aws.amazon.com/ec2/instance-types/p5/) and [PurchaseReservedDBInstancesOffering](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/capacity-blocks-pricing-billing.html).

~> **NOTE:** Due to the expense of testing this resource, we provide it as best effort. If you find it useful, and have the ability to help test or notice issues, consider reaching out to us on [GitHub](https://github.com/hashicorp/terraform-provider-aws).

## Example Usage

```terraform

data "aws_ec2_capacity_block_offering" "test" {
  capacity_duration_hours = 24
  end_date_range          = "2024-05-30T15:04:05Z"
  instance_count          = 1
  instance_type           = "p4d.24xlarge"
  start_date_range        = "2024-04-28T15:04:05Z"
}

resource "aws_ec2_capacity_block_reservation" "example" {
  capacity_block_offering_id = data.aws_ec2_capacity_block_offering.test.capacity_block_offering_id
  instance_platform          = "Linux/UNIX"
  tags = {
    "Environment" = "dev"
  }
}
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `capacity_block_offering_id` - (Required) The Capacity Block Reservation ID.
* `instance_platform` - (Required) The type of operating system for which to reserve capacity. Valid options are `Linux/UNIX`, `Red Hat Enterprise Linux`, `SUSE Linux`, `Windows`, `Windows with SQL Server`, `Windows with SQL Server Enterprise`, `Windows with SQL Server Standard` or `Windows with SQL Server Web`.
* `tags` - (Optional) A map of tags to assign to the resource. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - The ARN of the reservation.
* `availability_zone` - The Availability Zone in which to create the Capacity Block Reservation.
* `created_date` - The date and time at which the Capacity Block Reservation was created.
* `ebs_optimized` - Indicates whether the Capacity Reservation supports EBS-optimized instances.
* `end_date` - The date and time at which the Capacity Block Reservation expires. When a Capacity Block Reservation expires, the reserved capacity is released and you can no longer launch instances into it. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
* `end_date_type` - Indicates the way in which the Capacity Reservation ends.
* `id` - The ID of the Capacity Block Reservation.
* `instance_count` - The number of instances for which to reserve capacity.
* `instance_type` - The instance type for which to reserve capacity.
* `outpost_arn` - The ARN of the Outpost on which to create the Capacity Block Reservation.
* `placement_group_arn` - The ARN of the placement group in which to create the Capacity Block Reservation.
* `reservation_type` - The type of Capacity Reservation.
* `start_date` - The date and time at which the Capacity Block Reservation starts. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
* `tenancy` - Indicates the tenancy of the Capacity Block Reservation. Specify either `default` or `dedicated`.
* `tags_all` - A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block)
