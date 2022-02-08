// Code generated by generators/resource/main.go; DO NOT EDIT.

package logs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_logs_log_group", logGroupResourceType)
}

// logGroupResourceType returns the Terraform awscc_logs_log_group resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Logs::LogGroup resource type.
func logGroupResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The CloudWatch log group ARN.",
			//   "type": "string"
			// }
			Description: "The CloudWatch log group ARN.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"kms_key_id": {
			// Property: KmsKeyId
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.",
			//   "maxLength": 256,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(256),
			},
		},
		"log_group_name": {
			// Property: LogGroupName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the log group. If you don't specify a name, AWS CloudFormation generates a unique ID for the log group.",
			//   "maxLength": 512,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the log group. If you don't specify a name, AWS CloudFormation generates a unique ID for the log group.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 512),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
				tfsdk.RequiresReplace(),
			},
		},
		"retention_in_days": {
			// Property: RetentionInDays
			// CloudFormation resource type schema:
			// {
			//   "description": "The number of days to retain the log events in the specified log group. Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1827, and 3653.",
			//   "enum": [
			//     1,
			//     3,
			//     5,
			//     7,
			//     14,
			//     30,
			//     60,
			//     90,
			//     120,
			//     150,
			//     180,
			//     365,
			//     400,
			//     545,
			//     731,
			//     1827,
			//     3653
			//   ],
			//   "type": "integer"
			// }
			Description: "The number of days to retain the log events in the specified log group. Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1827, and 3653.",
			Type:        types.Int64Type,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.IntInSlice([]int{
					1,
					3,
					5,
					7,
					14,
					30,
					60,
					90,
					120,
					150,
					180,
					365,
					400,
					545,
					731,
					1827,
					3653,
				}),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., :, /, =, +, - and @.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., :, /, =, +, - and @.",
			//         "maxLength": 256,
			//         "minLength": 0,
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
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., :, /, =, +, - and @.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., :, /, =, +, - and @.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
					},
				},
				tfsdk.SetNestedAttributesOptions{},
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
		Description: "Resource schema for AWS::Logs::LogGroup",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Logs::LogGroup").WithTerraformTypeName("awscc_logs_log_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":               "Arn",
		"key":               "Key",
		"kms_key_id":        "KmsKeyId",
		"log_group_name":    "LogGroupName",
		"retention_in_days": "RetentionInDays",
		"tags":              "Tags",
		"value":             "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
