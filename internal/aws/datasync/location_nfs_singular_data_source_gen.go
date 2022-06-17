// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package datasync

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_datasync_location_nfs", locationNFSDataSourceType)
}

// locationNFSDataSourceType returns the Terraform awscc_datasync_location_nfs data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::DataSync::LocationNFS resource type.
func locationNFSDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"location_arn": {
			// Property: LocationArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the NFS location.",
			//   "maxLength": 128,
			//   "pattern": "^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):datasync:[a-z\\-0-9]+:[0-9]{12}:location/loc-[0-9a-z]{17}$",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the NFS location.",
			Type:        types.StringType,
			Computed:    true,
		},
		"location_uri": {
			// Property: LocationUri
			// CloudFormation resource type schema:
			// {
			//   "description": "The URL of the NFS location that was described.",
			//   "maxLength": 4356,
			//   "pattern": "^(efs|nfs|s3|smb|fsxw)://[a-zA-Z0-9./\\-]+$",
			//   "type": "string"
			// }
			Description: "The URL of the NFS location that was described.",
			Type:        types.StringType,
			Computed:    true,
		},
		"mount_options": {
			// Property: MountOptions
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "default": {
			//     "Version": "AUTOMATIC"
			//   },
			//   "description": "The NFS mount options that DataSync can use to mount your NFS share.",
			//   "properties": {
			//     "Version": {
			//       "description": "The specific NFS version that you want DataSync to use to mount your NFS share.",
			//       "enum": [
			//         "AUTOMATIC",
			//         "NFS3",
			//         "NFS4_0",
			//         "NFS4_1"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The NFS mount options that DataSync can use to mount your NFS share.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"version": {
						// Property: Version
						Description: "The specific NFS version that you want DataSync to use to mount your NFS share.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"on_prem_config": {
			// Property: OnPremConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Contains a list of Amazon Resource Names (ARNs) of agents that are used to connect an NFS server.",
			//   "properties": {
			//     "AgentArns": {
			//       "description": "ARN(s) of the agent(s) to use for an NFS location.",
			//       "insertionOrder": false,
			//       "items": {
			//         "maxLength": 128,
			//         "pattern": "^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):datasync:[a-z\\-0-9]+:[0-9]{12}:agent/agent-[0-9a-z]{17}$",
			//         "type": "string"
			//       },
			//       "maxItems": 4,
			//       "minItems": 1,
			//       "type": "array"
			//     }
			//   },
			//   "required": [
			//     "AgentArns"
			//   ],
			//   "type": "object"
			// }
			Description: "Contains a list of Amazon Resource Names (ARNs) of agents that are used to connect an NFS server.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"agent_arns": {
						// Property: AgentArns
						Description: "ARN(s) of the agent(s) to use for an NFS location.",
						Type:        types.ListType{ElemType: types.StringType},
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"server_hostname": {
			// Property: ServerHostname
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the NFS server. This value is the IP address or DNS name of the NFS server.",
			//   "maxLength": 255,
			//   "pattern": "^(([a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)*([A-Za-z0-9\\-]*[A-Za-z0-9])$",
			//   "type": "string"
			// }
			Description: "The name of the NFS server. This value is the IP address or DNS name of the NFS server.",
			Type:        types.StringType,
			Computed:    true,
		},
		"subdirectory": {
			// Property: Subdirectory
			// CloudFormation resource type schema:
			// {
			//   "description": "The subdirectory in the NFS file system that is used to read data from the NFS source location or write data to the NFS destination.",
			//   "maxLength": 4096,
			//   "pattern": "^[a-zA-Z0-9_\\-\\+\\./\\(\\)\\$\\p{Zs}]+$",
			//   "type": "string"
			// }
			Description: "The subdirectory in the NFS file system that is used to read data from the NFS source location or write data to the NFS destination.",
			Type:        types.StringType,
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
			//         "description": "The key for an AWS resource tag.",
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "pattern": "^[a-zA-Z0-9\\s+=._:/-]+$",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for an AWS resource tag.",
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "pattern": "^[a-zA-Z0-9\\s+=._:@/-]+$",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key for an AWS resource tag.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for an AWS resource tag.",
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
		Description: "Data Source schema for AWS::DataSync::LocationNFS",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::DataSync::LocationNFS").WithTerraformTypeName("awscc_datasync_location_nfs")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"agent_arns":      "AgentArns",
		"key":             "Key",
		"location_arn":    "LocationArn",
		"location_uri":    "LocationUri",
		"mount_options":   "MountOptions",
		"on_prem_config":  "OnPremConfig",
		"server_hostname": "ServerHostname",
		"subdirectory":    "Subdirectory",
		"tags":            "Tags",
		"value":           "Value",
		"version":         "Version",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
