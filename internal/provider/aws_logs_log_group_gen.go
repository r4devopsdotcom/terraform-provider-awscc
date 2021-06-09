// Code generated by generators/resource/main.go; DO NOT EDIT.

package provider

import (
	"context"

	tfsdk "github.com/hashicorp/terraform-plugin-framework"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func init() {
	RegisterResourceType("aws_logs_log_group", resourceAwsLogsLogGroup)
}

// resourceAwsLogsLogGroup returns the Terraform aws_logs_log_group resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Logs::LogGroup resource type.
func resourceAwsLogsLogGroup(ctx context.Context) (tfsdk.ResourceType, error) {
	// Subproperty definitions.

	// Root property definitions.

	// LogGroupName
	/*
	   {
	     "description": "The name of the log group. If you don't specify a name, AWS CloudFormation generates a unique ID for the log group.",
	     "maxLength": 512,
	     "minLength": 1,
	     "pattern": "^[.\\-_/#A-Za-z0-9]{1,512}$",
	     "type": "string"
	   }
	*/
	propLogGroupNameAttribute := schema.Attribute{}
	propLogGroupNameAttribute.Type = types.StringType
	propLogGroupNameAttribute.Optional = true
	propLogGroupNameAttribute.Description = `The name of the log group. If you don't specify a name, AWS CloudFormation generates a unique ID for the log group.`

	// KmsKeyId
	/*
	   {
	     "description": "The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.",
	     "maxLength": 256,
	     "pattern": "^arn:[a-z0-9-]+:kms:[a-z0-9-]+:\\d{12}:(key|alias)/.+$",
	     "type": "string"
	   }
	*/
	propKmsKeyIdAttribute := schema.Attribute{}
	propKmsKeyIdAttribute.Type = types.StringType
	propKmsKeyIdAttribute.Optional = true
	propKmsKeyIdAttribute.Description = `The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.`

	// RetentionInDays
	/*
	   {
	     "description": "The number of days to retain the log events in the specified log group. Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1827, and 3653.",
	     "enum": [
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
	       3653
	     ],
	     "type": "integer"
	   }
	*/
	propRetentionInDaysAttribute := schema.Attribute{}
	propRetentionInDaysAttribute.Type = types.NumberType
	propRetentionInDaysAttribute.Optional = true
	propRetentionInDaysAttribute.Description = `The number of days to retain the log events in the specified log group. Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1827, and 3653.`

	// Arn
	/*
	   {
	     "description": "The CloudWatch log group ARN.",
	     "type": "string"
	   }
	*/
	propArnAttribute := schema.Attribute{}
	propArnAttribute.Type = types.StringType
	propArnAttribute.Computed = true
	propArnAttribute.Description = `The CloudWatch log group ARN.`

	attributes := make(map[string]schema.Attribute)
	attributes["arn"] = propArnAttribute
	attributes["kms_key_id"] = propKmsKeyIdAttribute
	attributes["log_group_name"] = propLogGroupNameAttribute
	attributes["retention_in_days"] = propRetentionInDaysAttribute

	schema := schema.Schema{
		Version:    1,
		Attributes: attributes,
	}

	resourceType := NewGenericResourceType(
		"AWS::Logs::LogGroup",
		"aws_logs_log_group",
		schema,
	)

	return resourceType, nil
}
