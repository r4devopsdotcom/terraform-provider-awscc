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
	registry.AddDataSourceTypeFactory("awscc_ec2_capacity_reservation_fleet", capacityReservationFleetDataSourceType)
}

// capacityReservationFleetDataSourceType returns the Terraform awscc_ec2_capacity_reservation_fleet data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::EC2::CapacityReservationFleet resource type.
func capacityReservationFleetDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"allocation_strategy": {
			// Property: AllocationStrategy
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"capacity_reservation_fleet_id": {
			// Property: CapacityReservationFleetId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"end_date": {
			// Property: EndDate
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"instance_match_criteria": {
			// Property: InstanceMatchCriteria
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "open"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"instance_type_specifications": {
			// Property: InstanceTypeSpecifications
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "AvailabilityZone": {
			//         "type": "string"
			//       },
			//       "AvailabilityZoneId": {
			//         "type": "string"
			//       },
			//       "EbsOptimized": {
			//         "type": "boolean"
			//       },
			//       "InstancePlatform": {
			//         "type": "string"
			//       },
			//       "InstanceType": {
			//         "type": "string"
			//       },
			//       "Priority": {
			//         "maximum": 999,
			//         "minimum": 0,
			//         "type": "integer"
			//       },
			//       "Weight": {
			//         "type": "number"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"availability_zone": {
						// Property: AvailabilityZone
						Type:     types.StringType,
						Computed: true,
					},
					"availability_zone_id": {
						// Property: AvailabilityZoneId
						Type:     types.StringType,
						Computed: true,
					},
					"ebs_optimized": {
						// Property: EbsOptimized
						Type:     types.BoolType,
						Computed: true,
					},
					"instance_platform": {
						// Property: InstancePlatform
						Type:     types.StringType,
						Computed: true,
					},
					"instance_type": {
						// Property: InstanceType
						Type:     types.StringType,
						Computed: true,
					},
					"priority": {
						// Property: Priority
						Type:     types.NumberType,
						Computed: true,
					},
					"weight": {
						// Property: Weight
						Type:     types.NumberType,
						Computed: true,
					},
				},
				tfsdk.SetNestedAttributesOptions{},
			),
			Computed: true,
		},
		"no_remove_end_date": {
			// Property: NoRemoveEndDate
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Computed: true,
		},
		"remove_end_date": {
			// Property: RemoveEndDate
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Computed: true,
		},
		"tag_specifications": {
			// Property: TagSpecifications
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "ResourceType": {
			//         "type": "string"
			//       },
			//       "Tags": {
			//         "insertionOrder": false,
			//         "items": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "Key": {
			//               "type": "string"
			//             },
			//             "Value": {
			//               "type": "string"
			//             }
			//           },
			//           "required": [
			//             "Value",
			//             "Key"
			//           ],
			//           "type": "object"
			//         },
			//         "type": "array",
			//         "uniqueItems": false
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"resource_type": {
						// Property: ResourceType
						Type:     types.StringType,
						Computed: true,
					},
					"tags": {
						// Property: Tags
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
							tfsdk.ListNestedAttributesOptions{},
						),
						Computed: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
		"tenancy": {
			// Property: Tenancy
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "default"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"total_target_capacity": {
			// Property: TotalTargetCapacity
			// CloudFormation resource type schema:
			// {
			//   "maximum": 25000,
			//   "minimum": 1,
			//   "type": "integer"
			// }
			Type:     types.NumberType,
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::EC2::CapacityReservationFleet",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::CapacityReservationFleet").WithTerraformTypeName("awscc_ec2_capacity_reservation_fleet")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"allocation_strategy":           "AllocationStrategy",
		"availability_zone":             "AvailabilityZone",
		"availability_zone_id":          "AvailabilityZoneId",
		"capacity_reservation_fleet_id": "CapacityReservationFleetId",
		"ebs_optimized":                 "EbsOptimized",
		"end_date":                      "EndDate",
		"instance_match_criteria":       "InstanceMatchCriteria",
		"instance_platform":             "InstancePlatform",
		"instance_type":                 "InstanceType",
		"instance_type_specifications":  "InstanceTypeSpecifications",
		"key":                           "Key",
		"no_remove_end_date":            "NoRemoveEndDate",
		"priority":                      "Priority",
		"remove_end_date":               "RemoveEndDate",
		"resource_type":                 "ResourceType",
		"tag_specifications":            "TagSpecifications",
		"tags":                          "Tags",
		"tenancy":                       "Tenancy",
		"total_target_capacity":         "TotalTargetCapacity",
		"value":                         "Value",
		"weight":                        "Weight",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
