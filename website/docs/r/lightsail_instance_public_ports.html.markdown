---
subcategory: "Lightsail"
layout: "aws"
page_title: "AWS: aws_lightsail_instance_public_ports"
description: |-
  Manages public ports for a Lightsail instance.
---

# Resource: aws_lightsail_instance_public_ports

Manages public ports for a Lightsail instance. Use this resource to open ports for a specific Amazon Lightsail instance and specify the IP addresses allowed to connect to the instance through the ports and the protocol.

-> See [What is Amazon Lightsail?](https://lightsail.aws.amazon.com/ls/docs/getting-started/article/what-is-amazon-lightsail) for more information.

~> **Note:** Lightsail is currently only supported in a limited number of AWS Regions, please see ["Regions and Availability Zones in Amazon Lightsail"](https://lightsail.aws.amazon.com/ls/docs/overview/article/understanding-regions-and-availability-zones-in-amazon-lightsail) for more details.

## Example Usage

```terraform
data "aws_availability_zones" "available" {
  state = "available"

  filter {
    name   = "opt-in-status"
    values = ["opt-in-not-required"]
  }
}

resource "aws_lightsail_instance" "example" {
  name              = "example-instance"
  availability_zone = data.aws_availability_zones.available.names[0]
  blueprint_id      = "amazon_linux_2"
  bundle_id         = "nano_3_0"
}

resource "aws_lightsail_instance_public_ports" "example" {
  instance_name = aws_lightsail_instance.example.name

  port_info {
    protocol  = "tcp"
    from_port = 80
    to_port   = 80
  }

  port_info {
    protocol  = "tcp"
    from_port = 443
    to_port   = 443
    cidrs     = ["192.168.1.0/24"]
  }
}
```

## Argument Reference

The following arguments are required:

* `instance_name` - (Required) Name of the Lightsail Instance.
* `port_info` - (Required) Configuration block with port information. AWS closes all currently open ports that are not included in the `port_info`. [See below](#port_info).

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).

### port_info

The following arguments are required:

* `from_port` - (Required) First port in a range of open ports on an instance.
* `protocol` - (Required) IP protocol name. Valid values: `tcp`, `all`, `udp`, `icmp`.
* `to_port` - (Required) Last port in a range of open ports on an instance.

The following arguments are optional:

* `cidr_list_aliases` - (Optional) Set of CIDR aliases that define access for a preconfigured range of IP addresses.
* `cidrs` - (Optional) Set of CIDR blocks.
* `ipv6_cidrs` - (Optional) Set of IPv6 CIDR blocks.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - ID of the resource.
