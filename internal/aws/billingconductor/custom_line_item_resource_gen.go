// Code generated by generators/resource/main.go; DO NOT EDIT.

package billingconductor

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_billingconductor_custom_line_item", customLineItemResourceType)
}

// customLineItemResourceType returns the Terraform awscc_billingconductor_custom_line_item resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::BillingConductor::CustomLineItem resource type.
func customLineItemResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "ARN",
			//   "pattern": "(arn:aws(-cn)?:billingconductor::[0-9]{12}:customlineitem/)?[a-zA-Z0-9]{10}",
			//   "type": "string"
			// }
			Description: "ARN",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"association_size": {
			// Property: AssociationSize
			// CloudFormation resource type schema:
			// {
			//   "description": "Number of source values associated to this custom line item",
			//   "type": "integer"
			// }
			Description: "Number of source values associated to this custom line item",
			Type:        types.Int64Type,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"billing_group_arn": {
			// Property: BillingGroupArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Billing Group ARN",
			//   "pattern": "arn:aws(-cn)?:billingconductor::[0-9]{12}:billinggroup/?[0-9]{12}",
			//   "type": "string"
			// }
			Description: "Billing Group ARN",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringMatch(regexp.MustCompile("arn:aws(-cn)?:billingconductor::[0-9]{12}:billinggroup/?[0-9]{12}"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"billing_period_range": {
			// Property: BillingPeriodRange
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "ExclusiveEndBillingPeriod": {
			//       "pattern": "\\d{4}-(0?[1-9]|1[012])",
			//       "type": "string"
			//     },
			//     "InclusiveStartBillingPeriod": {
			//       "pattern": "\\d{4}-(0?[1-9]|1[012])",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"exclusive_end_billing_period": {
						// Property: ExclusiveEndBillingPeriod
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringMatch(regexp.MustCompile("\\d{4}-(0?[1-9]|1[012])"), ""),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							tfsdk.UseStateForUnknown(),
							tfsdk.RequiresReplace(),
						},
					},
					"inclusive_start_billing_period": {
						// Property: InclusiveStartBillingPeriod
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringMatch(regexp.MustCompile("\\d{4}-(0?[1-9]|1[012])"), ""),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							tfsdk.UseStateForUnknown(),
							tfsdk.RequiresReplace(),
						},
					},
				},
			),
			Optional: true,
		},
		"creation_time": {
			// Property: CreationTime
			// CloudFormation resource type schema:
			// {
			//   "description": "Creation timestamp in UNIX epoch time format",
			//   "type": "integer"
			// }
			Description: "Creation timestamp in UNIX epoch time format",
			Type:        types.Int64Type,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"currency_code": {
			// Property: CurrencyCode
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "USD",
			//     "CNY"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"custom_line_item_charge_details": {
			// Property: CustomLineItemChargeDetails
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "Flat": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "ChargeValue": {
			//           "maximum": 1000000,
			//           "minimum": 0,
			//           "type": "number"
			//         }
			//       },
			//       "required": [
			//         "ChargeValue"
			//       ],
			//       "type": "object"
			//     },
			//     "Percentage": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "ChildAssociatedResources": {
			//           "insertionOrder": false,
			//           "items": {
			//             "pattern": "(arn:aws(-cn)?:billingconductor::[0-9]{12}:(customlineitem|billinggroup)/)?[a-zA-Z0-9]{10,12}",
			//             "type": "string"
			//           },
			//           "type": "array",
			//           "uniqueItems": true
			//         },
			//         "PercentageValue": {
			//           "maximum": 10000,
			//           "minimum": 0,
			//           "type": "number"
			//         }
			//       },
			//       "required": [
			//         "PercentageValue"
			//       ],
			//       "type": "object"
			//     },
			//     "Type": {
			//       "enum": [
			//         "FEE",
			//         "CREDIT"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "Type"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"flat": {
						// Property: Flat
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"charge_value": {
									// Property: ChargeValue
									Type:     types.Float64Type,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.FloatBetween(0.000000, 1000000.000000),
									},
								},
							},
						),
						Optional: true,
					},
					"percentage": {
						// Property: Percentage
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"child_associated_resources": {
									// Property: ChildAssociatedResources
									Type:     types.SetType{ElemType: types.StringType},
									Optional: true,
									Validators: []tfsdk.AttributeValidator{
										validate.ArrayForEach(validate.StringMatch(regexp.MustCompile("(arn:aws(-cn)?:billingconductor::[0-9]{12}:(customlineitem|billinggroup)/)?[a-zA-Z0-9]{10,12}"), "")),
									},
								},
								"percentage_value": {
									// Property: PercentageValue
									Type:     types.Float64Type,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.FloatBetween(0.000000, 10000.000000),
									},
								},
							},
						),
						Optional: true,
					},
					"type": {
						// Property: Type
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"FEE",
								"CREDIT",
							}),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							tfsdk.RequiresReplace(),
						},
					},
				},
			),
			Optional: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 255,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(255),
			},
		},
		"last_modified_time": {
			// Property: LastModifiedTime
			// CloudFormation resource type schema:
			// {
			//   "description": "Latest modified timestamp in UNIX epoch time format",
			//   "type": "integer"
			// }
			Description: "Latest modified timestamp in UNIX epoch time format",
			Type:        types.Int64Type,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "[a-zA-Z0-9_\\+=\\.\\-@]+",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 128),
				validate.StringMatch(regexp.MustCompile("[a-zA-Z0-9_\\+=\\.\\-@]+"), ""),
			},
		},
		"product_code": {
			// Property: ProductCode
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 29,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 256),
						},
					},
				},
			),
			Optional: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			tfsdk.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "A custom line item is an one time charge that is applied to a specific billing group's bill.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::BillingConductor::CustomLineItem").WithTerraformTypeName("awscc_billingconductor_custom_line_item")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                             "Arn",
		"association_size":                "AssociationSize",
		"billing_group_arn":               "BillingGroupArn",
		"billing_period_range":            "BillingPeriodRange",
		"charge_value":                    "ChargeValue",
		"child_associated_resources":      "ChildAssociatedResources",
		"creation_time":                   "CreationTime",
		"currency_code":                   "CurrencyCode",
		"custom_line_item_charge_details": "CustomLineItemChargeDetails",
		"description":                     "Description",
		"exclusive_end_billing_period":    "ExclusiveEndBillingPeriod",
		"flat":                            "Flat",
		"inclusive_start_billing_period":  "InclusiveStartBillingPeriod",
		"key":                             "Key",
		"last_modified_time":              "LastModifiedTime",
		"name":                            "Name",
		"percentage":                      "Percentage",
		"percentage_value":                "PercentageValue",
		"product_code":                    "ProductCode",
		"tags":                            "Tags",
		"type":                            "Type",
		"value":                           "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
