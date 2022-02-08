// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_pinpoint_in_app_template", inAppTemplateDataSourceType)
}

// inAppTemplateDataSourceType returns the Terraform awscc_pinpoint_in_app_template data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Pinpoint::InAppTemplate resource type.
func inAppTemplateDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"content": {
			// Property: Content
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": true,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "BackgroundColor": {
			//         "type": "string"
			//       },
			//       "BodyConfig": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "Alignment": {
			//             "enum": [
			//               "LEFT",
			//               "CENTER",
			//               "RIGHT"
			//             ],
			//             "type": "string"
			//           },
			//           "Body": {
			//             "type": "string"
			//           },
			//           "TextColor": {
			//             "type": "string"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "HeaderConfig": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "Alignment": {
			//             "enum": [
			//               "LEFT",
			//               "CENTER",
			//               "RIGHT"
			//             ],
			//             "type": "string"
			//           },
			//           "Header": {
			//             "type": "string"
			//           },
			//           "TextColor": {
			//             "type": "string"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "ImageUrl": {
			//         "type": "string"
			//       },
			//       "PrimaryBtn": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "Android": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "ButtonAction": {
			//                 "enum": [
			//                   "LINK",
			//                   "DEEP_LINK",
			//                   "CLOSE"
			//                 ],
			//                 "type": "string"
			//               },
			//               "Link": {
			//                 "type": "string"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "DefaultConfig": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "BackgroundColor": {
			//                 "type": "string"
			//               },
			//               "BorderRadius": {
			//                 "type": "integer"
			//               },
			//               "ButtonAction": {
			//                 "enum": [
			//                   "LINK",
			//                   "DEEP_LINK",
			//                   "CLOSE"
			//                 ],
			//                 "type": "string"
			//               },
			//               "Link": {
			//                 "type": "string"
			//               },
			//               "Text": {
			//                 "type": "string"
			//               },
			//               "TextColor": {
			//                 "type": "string"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "IOS": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "ButtonAction": {
			//                 "enum": [
			//                   "LINK",
			//                   "DEEP_LINK",
			//                   "CLOSE"
			//                 ],
			//                 "type": "string"
			//               },
			//               "Link": {
			//                 "type": "string"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "Web": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "ButtonAction": {
			//                 "enum": [
			//                   "LINK",
			//                   "DEEP_LINK",
			//                   "CLOSE"
			//                 ],
			//                 "type": "string"
			//               },
			//               "Link": {
			//                 "type": "string"
			//               }
			//             },
			//             "type": "object"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "SecondaryBtn": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "Android": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "ButtonAction": {
			//                 "enum": [
			//                   "LINK",
			//                   "DEEP_LINK",
			//                   "CLOSE"
			//                 ],
			//                 "type": "string"
			//               },
			//               "Link": {
			//                 "type": "string"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "DefaultConfig": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "BackgroundColor": {
			//                 "type": "string"
			//               },
			//               "BorderRadius": {
			//                 "type": "integer"
			//               },
			//               "ButtonAction": {
			//                 "enum": [
			//                   "LINK",
			//                   "DEEP_LINK",
			//                   "CLOSE"
			//                 ],
			//                 "type": "string"
			//               },
			//               "Link": {
			//                 "type": "string"
			//               },
			//               "Text": {
			//                 "type": "string"
			//               },
			//               "TextColor": {
			//                 "type": "string"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "IOS": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "ButtonAction": {
			//                 "enum": [
			//                   "LINK",
			//                   "DEEP_LINK",
			//                   "CLOSE"
			//                 ],
			//                 "type": "string"
			//               },
			//               "Link": {
			//                 "type": "string"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "Web": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "ButtonAction": {
			//                 "enum": [
			//                   "LINK",
			//                   "DEEP_LINK",
			//                   "CLOSE"
			//                 ],
			//                 "type": "string"
			//               },
			//               "Link": {
			//                 "type": "string"
			//               }
			//             },
			//             "type": "object"
			//           }
			//         },
			//         "type": "object"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"background_color": {
						// Property: BackgroundColor
						Type:     types.StringType,
						Computed: true,
					},
					"body_config": {
						// Property: BodyConfig
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"alignment": {
									// Property: Alignment
									Type:     types.StringType,
									Computed: true,
								},
								"body": {
									// Property: Body
									Type:     types.StringType,
									Computed: true,
								},
								"text_color": {
									// Property: TextColor
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"header_config": {
						// Property: HeaderConfig
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"alignment": {
									// Property: Alignment
									Type:     types.StringType,
									Computed: true,
								},
								"header": {
									// Property: Header
									Type:     types.StringType,
									Computed: true,
								},
								"text_color": {
									// Property: TextColor
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"image_url": {
						// Property: ImageUrl
						Type:     types.StringType,
						Computed: true,
					},
					"primary_btn": {
						// Property: PrimaryBtn
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"android": {
									// Property: Android
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"button_action": {
												// Property: ButtonAction
												Type:     types.StringType,
												Computed: true,
											},
											"link": {
												// Property: Link
												Type:     types.StringType,
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"default_config": {
									// Property: DefaultConfig
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"background_color": {
												// Property: BackgroundColor
												Type:     types.StringType,
												Computed: true,
											},
											"border_radius": {
												// Property: BorderRadius
												Type:     types.Int64Type,
												Computed: true,
											},
											"button_action": {
												// Property: ButtonAction
												Type:     types.StringType,
												Computed: true,
											},
											"link": {
												// Property: Link
												Type:     types.StringType,
												Computed: true,
											},
											"text": {
												// Property: Text
												Type:     types.StringType,
												Computed: true,
											},
											"text_color": {
												// Property: TextColor
												Type:     types.StringType,
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"ios": {
									// Property: IOS
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"button_action": {
												// Property: ButtonAction
												Type:     types.StringType,
												Computed: true,
											},
											"link": {
												// Property: Link
												Type:     types.StringType,
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"web": {
									// Property: Web
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"button_action": {
												// Property: ButtonAction
												Type:     types.StringType,
												Computed: true,
											},
											"link": {
												// Property: Link
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
					"secondary_btn": {
						// Property: SecondaryBtn
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"android": {
									// Property: Android
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"button_action": {
												// Property: ButtonAction
												Type:     types.StringType,
												Computed: true,
											},
											"link": {
												// Property: Link
												Type:     types.StringType,
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"default_config": {
									// Property: DefaultConfig
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"background_color": {
												// Property: BackgroundColor
												Type:     types.StringType,
												Computed: true,
											},
											"border_radius": {
												// Property: BorderRadius
												Type:     types.Int64Type,
												Computed: true,
											},
											"button_action": {
												// Property: ButtonAction
												Type:     types.StringType,
												Computed: true,
											},
											"link": {
												// Property: Link
												Type:     types.StringType,
												Computed: true,
											},
											"text": {
												// Property: Text
												Type:     types.StringType,
												Computed: true,
											},
											"text_color": {
												// Property: TextColor
												Type:     types.StringType,
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"ios": {
									// Property: IOS
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"button_action": {
												// Property: ButtonAction
												Type:     types.StringType,
												Computed: true,
											},
											"link": {
												// Property: Link
												Type:     types.StringType,
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"web": {
									// Property: Web
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"button_action": {
												// Property: ButtonAction
												Type:     types.StringType,
												Computed: true,
											},
											"link": {
												// Property: Link
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
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
		"custom_config": {
			// Property: CustomConfig
			// CloudFormation resource type schema:
			// {
			//   "type": "object"
			// }
			Type:     types.MapType{ElemType: types.StringType},
			Computed: true,
		},
		"layout": {
			// Property: Layout
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "BOTTOM_BANNER",
			//     "TOP_BANNER",
			//     "OVERLAYS",
			//     "MOBILE_FEED",
			//     "MIDDLE_BANNER",
			//     "CAROUSEL"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "type": "object"
			// }
			Type:     types.MapType{ElemType: types.StringType},
			Computed: true,
		},
		"template_description": {
			// Property: TemplateDescription
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"template_name": {
			// Property: TemplateName
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::Pinpoint::InAppTemplate",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Pinpoint::InAppTemplate").WithTerraformTypeName("awscc_pinpoint_in_app_template")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"alignment":            "Alignment",
		"android":              "Android",
		"arn":                  "Arn",
		"background_color":     "BackgroundColor",
		"body":                 "Body",
		"body_config":          "BodyConfig",
		"border_radius":        "BorderRadius",
		"button_action":        "ButtonAction",
		"content":              "Content",
		"custom_config":        "CustomConfig",
		"default_config":       "DefaultConfig",
		"header":               "Header",
		"header_config":        "HeaderConfig",
		"image_url":            "ImageUrl",
		"ios":                  "IOS",
		"layout":               "Layout",
		"link":                 "Link",
		"primary_btn":          "PrimaryBtn",
		"secondary_btn":        "SecondaryBtn",
		"tags":                 "Tags",
		"template_description": "TemplateDescription",
		"template_name":        "TemplateName",
		"text":                 "Text",
		"text_color":           "TextColor",
		"web":                  "Web",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
