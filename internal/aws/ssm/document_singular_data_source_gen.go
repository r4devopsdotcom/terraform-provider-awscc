// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_ssm_document", documentDataSourceType)
}

// documentDataSourceType returns the Terraform awscc_ssm_document data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::SSM::Document resource type.
func documentDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"attachments": {
			// Property: Attachments
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of key and value pairs that describe attachments to a version of a document.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "description": "The key of a key-value pair that identifies the location of an attachment to a document.",
			//         "enum": [
			//           "SourceUrl",
			//           "S3FileUrl",
			//           "AttachmentReference"
			//         ],
			//         "type": "string"
			//       },
			//       "Name": {
			//         "description": "The name of the document attachment file.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$",
			//         "type": "string"
			//       },
			//       "Values": {
			//         "description": "The value of a key-value pair that identifies the location of an attachment to a document. The format for Value depends on the type of key you specify.",
			//         "insertionOrder": false,
			//         "items": {
			//           "maxLength": 100000,
			//           "minLength": 1,
			//           "type": "string"
			//         },
			//         "maxItems": 1,
			//         "minItems": 1,
			//         "type": "array"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "maxItems": 20,
			//   "minItems": 0,
			//   "type": "array"
			// }
			Description: "A list of key and value pairs that describe attachments to a version of a document.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key of a key-value pair that identifies the location of an attachment to a document.",
						Type:        types.StringType,
						Computed:    true,
					},
					"name": {
						// Property: Name
						Description: "The name of the document attachment file.",
						Type:        types.StringType,
						Computed:    true,
					},
					"values": {
						// Property: Values
						Description: "The value of a key-value pair that identifies the location of an attachment to a document. The format for Value depends on the type of key you specify.",
						Type:        types.ListType{ElemType: types.StringType},
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"content": {
			// Property: Content
			// CloudFormation resource type schema:
			// {
			//   "description": "The content for the Systems Manager document in JSON, YAML or String format.",
			//   "type": "string"
			// }
			Description: "The content for the Systems Manager document in JSON, YAML or String format.",
			Type:        types.StringType,
			Computed:    true,
		},
		"document_format": {
			// Property: DocumentFormat
			// CloudFormation resource type schema:
			// {
			//   "default": "JSON",
			//   "description": "Specify the document format for the request. The document format can be either JSON or YAML. JSON is the default format.",
			//   "enum": [
			//     "YAML",
			//     "JSON",
			//     "TEXT"
			//   ],
			//   "type": "string"
			// }
			Description: "Specify the document format for the request. The document format can be either JSON or YAML. JSON is the default format.",
			Type:        types.StringType,
			Computed:    true,
		},
		"document_type": {
			// Property: DocumentType
			// CloudFormation resource type schema:
			// {
			//   "description": "The type of document to create.",
			//   "enum": [
			//     "ApplicationConfiguration",
			//     "ApplicationConfigurationSchema",
			//     "Automation",
			//     "Automation.ChangeTemplate",
			//     "ChangeCalendar",
			//     "CloudFormation",
			//     "Command",
			//     "DeploymentStrategy",
			//     "Package",
			//     "Policy",
			//     "ProblemAnalysis",
			//     "ProblemAnalysisTemplate",
			//     "Session"
			//   ],
			//   "type": "string"
			// }
			Description: "The type of document to create.",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "A name for the Systems Manager document.",
			//   "pattern": "^[a-zA-Z0-9_\\-.]{3,128}$",
			//   "type": "string"
			// }
			Description: "A name for the Systems Manager document.",
			Type:        types.StringType,
			Computed:    true,
		},
		"requires": {
			// Property: Requires
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of SSM documents required by a document. For example, an ApplicationConfiguration document requires an ApplicationConfigurationSchema document.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Name": {
			//         "description": "The name of the required SSM document. The name can be an Amazon Resource Name (ARN).",
			//         "maxLength": 200,
			//         "pattern": "^[a-zA-Z0-9_\\-.:/]{3,200}$",
			//         "type": "string"
			//       },
			//       "Version": {
			//         "description": "The document version required by the current document.",
			//         "maxLength": 8,
			//         "pattern": "([$]LATEST|[$]DEFAULT|^[1-9][0-9]*$)",
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "minItems": 1,
			//   "type": "array"
			// }
			Description: "A list of SSM documents required by a document. For example, an ApplicationConfiguration document requires an ApplicationConfigurationSchema document.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"name": {
						// Property: Name
						Description: "The name of the required SSM document. The name can be an Amazon Resource Name (ARN).",
						Type:        types.StringType,
						Computed:    true,
					},
					"version": {
						// Property: Version
						Description: "The document version required by the current document.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "Optional metadata that you assign to a resource. Tags enable you to categorize a resource in different ways, such as by purpose, owner, or environment.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "description": "The name of the tag.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value of the tag.",
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "pattern": "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$",
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "maxItems": 1000,
			//   "type": "array"
			// }
			Description: "Optional metadata that you assign to a resource. Tags enable you to categorize a resource in different ways, such as by purpose, owner, or environment.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The name of the tag.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value of the tag.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"target_type": {
			// Property: TargetType
			// CloudFormation resource type schema:
			// {
			//   "description": "Specify a target type to define the kinds of resources the document can run on.",
			//   "pattern": "^\\/[\\w\\.\\-\\:\\/]*$",
			//   "type": "string"
			// }
			Description: "Specify a target type to define the kinds of resources the document can run on.",
			Type:        types.StringType,
			Computed:    true,
		},
		"update_method": {
			// Property: UpdateMethod
			// CloudFormation resource type schema:
			// {
			//   "default": "Replace",
			//   "description": "Update method - when set to 'Replace', the update will replace the existing document; when set to 'NewVersion', the update will create a new version.",
			//   "enum": [
			//     "Replace",
			//     "NewVersion"
			//   ],
			//   "type": "string"
			// }
			Description: "Update method - when set to 'Replace', the update will replace the existing document; when set to 'NewVersion', the update will create a new version.",
			Type:        types.StringType,
			Computed:    true,
		},
		"version_name": {
			// Property: VersionName
			// CloudFormation resource type schema:
			// {
			//   "description": "An optional field specifying the version of the artifact you are creating with the document. This value is unique across all versions of a document, and cannot be changed.",
			//   "pattern": "^[a-zA-Z0-9_\\-.]{1,128}$",
			//   "type": "string"
			// }
			Description: "An optional field specifying the version of the artifact you are creating with the document. This value is unique across all versions of a document, and cannot be changed.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::SSM::Document",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SSM::Document").WithTerraformTypeName("awscc_ssm_document")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"attachments":     "Attachments",
		"content":         "Content",
		"document_format": "DocumentFormat",
		"document_type":   "DocumentType",
		"key":             "Key",
		"name":            "Name",
		"requires":        "Requires",
		"tags":            "Tags",
		"target_type":     "TargetType",
		"update_method":   "UpdateMethod",
		"value":           "Value",
		"values":          "Values",
		"version":         "Version",
		"version_name":    "VersionName",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
