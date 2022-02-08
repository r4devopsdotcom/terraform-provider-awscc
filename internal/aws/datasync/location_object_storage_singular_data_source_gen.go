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
	registry.AddDataSourceTypeFactory("awscc_datasync_location_object_storage", locationObjectStorageDataSourceType)
}

// locationObjectStorageDataSourceType returns the Terraform awscc_datasync_location_object_storage data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::DataSync::LocationObjectStorage resource type.
func locationObjectStorageDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"access_key": {
			// Property: AccessKey
			// CloudFormation resource type schema:
			// {
			//   "description": "Optional. The access key is used if credentials are required to access the self-managed object storage server.",
			//   "maxLength": 200,
			//   "minLength": 8,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Optional. The access key is used if credentials are required to access the self-managed object storage server.",
			Type:        types.StringType,
			Computed:    true,
		},
		"agent_arns": {
			// Property: AgentArns
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the agents associated with the self-managed object storage server location.",
			//   "insertionOrder": false,
			//   "items": {
			//     "maxLength": 128,
			//     "pattern": "",
			//     "type": "string"
			//   },
			//   "maxItems": 4,
			//   "minItems": 1,
			//   "type": "array"
			// }
			Description: "The Amazon Resource Name (ARN) of the agents associated with the self-managed object storage server location.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"bucket_name": {
			// Property: BucketName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the bucket on the self-managed object storage server.",
			//   "maxLength": 63,
			//   "minLength": 3,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the bucket on the self-managed object storage server.",
			Type:        types.StringType,
			Computed:    true,
		},
		"location_arn": {
			// Property: LocationArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the location that is created.",
			//   "maxLength": 128,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the location that is created.",
			Type:        types.StringType,
			Computed:    true,
		},
		"location_uri": {
			// Property: LocationUri
			// CloudFormation resource type schema:
			// {
			//   "description": "The URL of the object storage location that was described.",
			//   "maxLength": 4356,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The URL of the object storage location that was described.",
			Type:        types.StringType,
			Computed:    true,
		},
		"secret_key": {
			// Property: SecretKey
			// CloudFormation resource type schema:
			// {
			//   "description": "Optional. The secret key is used if credentials are required to access the self-managed object storage server.",
			//   "maxLength": 200,
			//   "minLength": 8,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Optional. The secret key is used if credentials are required to access the self-managed object storage server.",
			Type:        types.StringType,
			Computed:    true,
		},
		"server_hostname": {
			// Property: ServerHostname
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the self-managed object storage server. This value is the IP address or Domain Name Service (DNS) name of the object storage server.",
			//   "maxLength": 255,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the self-managed object storage server. This value is the IP address or Domain Name Service (DNS) name of the object storage server.",
			Type:        types.StringType,
			Computed:    true,
		},
		"server_port": {
			// Property: ServerPort
			// CloudFormation resource type schema:
			// {
			//   "description": "The port that your self-managed server accepts inbound network traffic on.",
			//   "maximum": 65536,
			//   "minimum": 1,
			//   "type": "integer"
			// }
			Description: "The port that your self-managed server accepts inbound network traffic on.",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"server_protocol": {
			// Property: ServerProtocol
			// CloudFormation resource type schema:
			// {
			//   "description": "The protocol that the object storage server uses to communicate.",
			//   "enum": [
			//     "HTTPS",
			//     "HTTP"
			//   ],
			//   "type": "string"
			// }
			Description: "The protocol that the object storage server uses to communicate.",
			Type:        types.StringType,
			Computed:    true,
		},
		"subdirectory": {
			// Property: Subdirectory
			// CloudFormation resource type schema:
			// {
			//   "description": "The subdirectory in the self-managed object storage server that is used to read data from.",
			//   "maxLength": 4096,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The subdirectory in the self-managed object storage server that is used to read data from.",
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
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for an AWS resource tag.",
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "pattern": "",
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
				tfsdk.SetNestedAttributesOptions{},
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
		Description: "Data Source schema for AWS::DataSync::LocationObjectStorage",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::DataSync::LocationObjectStorage").WithTerraformTypeName("awscc_datasync_location_object_storage")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"access_key":      "AccessKey",
		"agent_arns":      "AgentArns",
		"bucket_name":     "BucketName",
		"key":             "Key",
		"location_arn":    "LocationArn",
		"location_uri":    "LocationUri",
		"secret_key":      "SecretKey",
		"server_hostname": "ServerHostname",
		"server_port":     "ServerPort",
		"server_protocol": "ServerProtocol",
		"subdirectory":    "Subdirectory",
		"tags":            "Tags",
		"value":           "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
