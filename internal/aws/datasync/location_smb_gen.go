// Code generated by generators/resource/main.go; DO NOT EDIT.

package datasync

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	providertypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddResourceTypeFactory("awscc_datasync_location_smb", locationSMBResourceType)
}

// locationSMBResourceType returns the Terraform awscc_datasync_location_smb resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::DataSync::LocationSMB resource type.
func locationSMBResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"agent_arns": {
			// Property: AgentArns
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Names (ARNs) of agents to use for a Simple Message Block (SMB) location.",
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
			Description: "The Amazon Resource Names (ARNs) of agents to use for a Simple Message Block (SMB) location.",
			Type:        types.ListType{ElemType: types.StringType},
			Required:    true,
		},
		"domain": {
			// Property: Domain
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the Windows domain that the SMB server belongs to.",
			//   "maxLength": 253,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the Windows domain that the SMB server belongs to.",
			Type:        types.StringType,
			Optional:    true,
		},
		"location_arn": {
			// Property: LocationArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the SMB location that is created.",
			//   "maxLength": 128,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the SMB location that is created.",
			Type:        types.StringType,
			Computed:    true,
		},
		"location_uri": {
			// Property: LocationUri
			// CloudFormation resource type schema:
			// {
			//   "description": "The URL of the SMB location that was described.",
			//   "maxLength": 4356,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The URL of the SMB location that was described.",
			Type:        types.StringType,
			Computed:    true,
		},
		"mount_options": {
			// Property: MountOptions
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The mount options used by DataSync to access the SMB server.",
			//   "properties": {
			//     "Version": {
			//       "description": "The specific SMB version that you want DataSync to use to mount your SMB share.",
			//       "enum": [
			//         "AUTOMATIC",
			//         "SMB2",
			//         "SMB3"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The mount options used by DataSync to access the SMB server.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"version": {
						// Property: Version
						Description: "The specific SMB version that you want DataSync to use to mount your SMB share.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Optional: true,
		},
		"password": {
			// Property: Password
			// CloudFormation resource type schema:
			// {
			//   "description": "The password of the user who can mount the share and has the permissions to access files and folders in the SMB share.",
			//   "maxLength": 104,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The password of the user who can mount the share and has the permissions to access files and folders in the SMB share.",
			Type:        types.StringType,
			Required:    true,
			// Password is a write-only attribute.
		},
		"server_hostname": {
			// Property: ServerHostname
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the SMB server. This value is the IP address or Domain Name Service (DNS) name of the SMB server.",
			//   "maxLength": 255,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the SMB server. This value is the IP address or Domain Name Service (DNS) name of the SMB server.",
			Type:        types.StringType,
			Required:    true,
			// ServerHostname is a force-new attribute.
			// ServerHostname is a write-only attribute.
		},
		"subdirectory": {
			// Property: Subdirectory
			// CloudFormation resource type schema:
			// {
			//   "description": "The subdirectory in the SMB file system that is used to read data from the SMB source location or write data to the SMB destination",
			//   "maxLength": 4096,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The subdirectory in the SMB file system that is used to read data from the SMB source location or write data to the SMB destination",
			Type:        types.StringType,
			Required:    true,
			// Subdirectory is a write-only attribute.
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
			Attributes: providertypes.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key for an AWS resource tag.",
						Type:        types.StringType,
						Required:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for an AWS resource tag.",
						Type:        types.StringType,
						Required:    true,
					},
				},
				providertypes.SetNestedAttributesOptions{
					MaxItems: 50,
				},
			),
			Optional: true,
		},
		"user": {
			// Property: User
			// CloudFormation resource type schema:
			// {
			//   "description": "The user who can mount the share, has the permissions to access files and folders in the SMB share.",
			//   "maxLength": 104,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The user who can mount the share, has the permissions to access files and folders in the SMB share.",
			Type:        types.StringType,
			Required:    true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource schema for AWS::DataSync::LocationSMB.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::DataSync::LocationSMB").WithTerraformTypeName("awscc_datasync_location_smb").WithTerraformSchema(schema)

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Password",
		"/properties/Subdirectory",
		"/properties/ServerHostname",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_datasync_location_smb", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
