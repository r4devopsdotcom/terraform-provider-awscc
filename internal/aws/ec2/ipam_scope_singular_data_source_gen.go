// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_ec2_ipam_scope", iPAMScopeDataSourceType)
}

// iPAMScopeDataSourceType returns the Terraform awscc_ec2_ipam_scope data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::EC2::IPAMScope resource type.
func iPAMScopeDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the IPAM scope.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the IPAM scope.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"ipam_arn": {
			// Property: IpamArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the IPAM this scope is a part of.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the IPAM this scope is a part of.",
			Type:        types.StringType,
			Computed:    true,
		},
		"ipam_id": {
			// Property: IpamId
			// CloudFormation resource type schema:
			// {
			//   "description": "The Id of the IPAM this scope is a part of.",
			//   "type": "string"
			// }
			Description: "The Id of the IPAM this scope is a part of.",
			Type:        types.StringType,
			Computed:    true,
		},
		"ipam_scope_id": {
			// Property: IpamScopeId
			// CloudFormation resource type schema:
			// {
			//   "description": "Id of the IPAM scope.",
			//   "type": "string"
			// }
			Description: "Id of the IPAM scope.",
			Type:        types.StringType,
			Computed:    true,
		},
		"ipam_scope_type": {
			// Property: IpamScopeType
			// CloudFormation resource type schema:
			// {
			//   "description": "Determines whether this scope contains publicly routable space or space for a private network",
			//   "enum": [
			//     "public",
			//     "private"
			//   ],
			//   "type": "string"
			// }
			Description: "Determines whether this scope contains publicly routable space or space for a private network",
			Type:        types.StringType,
			Computed:    true,
		},
		"is_default": {
			// Property: IsDefault
			// CloudFormation resource type schema:
			// {
			//   "description": "Is this one of the default scopes created with the IPAM.",
			//   "type": "boolean"
			// }
			Description: "Is this one of the default scopes created with the IPAM.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"pool_count": {
			// Property: PoolCount
			// CloudFormation resource type schema:
			// {
			//   "description": "The number of pools that currently exist in this scope.",
			//   "type": "integer"
			// }
			Description: "The number of pools that currently exist in this scope.",
			Type:        types.Int64Type,
			Computed:    true,
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
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
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
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::EC2::IPAMScope",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::IPAMScope").WithTerraformTypeName("awscc_ec2_ipam_scope")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":             "Arn",
		"description":     "Description",
		"ipam_arn":        "IpamArn",
		"ipam_id":         "IpamId",
		"ipam_scope_id":   "IpamScopeId",
		"ipam_scope_type": "IpamScopeType",
		"is_default":      "IsDefault",
		"key":             "Key",
		"pool_count":      "PoolCount",
		"tags":            "Tags",
		"value":           "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
