// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package redshift

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	awstypes "github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/enum"
	"github.com/hashicorp/terraform-provider-aws/internal/errs"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	"github.com/hashicorp/terraform-provider-aws/internal/flex"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKResource("aws_redshift_event_subscription", name="Event Subscription")
// @Tags(identifierAttribute="arn")
func resourceEventSubscription() *schema.Resource {
	return &schema.Resource{
		CreateWithoutTimeout: resourceEventSubscriptionCreate,
		ReadWithoutTimeout:   resourceEventSubscriptionRead,
		UpdateWithoutTimeout: resourceEventSubscriptionUpdate,
		DeleteWithoutTimeout: resourceEventSubscriptionDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(40 * time.Minute),
			Delete: schema.DefaultTimeout(40 * time.Minute),
			Update: schema.DefaultTimeout(40 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			names.AttrARN: {
				Type:     schema.TypeString,
				Computed: true,
			},
			"customer_aws_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			names.AttrEnabled: {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"event_categories": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{
						names.AttrConfiguration,
						"management",
						"monitoring",
						"security",
						"pending",
					}, false),
				},
			},
			names.AttrName: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"severity": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "INFO",
				ValidateFunc: validation.StringInSlice([]string{
					"ERROR",
					"INFO",
				}, false),
			},
			names.AttrSNSTopicARN: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: verify.ValidARN,
			},
			"source_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			names.AttrSourceType: {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: enum.Validate[awstypes.SourceType](),
			},
			names.AttrStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},
			names.AttrTags:    tftags.TagsSchema(),
			names.AttrTagsAll: tftags.TagsSchemaComputed(),
		},
	}
}

func resourceEventSubscriptionCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).RedshiftClient(ctx)

	request := &redshift.CreateEventSubscriptionInput{
		SubscriptionName: aws.String(d.Get(names.AttrName).(string)),
		SnsTopicArn:      aws.String(d.Get(names.AttrSNSTopicARN).(string)),
		Enabled:          aws.Bool(d.Get(names.AttrEnabled).(bool)),
		Tags:             getTagsIn(ctx),
	}

	if v, ok := d.GetOk("event_categories"); ok && v.(*schema.Set).Len() > 0 {
		request.EventCategories = flex.ExpandStringValueSet(v.(*schema.Set))
	}

	if v, ok := d.GetOk("source_ids"); ok && v.(*schema.Set).Len() > 0 {
		request.SourceIds = flex.ExpandStringValueSet(v.(*schema.Set))
	}

	if v, ok := d.GetOk("severity"); ok {
		request.Severity = aws.String(v.(string))
	}

	if v, ok := d.GetOk(names.AttrSourceType); ok {
		request.SourceType = aws.String(v.(string))
	}

	log.Println("[DEBUG] Create Redshift Event Subscription:", request)

	output, err := conn.CreateEventSubscription(ctx, request)
	if err != nil || output.EventSubscription == nil {
		return sdkdiag.AppendErrorf(diags, "creating Redshift Event Subscription %s: %s", d.Get(names.AttrName).(string), err)
	}

	d.SetId(aws.ToString(output.EventSubscription.CustSubscriptionId))

	return append(diags, resourceEventSubscriptionRead(ctx, d, meta)...)
}

func resourceEventSubscriptionRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).RedshiftClient(ctx)

	sub, err := findEventSubscriptionByName(ctx, conn, d.Id())

	if !d.IsNewResource() && tfresource.NotFound(err) {
		log.Printf("[WARN] Redshift Event Subscription (%s) not found, removing from state", d.Id())
		d.SetId("")
		return diags
	}
	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading Redshift Event Subscription (%s): %s", d.Id(), err)
	}

	arn := arn.ARN{
		Partition: meta.(*conns.AWSClient).Partition(ctx),
		Service:   names.Redshift,
		Region:    meta.(*conns.AWSClient).Region(ctx),
		AccountID: meta.(*conns.AWSClient).AccountID(ctx),
		Resource:  fmt.Sprintf("eventsubscription:%s", d.Id()),
	}.String()
	d.Set(names.AttrARN, arn)
	d.Set("customer_aws_id", sub.CustomerAwsId)
	d.Set(names.AttrEnabled, sub.Enabled)
	d.Set("event_categories", sub.EventCategoriesList)
	d.Set(names.AttrName, sub.CustSubscriptionId)
	d.Set("severity", sub.Severity)
	d.Set(names.AttrSNSTopicARN, sub.SnsTopicArn)
	d.Set("source_ids", sub.SourceIdsList)
	d.Set(names.AttrSourceType, sub.SourceType)
	d.Set(names.AttrStatus, sub.Status)

	setTagsOut(ctx, sub.Tags)

	return diags
}

func resourceEventSubscriptionUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).RedshiftClient(ctx)

	if d.HasChangesExcept(names.AttrTags, names.AttrTagsAll) {
		req := &redshift.ModifyEventSubscriptionInput{
			SubscriptionName: aws.String(d.Id()),
			SnsTopicArn:      aws.String(d.Get(names.AttrSNSTopicARN).(string)),
			Enabled:          aws.Bool(d.Get(names.AttrEnabled).(bool)),
			SourceIds:        flex.ExpandStringValueSet(d.Get("source_ids").(*schema.Set)),
			SourceType:       aws.String(d.Get(names.AttrSourceType).(string)),
			Severity:         aws.String(d.Get("severity").(string)),
			EventCategories:  flex.ExpandStringValueSet(d.Get("event_categories").(*schema.Set)),
		}

		log.Printf("[DEBUG] Redshift Event Subscription modification request: %#v", req)
		_, err := conn.ModifyEventSubscription(ctx, req)
		if err != nil {
			return sdkdiag.AppendErrorf(diags, "Modifying Redshift Event Subscription %s failed: %s", d.Id(), err)
		}
	}

	return append(diags, resourceEventSubscriptionRead(ctx, d, meta)...)
}

func resourceEventSubscriptionDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).RedshiftClient(ctx)
	deleteOpts := redshift.DeleteEventSubscriptionInput{
		SubscriptionName: aws.String(d.Id()),
	}

	if _, err := conn.DeleteEventSubscription(ctx, &deleteOpts); err != nil {
		if errs.IsA[*awstypes.SubscriptionNotFoundFault](err) {
			return diags
		}
		return sdkdiag.AppendErrorf(diags, "deleting Redshift Event Subscription %s: %s", d.Id(), err)
	}

	return diags
}
