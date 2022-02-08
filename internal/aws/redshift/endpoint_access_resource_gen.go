// Code generated by generators/resource/main.go; DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("awscc_redshift_endpoint_access", endpointAccessResourceType)
}

// endpointAccessResourceType returns the Terraform awscc_redshift_endpoint_access resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Redshift::EndpointAccess resource type.
func endpointAccessResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"address": {
			// Property: Address
			// CloudFormation resource type schema:
			// {
			//   "description": "The DNS address of the endpoint.",
			//   "type": "string"
			// }
			Description: "The DNS address of the endpoint.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"cluster_identifier": {
			// Property: ClusterIdentifier
			// CloudFormation resource type schema:
			// {
			//   "description": "A unique identifier for the cluster. You use this identifier to refer to the cluster for any subsequent cluster operations such as deleting or modifying. All alphabetical characters must be lower case, no hypens at the end, no two consecutive hyphens. Cluster name should be unique for all clusters within an AWS account",
			//   "type": "string"
			// }
			Description: "A unique identifier for the cluster. You use this identifier to refer to the cluster for any subsequent cluster operations such as deleting or modifying. All alphabetical characters must be lower case, no hypens at the end, no two consecutive hyphens. Cluster name should be unique for all clusters within an AWS account",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
				tfsdk.RequiresReplace(),
			},
		},
		"endpoint_create_time": {
			// Property: EndpointCreateTime
			// CloudFormation resource type schema:
			// {
			//   "description": "The time (UTC) that the endpoint was created.",
			//   "type": "string"
			// }
			Description: "The time (UTC) that the endpoint was created.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"endpoint_name": {
			// Property: EndpointName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the endpoint.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the endpoint.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"endpoint_status": {
			// Property: EndpointStatus
			// CloudFormation resource type schema:
			// {
			//   "description": "The status of the endpoint.",
			//   "type": "string"
			// }
			Description: "The status of the endpoint.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"port": {
			// Property: Port
			// CloudFormation resource type schema:
			// {
			//   "description": "The port number on which the cluster accepts incoming connections.",
			//   "type": "integer"
			// }
			Description: "The port number on which the cluster accepts incoming connections.",
			Type:        types.Int64Type,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"resource_owner": {
			// Property: ResourceOwner
			// CloudFormation resource type schema:
			// {
			//   "description": "The AWS account ID of the owner of the cluster.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The AWS account ID of the owner of the cluster.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
				tfsdk.RequiresReplace(),
			},
		},
		"subnet_group_name": {
			// Property: SubnetGroupName
			// CloudFormation resource type schema:
			// {
			//   "description": "The subnet group name where Amazon Redshift chooses to deploy the endpoint.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The subnet group name where Amazon Redshift chooses to deploy the endpoint.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
				tfsdk.RequiresReplace(),
			},
		},
		"vpc_endpoint": {
			// Property: VpcEndpoint
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The connection endpoint for connecting to an Amazon Redshift cluster through the proxy.",
			//   "properties": {
			//     "NetworkInterfaces": {
			//       "description": "One or more network interfaces of the endpoint. Also known as an interface endpoint.",
			//       "insertionOrder": false,
			//       "items": {
			//         "additionalProperties": false,
			//         "description": "Describes a network interface.",
			//         "properties": {
			//           "AvailabilityZone": {
			//             "description": "The Availability Zone.",
			//             "type": "string"
			//           },
			//           "NetworkInterfaceId": {
			//             "description": "The network interface identifier.",
			//             "type": "string"
			//           },
			//           "PrivateIpAddress": {
			//             "description": "The IPv4 address of the network interface within the subnet.",
			//             "type": "string"
			//           },
			//           "SubnetId": {
			//             "description": "The subnet identifier.",
			//             "type": "string"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "type": "array"
			//     },
			//     "VpcEndpointId": {
			//       "description": "The connection endpoint ID for connecting an Amazon Redshift cluster through the proxy.",
			//       "type": "string"
			//     },
			//     "VpcId": {
			//       "description": "The VPC identifier that the endpoint is associated.",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The connection endpoint for connecting to an Amazon Redshift cluster through the proxy.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"network_interfaces": {
						// Property: NetworkInterfaces
						Description: "One or more network interfaces of the endpoint. Also known as an interface endpoint.",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"availability_zone": {
									// Property: AvailabilityZone
									Description: "The Availability Zone.",
									Type:        types.StringType,
									Optional:    true,
								},
								"network_interface_id": {
									// Property: NetworkInterfaceId
									Description: "The network interface identifier.",
									Type:        types.StringType,
									Optional:    true,
								},
								"private_ip_address": {
									// Property: PrivateIpAddress
									Description: "The IPv4 address of the network interface within the subnet.",
									Type:        types.StringType,
									Optional:    true,
								},
								"subnet_id": {
									// Property: SubnetId
									Description: "The subnet identifier.",
									Type:        types.StringType,
									Optional:    true,
								},
							},
							tfsdk.ListNestedAttributesOptions{},
						),
						Optional: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							Multiset(),
						},
					},
					"vpc_endpoint_id": {
						// Property: VpcEndpointId
						Description: "The connection endpoint ID for connecting an Amazon Redshift cluster through the proxy.",
						Type:        types.StringType,
						Optional:    true,
					},
					"vpc_id": {
						// Property: VpcId
						Description: "The VPC identifier that the endpoint is associated.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"vpc_security_group_ids": {
			// Property: VpcSecurityGroupIds
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of vpc security group ids to apply to the created endpoint access.",
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "A list of vpc security group ids to apply to the created endpoint access.",
			Type:        types.ListType{ElemType: types.StringType},
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
			},
		},
		"vpc_security_groups": {
			// Property: VpcSecurityGroups
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of Virtual Private Cloud (VPC) security groups to be associated with the endpoint.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "Describes the members of a VPC security group.",
			//     "properties": {
			//       "Status": {
			//         "description": "The status of the VPC security group.",
			//         "type": "string"
			//       },
			//       "VpcSecurityGroupId": {
			//         "description": "The identifier of the VPC security group.",
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "A list of Virtual Private Cloud (VPC) security groups to be associated with the endpoint.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"status": {
						// Property: Status
						Description: "The status of the VPC security group.",
						Type:        types.StringType,
						Optional:    true,
					},
					"vpc_security_group_id": {
						// Property: VpcSecurityGroupId
						Description: "The identifier of the VPC security group.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
				tfsdk.UseStateForUnknown(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			tfsdk.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource schema for a Redshift-managed VPC endpoint.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Redshift::EndpointAccess").WithTerraformTypeName("awscc_redshift_endpoint_access")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"address":                "Address",
		"availability_zone":      "AvailabilityZone",
		"cluster_identifier":     "ClusterIdentifier",
		"endpoint_create_time":   "EndpointCreateTime",
		"endpoint_name":          "EndpointName",
		"endpoint_status":        "EndpointStatus",
		"network_interface_id":   "NetworkInterfaceId",
		"network_interfaces":     "NetworkInterfaces",
		"port":                   "Port",
		"private_ip_address":     "PrivateIpAddress",
		"resource_owner":         "ResourceOwner",
		"status":                 "Status",
		"subnet_group_name":      "SubnetGroupName",
		"subnet_id":              "SubnetId",
		"vpc_endpoint":           "VpcEndpoint",
		"vpc_endpoint_id":        "VpcEndpointId",
		"vpc_id":                 "VpcId",
		"vpc_security_group_id":  "VpcSecurityGroupId",
		"vpc_security_group_ids": "VpcSecurityGroupIds",
		"vpc_security_groups":    "VpcSecurityGroups",
	})

	opts = opts.WithCreateTimeoutInMinutes(60).WithDeleteTimeoutInMinutes(60)

	opts = opts.WithUpdateTimeoutInMinutes(60)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
