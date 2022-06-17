// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package frauddetector

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_frauddetector_event_type", eventTypeDataSourceType)
}

// eventTypeDataSourceType returns the Terraform awscc_frauddetector_event_type data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::FraudDetector::EventType resource type.
func eventTypeDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the event type.",
			//   "type": "string"
			// }
			Description: "The ARN of the event type.",
			Type:        types.StringType,
			Computed:    true,
		},
		"created_time": {
			// Property: CreatedTime
			// CloudFormation resource type schema:
			// {
			//   "description": "The time when the event type was created.",
			//   "type": "string"
			// }
			Description: "The time when the event type was created.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "The description of the event type.",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The description of the event type.",
			Type:        types.StringType,
			Computed:    true,
		},
		"entity_types": {
			// Property: EntityTypes
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Arn": {
			//         "type": "string"
			//       },
			//       "CreatedTime": {
			//         "description": "The time when the event type was created.",
			//         "type": "string"
			//       },
			//       "Description": {
			//         "description": "The description.",
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Inline": {
			//         "type": "boolean"
			//       },
			//       "LastUpdatedTime": {
			//         "description": "The time when the event type was last updated.",
			//         "type": "string"
			//       },
			//       "Name": {
			//         "type": "string"
			//       },
			//       "Tags": {
			//         "description": "Tags associated with this event type.",
			//         "insertionOrder": false,
			//         "items": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "Key": {
			//               "maxLength": 128,
			//               "minLength": 1,
			//               "type": "string"
			//             },
			//             "Value": {
			//               "maxLength": 256,
			//               "minLength": 0,
			//               "type": "string"
			//             }
			//           },
			//           "required": [
			//             "Key",
			//             "Value"
			//           ],
			//           "type": "object"
			//         },
			//         "maxItems": 200,
			//         "type": "array",
			//         "uniqueItems": false
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "minItems": 1,
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"arn": {
						// Property: Arn
						Type:     types.StringType,
						Computed: true,
					},
					"created_time": {
						// Property: CreatedTime
						Description: "The time when the event type was created.",
						Type:        types.StringType,
						Computed:    true,
					},
					"description": {
						// Property: Description
						Description: "The description.",
						Type:        types.StringType,
						Computed:    true,
					},
					"inline": {
						// Property: Inline
						Type:     types.BoolType,
						Computed: true,
					},
					"last_updated_time": {
						// Property: LastUpdatedTime
						Description: "The time when the event type was last updated.",
						Type:        types.StringType,
						Computed:    true,
					},
					"name": {
						// Property: Name
						Type:     types.StringType,
						Computed: true,
					},
					"tags": {
						// Property: Tags
						Description: "Tags associated with this event type.",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"key": {
									// Property: Key
									Type:     types.StringType,
									Computed: true,
								},
								"value": {
									// Property: Value
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"event_variables": {
			// Property: EventVariables
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Arn": {
			//         "type": "string"
			//       },
			//       "CreatedTime": {
			//         "description": "The time when the event type was created.",
			//         "type": "string"
			//       },
			//       "DataSource": {
			//         "enum": [
			//           "EVENT"
			//         ],
			//         "type": "string"
			//       },
			//       "DataType": {
			//         "enum": [
			//           "STRING",
			//           "INTEGER",
			//           "FLOAT",
			//           "BOOLEAN"
			//         ],
			//         "type": "string"
			//       },
			//       "DefaultValue": {
			//         "type": "string"
			//       },
			//       "Description": {
			//         "description": "The description.",
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Inline": {
			//         "type": "boolean"
			//       },
			//       "LastUpdatedTime": {
			//         "description": "The time when the event type was last updated.",
			//         "type": "string"
			//       },
			//       "Name": {
			//         "type": "string"
			//       },
			//       "Tags": {
			//         "description": "Tags associated with this event type.",
			//         "insertionOrder": false,
			//         "items": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "Key": {
			//               "maxLength": 128,
			//               "minLength": 1,
			//               "type": "string"
			//             },
			//             "Value": {
			//               "maxLength": 256,
			//               "minLength": 0,
			//               "type": "string"
			//             }
			//           },
			//           "required": [
			//             "Key",
			//             "Value"
			//           ],
			//           "type": "object"
			//         },
			//         "maxItems": 200,
			//         "type": "array",
			//         "uniqueItems": false
			//       },
			//       "VariableType": {
			//         "enum": [
			//           "AUTH_CODE",
			//           "AVS",
			//           "BILLING_ADDRESS_L1",
			//           "BILLING_ADDRESS_L2",
			//           "BILLING_CITY",
			//           "BILLING_COUNTRY",
			//           "BILLING_NAME",
			//           "BILLING_PHONE",
			//           "BILLING_STATE",
			//           "BILLING_ZIP",
			//           "CARD_BIN",
			//           "CATEGORICAL",
			//           "CURRENCY_CODE",
			//           "EMAIL_ADDRESS",
			//           "FINGERPRINT",
			//           "FRAUD_LABEL",
			//           "FREE_FORM_TEXT",
			//           "IP_ADDRESS",
			//           "NUMERIC",
			//           "ORDER_ID",
			//           "PAYMENT_TYPE",
			//           "PHONE_NUMBER",
			//           "PRICE",
			//           "PRODUCT_CATEGORY",
			//           "SHIPPING_ADDRESS_L1",
			//           "SHIPPING_ADDRESS_L2",
			//           "SHIPPING_CITY",
			//           "SHIPPING_COUNTRY",
			//           "SHIPPING_NAME",
			//           "SHIPPING_PHONE",
			//           "SHIPPING_STATE",
			//           "SHIPPING_ZIP",
			//           "USERAGENT"
			//         ],
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "minItems": 1,
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"arn": {
						// Property: Arn
						Type:     types.StringType,
						Computed: true,
					},
					"created_time": {
						// Property: CreatedTime
						Description: "The time when the event type was created.",
						Type:        types.StringType,
						Computed:    true,
					},
					"data_source": {
						// Property: DataSource
						Type:     types.StringType,
						Computed: true,
					},
					"data_type": {
						// Property: DataType
						Type:     types.StringType,
						Computed: true,
					},
					"default_value": {
						// Property: DefaultValue
						Type:     types.StringType,
						Computed: true,
					},
					"description": {
						// Property: Description
						Description: "The description.",
						Type:        types.StringType,
						Computed:    true,
					},
					"inline": {
						// Property: Inline
						Type:     types.BoolType,
						Computed: true,
					},
					"last_updated_time": {
						// Property: LastUpdatedTime
						Description: "The time when the event type was last updated.",
						Type:        types.StringType,
						Computed:    true,
					},
					"name": {
						// Property: Name
						Type:     types.StringType,
						Computed: true,
					},
					"tags": {
						// Property: Tags
						Description: "Tags associated with this event type.",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"key": {
									// Property: Key
									Type:     types.StringType,
									Computed: true,
								},
								"value": {
									// Property: Value
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"variable_type": {
						// Property: VariableType
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"labels": {
			// Property: Labels
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Arn": {
			//         "type": "string"
			//       },
			//       "CreatedTime": {
			//         "description": "The time when the event type was created.",
			//         "type": "string"
			//       },
			//       "Description": {
			//         "description": "The description.",
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Inline": {
			//         "type": "boolean"
			//       },
			//       "LastUpdatedTime": {
			//         "description": "The time when the event type was last updated.",
			//         "type": "string"
			//       },
			//       "Name": {
			//         "type": "string"
			//       },
			//       "Tags": {
			//         "description": "Tags associated with this event type.",
			//         "insertionOrder": false,
			//         "items": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "Key": {
			//               "maxLength": 128,
			//               "minLength": 1,
			//               "type": "string"
			//             },
			//             "Value": {
			//               "maxLength": 256,
			//               "minLength": 0,
			//               "type": "string"
			//             }
			//           },
			//           "required": [
			//             "Key",
			//             "Value"
			//           ],
			//           "type": "object"
			//         },
			//         "maxItems": 200,
			//         "type": "array",
			//         "uniqueItems": false
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "minItems": 2,
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"arn": {
						// Property: Arn
						Type:     types.StringType,
						Computed: true,
					},
					"created_time": {
						// Property: CreatedTime
						Description: "The time when the event type was created.",
						Type:        types.StringType,
						Computed:    true,
					},
					"description": {
						// Property: Description
						Description: "The description.",
						Type:        types.StringType,
						Computed:    true,
					},
					"inline": {
						// Property: Inline
						Type:     types.BoolType,
						Computed: true,
					},
					"last_updated_time": {
						// Property: LastUpdatedTime
						Description: "The time when the event type was last updated.",
						Type:        types.StringType,
						Computed:    true,
					},
					"name": {
						// Property: Name
						Type:     types.StringType,
						Computed: true,
					},
					"tags": {
						// Property: Tags
						Description: "Tags associated with this event type.",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"key": {
									// Property: Key
									Type:     types.StringType,
									Computed: true,
								},
								"value": {
									// Property: Value
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"last_updated_time": {
			// Property: LastUpdatedTime
			// CloudFormation resource type schema:
			// {
			//   "description": "The time when the event type was last updated.",
			//   "type": "string"
			// }
			Description: "The time when the event type was last updated.",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name for the event type",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "pattern": "^[0-9a-z_-]+$",
			//   "type": "string"
			// }
			Description: "The name for the event type",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "Tags associated with this event type.",
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
			//   "maxItems": 200,
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "Tags associated with this event type.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Computed: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Computed: true,
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
		Description: "Data Source schema for AWS::FraudDetector::EventType",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::FraudDetector::EventType").WithTerraformTypeName("awscc_frauddetector_event_type")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":               "Arn",
		"created_time":      "CreatedTime",
		"data_source":       "DataSource",
		"data_type":         "DataType",
		"default_value":     "DefaultValue",
		"description":       "Description",
		"entity_types":      "EntityTypes",
		"event_variables":   "EventVariables",
		"inline":            "Inline",
		"key":               "Key",
		"labels":            "Labels",
		"last_updated_time": "LastUpdatedTime",
		"name":              "Name",
		"tags":              "Tags",
		"value":             "Value",
		"variable_type":     "VariableType",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
