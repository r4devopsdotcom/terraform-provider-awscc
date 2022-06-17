// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_gamelift_game_server_group", gameServerGroupDataSourceType)
}

// gameServerGroupDataSourceType returns the Terraform awscc_gamelift_game_server_group data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::GameLift::GameServerGroup resource type.
func gameServerGroupDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"auto_scaling_group_arn": {
			// Property: AutoScalingGroupArn
			// CloudFormation resource type schema:
			// {
			//   "description": "A generated unique ID for the EC2 Auto Scaling group that is associated with this game server group.",
			//   "maxLength": 256,
			//   "minLength": 0,
			//   "pattern": "[ -퟿-�𐀀-􏿿\r\n\t]*",
			//   "type": "string"
			// }
			Description: "A generated unique ID for the EC2 Auto Scaling group that is associated with this game server group.",
			Type:        types.StringType,
			Computed:    true,
		},
		"auto_scaling_policy": {
			// Property: AutoScalingPolicy
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Configuration settings to define a scaling policy for the Auto Scaling group that is optimized for game hosting",
			//   "properties": {
			//     "EstimatedInstanceWarmup": {
			//       "description": "Length of time, in seconds, it takes for a new instance to start new game server processes and register with GameLift FleetIQ.",
			//       "type": "number"
			//     },
			//     "TargetTrackingConfiguration": {
			//       "additionalProperties": false,
			//       "description": "Settings for a target-based scaling policy applied to Auto Scaling group.",
			//       "properties": {
			//         "TargetValue": {
			//           "description": "Desired value to use with a game server group target-based scaling policy.",
			//           "type": "number"
			//         }
			//       },
			//       "required": [
			//         "TargetValue"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "required": [
			//     "TargetTrackingConfiguration"
			//   ],
			//   "type": "object"
			// }
			Description: "Configuration settings to define a scaling policy for the Auto Scaling group that is optimized for game hosting",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"estimated_instance_warmup": {
						// Property: EstimatedInstanceWarmup
						Description: "Length of time, in seconds, it takes for a new instance to start new game server processes and register with GameLift FleetIQ.",
						Type:        types.Float64Type,
						Computed:    true,
					},
					"target_tracking_configuration": {
						// Property: TargetTrackingConfiguration
						Description: "Settings for a target-based scaling policy applied to Auto Scaling group.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"target_value": {
									// Property: TargetValue
									Description: "Desired value to use with a game server group target-based scaling policy.",
									Type:        types.Float64Type,
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
		"balancing_strategy": {
			// Property: BalancingStrategy
			// CloudFormation resource type schema:
			// {
			//   "description": "The fallback balancing method to use for the game server group when Spot Instances in a Region become unavailable or are not viable for game hosting.",
			//   "enum": [
			//     "SPOT_ONLY",
			//     "SPOT_PREFERRED",
			//     "ON_DEMAND_ONLY"
			//   ],
			//   "type": "string"
			// }
			Description: "The fallback balancing method to use for the game server group when Spot Instances in a Region become unavailable or are not viable for game hosting.",
			Type:        types.StringType,
			Computed:    true,
		},
		"delete_option": {
			// Property: DeleteOption
			// CloudFormation resource type schema:
			// {
			//   "description": "The type of delete to perform.",
			//   "enum": [
			//     "SAFE_DELETE",
			//     "FORCE_DELETE",
			//     "RETAIN"
			//   ],
			//   "type": "string"
			// }
			Description: "The type of delete to perform.",
			Type:        types.StringType,
			Computed:    true,
		},
		"game_server_group_arn": {
			// Property: GameServerGroupArn
			// CloudFormation resource type schema:
			// {
			//   "description": "A generated unique ID for the game server group.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "^arn:.*:gameservergroup\\/[a-zA-Z0-9-\\.]*",
			//   "type": "string"
			// }
			Description: "A generated unique ID for the game server group.",
			Type:        types.StringType,
			Computed:    true,
		},
		"game_server_group_name": {
			// Property: GameServerGroupName
			// CloudFormation resource type schema:
			// {
			//   "description": "An identifier for the new game server group.",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "[a-zA-Z0-9-\\.]+",
			//   "type": "string"
			// }
			Description: "An identifier for the new game server group.",
			Type:        types.StringType,
			Computed:    true,
		},
		"game_server_protection_policy": {
			// Property: GameServerProtectionPolicy
			// CloudFormation resource type schema:
			// {
			//   "description": "A flag that indicates whether instances in the game server group are protected from early termination.",
			//   "enum": [
			//     "NO_PROTECTION",
			//     "FULL_PROTECTION"
			//   ],
			//   "type": "string"
			// }
			Description: "A flag that indicates whether instances in the game server group are protected from early termination.",
			Type:        types.StringType,
			Computed:    true,
		},
		"instance_definitions": {
			// Property: InstanceDefinitions
			// CloudFormation resource type schema:
			// {
			//   "description": "A set of EC2 instance types to use when creating instances in the group.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "An allowed instance type for your game server group.",
			//     "properties": {
			//       "InstanceType": {
			//         "description": "An EC2 instance type designation.",
			//         "type": "string"
			//       },
			//       "WeightedCapacity": {
			//         "description": "Instance weighting that indicates how much this instance type contributes to the total capacity of a game server group.",
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "InstanceType"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 20,
			//   "minItems": 2,
			//   "type": "array"
			// }
			Description: "A set of EC2 instance types to use when creating instances in the group.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"instance_type": {
						// Property: InstanceType
						Description: "An EC2 instance type designation.",
						Type:        types.StringType,
						Computed:    true,
					},
					"weighted_capacity": {
						// Property: WeightedCapacity
						Description: "Instance weighting that indicates how much this instance type contributes to the total capacity of a game server group.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"launch_template": {
			// Property: LaunchTemplate
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The EC2 launch template that contains configuration settings and game server code to be deployed to all instances in the game server group.",
			//   "properties": {
			//     "LaunchTemplateId": {
			//       "description": "A unique identifier for an existing EC2 launch template.",
			//       "type": "string"
			//     },
			//     "LaunchTemplateName": {
			//       "description": "A readable identifier for an existing EC2 launch template.",
			//       "type": "string"
			//     },
			//     "Version": {
			//       "description": "The version of the EC2 launch template to use.",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The EC2 launch template that contains configuration settings and game server code to be deployed to all instances in the game server group.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"launch_template_id": {
						// Property: LaunchTemplateId
						Description: "A unique identifier for an existing EC2 launch template.",
						Type:        types.StringType,
						Computed:    true,
					},
					"launch_template_name": {
						// Property: LaunchTemplateName
						Description: "A readable identifier for an existing EC2 launch template.",
						Type:        types.StringType,
						Computed:    true,
					},
					"version": {
						// Property: Version
						Description: "The version of the EC2 launch template to use.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"max_size": {
			// Property: MaxSize
			// CloudFormation resource type schema:
			// {
			//   "description": "The maximum number of instances allowed in the EC2 Auto Scaling group.",
			//   "minimum": 1,
			//   "type": "number"
			// }
			Description: "The maximum number of instances allowed in the EC2 Auto Scaling group.",
			Type:        types.Float64Type,
			Computed:    true,
		},
		"min_size": {
			// Property: MinSize
			// CloudFormation resource type schema:
			// {
			//   "description": "The minimum number of instances allowed in the EC2 Auto Scaling group.",
			//   "minimum": 0,
			//   "type": "number"
			// }
			Description: "The minimum number of instances allowed in the EC2 Auto Scaling group.",
			Type:        types.Float64Type,
			Computed:    true,
		},
		"role_arn": {
			// Property: RoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) for an IAM role that allows Amazon GameLift to access your EC2 Auto Scaling groups.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "^arn:.*:role\\/[\\w+=,.@-]+",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) for an IAM role that allows Amazon GameLift to access your EC2 Auto Scaling groups.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of labels to assign to the new game server group resource.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "description": "The key for a developer-defined key:value pair for tagging an AWS resource.",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for a developer-defined key:value pair for tagging an AWS resource.",
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "maxItems": 200,
			//   "minItems": 0,
			//   "type": "array"
			// }
			Description: "A list of labels to assign to the new game server group resource.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key for a developer-defined key:value pair for tagging an AWS resource.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for a developer-defined key:value pair for tagging an AWS resource.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"vpc_subnets": {
			// Property: VpcSubnets
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of virtual private cloud (VPC) subnets to use with instances in the game server group.",
			//   "insertionOrder": false,
			//   "items": {
			//     "maxLength": 24,
			//     "minLength": 15,
			//     "pattern": "^subnet-[0-9a-z]+$",
			//     "type": "string"
			//   },
			//   "maxItems": 20,
			//   "minItems": 1,
			//   "type": "array"
			// }
			Description: "A list of virtual private cloud (VPC) subnets to use with instances in the game server group.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::GameLift::GameServerGroup",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::GameLift::GameServerGroup").WithTerraformTypeName("awscc_gamelift_game_server_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"auto_scaling_group_arn":        "AutoScalingGroupArn",
		"auto_scaling_policy":           "AutoScalingPolicy",
		"balancing_strategy":            "BalancingStrategy",
		"delete_option":                 "DeleteOption",
		"estimated_instance_warmup":     "EstimatedInstanceWarmup",
		"game_server_group_arn":         "GameServerGroupArn",
		"game_server_group_name":        "GameServerGroupName",
		"game_server_protection_policy": "GameServerProtectionPolicy",
		"instance_definitions":          "InstanceDefinitions",
		"instance_type":                 "InstanceType",
		"key":                           "Key",
		"launch_template":               "LaunchTemplate",
		"launch_template_id":            "LaunchTemplateId",
		"launch_template_name":          "LaunchTemplateName",
		"max_size":                      "MaxSize",
		"min_size":                      "MinSize",
		"role_arn":                      "RoleArn",
		"tags":                          "Tags",
		"target_tracking_configuration": "TargetTrackingConfiguration",
		"target_value":                  "TargetValue",
		"value":                         "Value",
		"version":                       "Version",
		"vpc_subnets":                   "VpcSubnets",
		"weighted_capacity":             "WeightedCapacity",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
