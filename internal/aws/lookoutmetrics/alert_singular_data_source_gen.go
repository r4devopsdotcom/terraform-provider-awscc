// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package lookoutmetrics

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_lookoutmetrics_alert", alertDataSourceType)
}

// alertDataSourceType returns the Terraform awscc_lookoutmetrics_alert data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::LookoutMetrics::Alert resource type.
func alertDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"action": {
			// Property: Action
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The action to be taken by the alert when an anomaly is detected.",
			//   "properties": {
			//     "LambdaConfiguration": {
			//       "additionalProperties": false,
			//       "description": "Configuration options for a Lambda alert action.",
			//       "properties": {
			//         "LambdaArn": {
			//           "description": "ARN of a Lambda to send alert notifications to.",
			//           "maxLength": 256,
			//           "pattern": "",
			//           "type": "string"
			//         },
			//         "RoleArn": {
			//           "description": "ARN of an IAM role that LookoutMetrics should assume to access the Lambda function.",
			//           "maxLength": 256,
			//           "pattern": "",
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "RoleArn",
			//         "LambdaArn"
			//       ],
			//       "type": "object"
			//     },
			//     "SNSConfiguration": {
			//       "additionalProperties": false,
			//       "description": "Configuration options for an SNS alert action.",
			//       "properties": {
			//         "RoleArn": {
			//           "description": "ARN of an IAM role that LookoutMetrics should assume to access the SNS topic.",
			//           "maxLength": 256,
			//           "pattern": "",
			//           "type": "string"
			//         },
			//         "SnsTopicArn": {
			//           "description": "ARN of an SNS topic to send alert notifications to.",
			//           "maxLength": 256,
			//           "pattern": "",
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "RoleArn",
			//         "SnsTopicArn"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The action to be taken by the alert when an anomaly is detected.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"lambda_configuration": {
						// Property: LambdaConfiguration
						Description: "Configuration options for a Lambda alert action.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"lambda_arn": {
									// Property: LambdaArn
									Description: "ARN of a Lambda to send alert notifications to.",
									Type:        types.StringType,
									Computed:    true,
								},
								"role_arn": {
									// Property: RoleArn
									Description: "ARN of an IAM role that LookoutMetrics should assume to access the Lambda function.",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"sns_configuration": {
						// Property: SNSConfiguration
						Description: "Configuration options for an SNS alert action.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"role_arn": {
									// Property: RoleArn
									Description: "ARN of an IAM role that LookoutMetrics should assume to access the SNS topic.",
									Type:        types.StringType,
									Computed:    true,
								},
								"sns_topic_arn": {
									// Property: SnsTopicArn
									Description: "ARN of an SNS topic to send alert notifications to.",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"alert_description": {
			// Property: AlertDescription
			// CloudFormation resource type schema:
			// {
			//   "description": "A description for the alert.",
			//   "maxLength": 256,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "A description for the alert.",
			Type:        types.StringType,
			Computed:    true,
		},
		"alert_name": {
			// Property: AlertName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the alert. If not provided, a name is generated automatically.",
			//   "maxLength": 63,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the alert. If not provided, a name is generated automatically.",
			Type:        types.StringType,
			Computed:    true,
		},
		"alert_sensitivity_threshold": {
			// Property: AlertSensitivityThreshold
			// CloudFormation resource type schema:
			// {
			//   "description": "A number between 0 and 100 (inclusive) that tunes the sensitivity of the alert.",
			//   "maximum": 100,
			//   "minimum": 0,
			//   "type": "integer"
			// }
			Description: "A number between 0 and 100 (inclusive) that tunes the sensitivity of the alert.",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"anomaly_detector_arn": {
			// Property: AnomalyDetectorArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon resource name (ARN) of the Anomaly Detector to alert.",
			//   "maxLength": 256,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon resource name (ARN) of the Anomaly Detector to alert.",
			Type:        types.StringType,
			Computed:    true,
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "ARN assigned to the alert.",
			//   "maxLength": 256,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "ARN assigned to the alert.",
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
		Description: "Data Source schema for AWS::LookoutMetrics::Alert",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::LookoutMetrics::Alert").WithTerraformTypeName("awscc_lookoutmetrics_alert")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"action":                      "Action",
		"alert_description":           "AlertDescription",
		"alert_name":                  "AlertName",
		"alert_sensitivity_threshold": "AlertSensitivityThreshold",
		"anomaly_detector_arn":        "AnomalyDetectorArn",
		"arn":                         "Arn",
		"lambda_arn":                  "LambdaArn",
		"lambda_configuration":        "LambdaConfiguration",
		"role_arn":                    "RoleArn",
		"sns_configuration":           "SNSConfiguration",
		"sns_topic_arn":               "SnsTopicArn",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
