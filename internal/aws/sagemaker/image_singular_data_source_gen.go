// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_sagemaker_image", imageDataSourceType)
}

// imageDataSourceType returns the Terraform awscc_sagemaker_image data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::SageMaker::Image resource type.
func imageDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"image_arn": {
			// Property: ImageArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the image.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "^arn:aws(-[\\w]+)*:sagemaker:[a-z0-9\\-]*:[0-9]{12}:image\\/[a-z0-9]([-.]?[a-z0-9])*$",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the image.",
			Type:        types.StringType,
			Computed:    true,
		},
		"image_description": {
			// Property: ImageDescription
			// CloudFormation resource type schema:
			// {
			//   "description": "A description of the image.",
			//   "maxLength": 512,
			//   "minLength": 1,
			//   "pattern": ".+",
			//   "type": "string"
			// }
			Description: "A description of the image.",
			Type:        types.StringType,
			Computed:    true,
		},
		"image_display_name": {
			// Property: ImageDisplayName
			// CloudFormation resource type schema:
			// {
			//   "description": "The display name of the image.",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "^[A-Za-z0-9 -_]+$",
			//   "type": "string"
			// }
			Description: "The display name of the image.",
			Type:        types.StringType,
			Computed:    true,
		},
		"image_name": {
			// Property: ImageName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the image.",
			//   "maxLength": 63,
			//   "minLength": 1,
			//   "pattern": "^[a-zA-Z0-9]([-.]?[a-zA-Z0-9])*$",
			//   "type": "string"
			// }
			Description: "The name of the image.",
			Type:        types.StringType,
			Computed:    true,
		},
		"image_role_arn": {
			// Property: ImageRoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker to perform tasks on behalf of the customer.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "^arn:aws(-[\\w]+)*:iam::[0-9]{12}:role/.*$",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker to perform tasks on behalf of the customer.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 256,
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
			//   "minItems": 1,
			//   "type": "array"
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
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
		Description: "Data Source schema for AWS::SageMaker::Image",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SageMaker::Image").WithTerraformTypeName("awscc_sagemaker_image")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"image_arn":          "ImageArn",
		"image_description":  "ImageDescription",
		"image_display_name": "ImageDisplayName",
		"image_name":         "ImageName",
		"image_role_arn":     "ImageRoleArn",
		"key":                "Key",
		"tags":               "Tags",
		"value":              "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
