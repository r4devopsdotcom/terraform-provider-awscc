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
	registry.AddDataSourceTypeFactory("awscc_ec2_network_insights_access_scope", networkInsightsAccessScopeDataSourceType)
}

// networkInsightsAccessScopeDataSourceType returns the Terraform awscc_ec2_network_insights_access_scope data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::EC2::NetworkInsightsAccessScope resource type.
func networkInsightsAccessScopeDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"created_date": {
			// Property: CreatedDate
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"exclude_paths": {
			// Property: ExcludePaths
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": true,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Destination": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "PacketHeaderStatement": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "DestinationAddresses": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "DestinationPorts": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "DestinationPrefixLists": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "Protocols": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "enum": [
			//                     "tcp",
			//                     "udp"
			//                   ],
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "SourceAddresses": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "SourcePorts": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "SourcePrefixLists": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "ResourceStatement": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "ResourceTypes": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "Resources": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               }
			//             },
			//             "type": "object"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "Source": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "PacketHeaderStatement": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "DestinationAddresses": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "DestinationPorts": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "DestinationPrefixLists": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "Protocols": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "enum": [
			//                     "tcp",
			//                     "udp"
			//                   ],
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "SourceAddresses": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "SourcePorts": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "SourcePrefixLists": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "ResourceStatement": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "ResourceTypes": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "Resources": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               }
			//             },
			//             "type": "object"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "ThroughResources": {
			//         "insertionOrder": true,
			//         "items": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "ResourceStatement": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "ResourceTypes": {
			//                   "insertionOrder": true,
			//                   "items": {
			//                     "type": "string"
			//                   },
			//                   "type": "array"
			//                 },
			//                 "Resources": {
			//                   "insertionOrder": true,
			//                   "items": {
			//                     "type": "string"
			//                   },
			//                   "type": "array"
			//                 }
			//               },
			//               "type": "object"
			//             }
			//           },
			//           "type": "object"
			//         },
			//         "type": "array"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"destination": {
						// Property: Destination
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"packet_header_statement": {
									// Property: PacketHeaderStatement
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"destination_addresses": {
												// Property: DestinationAddresses
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
											"destination_ports": {
												// Property: DestinationPorts
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
											"destination_prefix_lists": {
												// Property: DestinationPrefixLists
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
											"protocols": {
												// Property: Protocols
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
											"source_addresses": {
												// Property: SourceAddresses
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
											"source_ports": {
												// Property: SourcePorts
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
											"source_prefix_lists": {
												// Property: SourcePrefixLists
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"resource_statement": {
									// Property: ResourceStatement
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"resource_types": {
												// Property: ResourceTypes
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
											"resources": {
												// Property: Resources
												Type:     types.ListType{ElemType: types.StringType},
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
					"source": {
						// Property: Source
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"packet_header_statement": {
									// Property: PacketHeaderStatement
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"destination_addresses": {
												// Property: DestinationAddresses
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
											"destination_ports": {
												// Property: DestinationPorts
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
											"destination_prefix_lists": {
												// Property: DestinationPrefixLists
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
											"protocols": {
												// Property: Protocols
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
											"source_addresses": {
												// Property: SourceAddresses
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
											"source_ports": {
												// Property: SourcePorts
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
											"source_prefix_lists": {
												// Property: SourcePrefixLists
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"resource_statement": {
									// Property: ResourceStatement
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"resource_types": {
												// Property: ResourceTypes
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
											"resources": {
												// Property: Resources
												Type:     types.ListType{ElemType: types.StringType},
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
					"through_resources": {
						// Property: ThroughResources
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"resource_statement": {
									// Property: ResourceStatement
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"resource_types": {
												// Property: ResourceTypes
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
											"resources": {
												// Property: Resources
												Type:     types.ListType{ElemType: types.StringType},
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
			),
			Computed: true,
		},
		"match_paths": {
			// Property: MatchPaths
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": true,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Destination": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "PacketHeaderStatement": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "DestinationAddresses": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "DestinationPorts": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "DestinationPrefixLists": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "Protocols": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "enum": [
			//                     "tcp",
			//                     "udp"
			//                   ],
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "SourceAddresses": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "SourcePorts": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "SourcePrefixLists": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "ResourceStatement": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "ResourceTypes": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "Resources": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               }
			//             },
			//             "type": "object"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "Source": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "PacketHeaderStatement": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "DestinationAddresses": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "DestinationPorts": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "DestinationPrefixLists": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "Protocols": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "enum": [
			//                     "tcp",
			//                     "udp"
			//                   ],
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "SourceAddresses": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "SourcePorts": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "SourcePrefixLists": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "ResourceStatement": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "ResourceTypes": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "Resources": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               }
			//             },
			//             "type": "object"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "ThroughResources": {
			//         "insertionOrder": true,
			//         "items": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "ResourceStatement": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "ResourceTypes": {
			//                   "insertionOrder": true,
			//                   "items": {
			//                     "type": "string"
			//                   },
			//                   "type": "array"
			//                 },
			//                 "Resources": {
			//                   "insertionOrder": true,
			//                   "items": {
			//                     "type": "string"
			//                   },
			//                   "type": "array"
			//                 }
			//               },
			//               "type": "object"
			//             }
			//           },
			//           "type": "object"
			//         },
			//         "type": "array"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"destination": {
						// Property: Destination
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"packet_header_statement": {
									// Property: PacketHeaderStatement
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"destination_addresses": {
												// Property: DestinationAddresses
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
											"destination_ports": {
												// Property: DestinationPorts
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
											"destination_prefix_lists": {
												// Property: DestinationPrefixLists
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
											"protocols": {
												// Property: Protocols
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
											"source_addresses": {
												// Property: SourceAddresses
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
											"source_ports": {
												// Property: SourcePorts
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
											"source_prefix_lists": {
												// Property: SourcePrefixLists
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"resource_statement": {
									// Property: ResourceStatement
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"resource_types": {
												// Property: ResourceTypes
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
											"resources": {
												// Property: Resources
												Type:     types.ListType{ElemType: types.StringType},
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
					"source": {
						// Property: Source
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"packet_header_statement": {
									// Property: PacketHeaderStatement
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"destination_addresses": {
												// Property: DestinationAddresses
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
											"destination_ports": {
												// Property: DestinationPorts
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
											"destination_prefix_lists": {
												// Property: DestinationPrefixLists
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
											"protocols": {
												// Property: Protocols
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
											"source_addresses": {
												// Property: SourceAddresses
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
											"source_ports": {
												// Property: SourcePorts
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
											"source_prefix_lists": {
												// Property: SourcePrefixLists
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"resource_statement": {
									// Property: ResourceStatement
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"resource_types": {
												// Property: ResourceTypes
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
											"resources": {
												// Property: Resources
												Type:     types.ListType{ElemType: types.StringType},
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
					"through_resources": {
						// Property: ThroughResources
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"resource_statement": {
									// Property: ResourceStatement
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"resource_types": {
												// Property: ResourceTypes
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
											"resources": {
												// Property: Resources
												Type:     types.ListType{ElemType: types.StringType},
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
			),
			Computed: true,
		},
		"network_insights_access_scope_arn": {
			// Property: NetworkInsightsAccessScopeArn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"network_insights_access_scope_id": {
			// Property: NetworkInsightsAccessScopeId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "type": "string"
			//       },
			//       "Value": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
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
		"updated_date": {
			// Property: UpdatedDate
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
		Description: "Data Source schema for AWS::EC2::NetworkInsightsAccessScope",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::NetworkInsightsAccessScope").WithTerraformTypeName("awscc_ec2_network_insights_access_scope")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"created_date":                      "CreatedDate",
		"destination":                       "Destination",
		"destination_addresses":             "DestinationAddresses",
		"destination_ports":                 "DestinationPorts",
		"destination_prefix_lists":          "DestinationPrefixLists",
		"exclude_paths":                     "ExcludePaths",
		"key":                               "Key",
		"match_paths":                       "MatchPaths",
		"network_insights_access_scope_arn": "NetworkInsightsAccessScopeArn",
		"network_insights_access_scope_id":  "NetworkInsightsAccessScopeId",
		"packet_header_statement":           "PacketHeaderStatement",
		"protocols":                         "Protocols",
		"resource_statement":                "ResourceStatement",
		"resource_types":                    "ResourceTypes",
		"resources":                         "Resources",
		"source":                            "Source",
		"source_addresses":                  "SourceAddresses",
		"source_ports":                      "SourcePorts",
		"source_prefix_lists":               "SourcePrefixLists",
		"tags":                              "Tags",
		"through_resources":                 "ThroughResources",
		"updated_date":                      "UpdatedDate",
		"value":                             "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
