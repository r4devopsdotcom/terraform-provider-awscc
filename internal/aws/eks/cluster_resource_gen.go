// Code generated by generators/resource/main.go; DO NOT EDIT.

package eks

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_eks_cluster", clusterResourceType)
}

// clusterResourceType returns the Terraform awscc_eks_cluster resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::EKS::Cluster resource type.
func clusterResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the cluster, such as arn:aws:eks:us-west-2:666666666666:cluster/prod.",
			//   "type": "string"
			// }
			Description: "The ARN of the cluster, such as arn:aws:eks:us-west-2:666666666666:cluster/prod.",
			Type:        types.StringType,
			Computed:    true,
		},
		"certificate_authority_data": {
			// Property: CertificateAuthorityData
			// CloudFormation resource type schema:
			// {
			//   "description": "The certificate-authority-data for your cluster.",
			//   "type": "string"
			// }
			Description: "The certificate-authority-data for your cluster.",
			Type:        types.StringType,
			Computed:    true,
		},
		"cluster_security_group_id": {
			// Property: ClusterSecurityGroupId
			// CloudFormation resource type schema:
			// {
			//   "description": "The cluster security group that was created by Amazon EKS for the cluster. Managed node groups use this security group for control plane to data plane communication.",
			//   "type": "string"
			// }
			Description: "The cluster security group that was created by Amazon EKS for the cluster. Managed node groups use this security group for control plane to data plane communication.",
			Type:        types.StringType,
			Computed:    true,
		},
		"encryption_config": {
			// Property: EncryptionConfig
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "The encryption configuration for the cluster",
			//     "properties": {
			//       "Provider": {
			//         "additionalProperties": false,
			//         "description": "The encryption provider for the cluster.",
			//         "properties": {
			//           "KeyArn": {
			//             "description": "Amazon Resource Name (ARN) or alias of the KMS key. The KMS key must be symmetric, created in the same region as the cluster, and if the KMS key was created in a different account, the user must have access to the KMS key.",
			//             "type": "string"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "Resources": {
			//         "description": "Specifies the resources to be encrypted. The only supported value is \"secrets\".",
			//         "insertionOrder": false,
			//         "items": {
			//           "type": "string"
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
					"provider": {
						// Property: Provider
						Description: "The encryption provider for the cluster.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"key_arn": {
									// Property: KeyArn
									Description: "Amazon Resource Name (ARN) or alias of the KMS key. The KMS key must be symmetric, created in the same region as the cluster, and if the KMS key was created in a different account, the user must have access to the KMS key.",
									Type:        types.StringType,
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
					"resources": {
						// Property: Resources
						Description: "Specifies the resources to be encrypted. The only supported value is \"secrets\".",
						Type:        types.ListType{ElemType: types.StringType},
						Optional:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							Multiset(),
						},
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
				tfsdk.RequiresReplace(),
			},
		},
		"encryption_config_key_arn": {
			// Property: EncryptionConfigKeyArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Amazon Resource Name (ARN) or alias of the customer master key (CMK).",
			//   "type": "string"
			// }
			Description: "Amazon Resource Name (ARN) or alias of the customer master key (CMK).",
			Type:        types.StringType,
			Computed:    true,
		},
		"endpoint": {
			// Property: Endpoint
			// CloudFormation resource type schema:
			// {
			//   "description": "The endpoint for your Kubernetes API server, such as https://5E1D0CEXAMPLEA591B746AFC5AB30262.yl4.us-west-2.eks.amazonaws.com.",
			//   "type": "string"
			// }
			Description: "The endpoint for your Kubernetes API server, such as https://5E1D0CEXAMPLEA591B746AFC5AB30262.yl4.us-west-2.eks.amazonaws.com.",
			Type:        types.StringType,
			Computed:    true,
		},
		"kubernetes_network_config": {
			// Property: KubernetesNetworkConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The Kubernetes network configuration for the cluster.",
			//   "properties": {
			//     "ServiceIpv4Cidr": {
			//       "description": "The CIDR block to assign Kubernetes service IP addresses from. If you don't specify a block, Kubernetes assigns addresses from either the 10.100.0.0/16 or 172.20.0.0/16 CIDR blocks. We recommend that you specify a block that does not overlap with resources in other networks that are peered or connected to your VPC. ",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The Kubernetes network configuration for the cluster.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"service_ipv_4_cidr": {
						// Property: ServiceIpv4Cidr
						Description: "The CIDR block to assign Kubernetes service IP addresses from. If you don't specify a block, Kubernetes assigns addresses from either the 10.100.0.0/16 or 172.20.0.0/16 CIDR blocks. We recommend that you specify a block that does not overlap with resources in other networks that are peered or connected to your VPC. ",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"logging": {
			// Property: Logging
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Enable exporting the Kubernetes control plane logs for your cluster to CloudWatch Logs based on log types. By default, cluster control plane logs aren't exported to CloudWatch Logs.",
			//   "properties": {
			//     "ClusterLogging": {
			//       "additionalProperties": false,
			//       "description": "The cluster control plane logging configuration for your cluster. ",
			//       "properties": {
			//         "EnabledTypes": {
			//           "description": "Enable control plane logs for your cluster, all log types will be disabled if the array is empty",
			//           "insertionOrder": false,
			//           "items": {
			//             "additionalProperties": false,
			//             "description": "Enabled Logging Type",
			//             "properties": {
			//               "Type": {
			//                 "description": "name of the log type",
			//                 "enum": [
			//                   "api",
			//                   "audit",
			//                   "authenticator",
			//                   "controllerManager",
			//                   "scheduler"
			//                 ],
			//                 "type": "string"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "type": "array"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Enable exporting the Kubernetes control plane logs for your cluster to CloudWatch Logs based on log types. By default, cluster control plane logs aren't exported to CloudWatch Logs.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"cluster_logging": {
						// Property: ClusterLogging
						Description: "The cluster control plane logging configuration for your cluster. ",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"enabled_types": {
									// Property: EnabledTypes
									Description: "Enable control plane logs for your cluster, all log types will be disabled if the array is empty",
									Attributes: tfsdk.ListNestedAttributes(
										map[string]tfsdk.Attribute{
											"type": {
												// Property: Type
												Description: "name of the log type",
												Type:        types.StringType,
												Optional:    true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringInSlice([]string{
														"api",
														"audit",
														"authenticator",
														"controllerManager",
														"scheduler",
													}),
												},
											},
										},
										tfsdk.ListNestedAttributesOptions{},
									),
									Optional: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										Multiset(),
									},
								},
							},
						),
						Optional: true,
					},
				},
			),
			Optional: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The unique name to give to your cluster.",
			//   "maxLength": 100,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The unique name to give to your cluster.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 100),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"open_id_connect_issuer_url": {
			// Property: OpenIdConnectIssuerUrl
			// CloudFormation resource type schema:
			// {
			//   "description": "The issuer URL for the cluster's OIDC identity provider, such as https://oidc.eks.us-west-2.amazonaws.com/id/EXAMPLED539D4633E53DE1B716D3041E. If you need to remove https:// from this output value, you can include the following code in your template.",
			//   "type": "string"
			// }
			Description: "The issuer URL for the cluster's OIDC identity provider, such as https://oidc.eks.us-west-2.amazonaws.com/id/EXAMPLED539D4633E53DE1B716D3041E. If you need to remove https:// from this output value, you can include the following code in your template.",
			Type:        types.StringType,
			Computed:    true,
		},
		"resources_vpc_config": {
			// Property: ResourcesVpcConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "An object representing the VPC configuration to use for an Amazon EKS cluster.",
			//   "properties": {
			//     "EndpointPrivateAccess": {
			//       "description": "Set this value to true to enable private access for your cluster's Kubernetes API server endpoint. If you enable private access, Kubernetes API requests from within your cluster's VPC use the private VPC endpoint. The default value for this parameter is false, which disables private access for your Kubernetes API server. If you disable private access and you have nodes or AWS Fargate pods in the cluster, then ensure that publicAccessCidrs includes the necessary CIDR blocks for communication with the nodes or Fargate pods.",
			//       "type": "boolean"
			//     },
			//     "EndpointPublicAccess": {
			//       "description": "Set this value to false to disable public access to your cluster's Kubernetes API server endpoint. If you disable public access, your cluster's Kubernetes API server can only receive requests from within the cluster VPC. The default value for this parameter is true, which enables public access for your Kubernetes API server.",
			//       "type": "boolean"
			//     },
			//     "PublicAccessCidrs": {
			//       "description": "The CIDR blocks that are allowed access to your cluster's public Kubernetes API server endpoint. Communication to the endpoint from addresses outside of the CIDR blocks that you specify is denied. The default value is 0.0.0.0/0. If you've disabled private endpoint access and you have nodes or AWS Fargate pods in the cluster, then ensure that you specify the necessary CIDR blocks.",
			//       "insertionOrder": false,
			//       "items": {
			//         "minItems": 1,
			//         "type": "string"
			//       },
			//       "type": "array"
			//     },
			//     "SecurityGroupIds": {
			//       "description": "Specify one or more security groups for the cross-account elastic network interfaces that Amazon EKS creates to use to allow communication between your worker nodes and the Kubernetes control plane. If you don't specify a security group, the default security group for your VPC is used.",
			//       "insertionOrder": false,
			//       "items": {
			//         "minItems": 1,
			//         "type": "string"
			//       },
			//       "type": "array"
			//     },
			//     "SubnetIds": {
			//       "description": "Specify subnets for your Amazon EKS nodes. Amazon EKS creates cross-account elastic network interfaces in these subnets to allow communication between your nodes and the Kubernetes control plane.",
			//       "insertionOrder": false,
			//       "items": {
			//         "minItems": 1,
			//         "type": "string"
			//       },
			//       "type": "array"
			//     }
			//   },
			//   "required": [
			//     "SubnetIds"
			//   ],
			//   "type": "object"
			// }
			Description: "An object representing the VPC configuration to use for an Amazon EKS cluster.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"endpoint_private_access": {
						// Property: EndpointPrivateAccess
						Description: "Set this value to true to enable private access for your cluster's Kubernetes API server endpoint. If you enable private access, Kubernetes API requests from within your cluster's VPC use the private VPC endpoint. The default value for this parameter is false, which disables private access for your Kubernetes API server. If you disable private access and you have nodes or AWS Fargate pods in the cluster, then ensure that publicAccessCidrs includes the necessary CIDR blocks for communication with the nodes or Fargate pods.",
						Type:        types.BoolType,
						Optional:    true,
					},
					"endpoint_public_access": {
						// Property: EndpointPublicAccess
						Description: "Set this value to false to disable public access to your cluster's Kubernetes API server endpoint. If you disable public access, your cluster's Kubernetes API server can only receive requests from within the cluster VPC. The default value for this parameter is true, which enables public access for your Kubernetes API server.",
						Type:        types.BoolType,
						Optional:    true,
					},
					"public_access_cidrs": {
						// Property: PublicAccessCidrs
						Description: "The CIDR blocks that are allowed access to your cluster's public Kubernetes API server endpoint. Communication to the endpoint from addresses outside of the CIDR blocks that you specify is denied. The default value is 0.0.0.0/0. If you've disabled private endpoint access and you have nodes or AWS Fargate pods in the cluster, then ensure that you specify the necessary CIDR blocks.",
						Type:        types.ListType{ElemType: types.StringType},
						Optional:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							Multiset(),
						},
					},
					"security_group_ids": {
						// Property: SecurityGroupIds
						Description: "Specify one or more security groups for the cross-account elastic network interfaces that Amazon EKS creates to use to allow communication between your worker nodes and the Kubernetes control plane. If you don't specify a security group, the default security group for your VPC is used.",
						Type:        types.ListType{ElemType: types.StringType},
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							Multiset(),
							tfsdk.RequiresReplace(),
						},
					},
					"subnet_ids": {
						// Property: SubnetIds
						Description: "Specify subnets for your Amazon EKS nodes. Amazon EKS creates cross-account elastic network interfaces in these subnets to allow communication between your nodes and the Kubernetes control plane.",
						Type:        types.ListType{ElemType: types.StringType},
						Required:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							Multiset(),
							tfsdk.RequiresReplace(),
						},
					},
				},
			),
			Required: true,
		},
		"role_arn": {
			// Property: RoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
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
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
					},
				},
				tfsdk.SetNestedAttributesOptions{},
			),
			Optional: true,
		},
		"version": {
			// Property: Version
			// CloudFormation resource type schema:
			// {
			//   "description": "The desired Kubernetes version for your cluster. If you don't specify a value here, the latest version available in Amazon EKS is used.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The desired Kubernetes version for your cluster. If you don't specify a value here, the latest version available in Amazon EKS is used.",
			Type:        types.StringType,
			Optional:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "An object representing an Amazon EKS cluster.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EKS::Cluster").WithTerraformTypeName("awscc_eks_cluster")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                        "Arn",
		"certificate_authority_data": "CertificateAuthorityData",
		"cluster_logging":            "ClusterLogging",
		"cluster_security_group_id":  "ClusterSecurityGroupId",
		"enabled_types":              "EnabledTypes",
		"encryption_config":          "EncryptionConfig",
		"encryption_config_key_arn":  "EncryptionConfigKeyArn",
		"endpoint":                   "Endpoint",
		"endpoint_private_access":    "EndpointPrivateAccess",
		"endpoint_public_access":     "EndpointPublicAccess",
		"key":                        "Key",
		"key_arn":                    "KeyArn",
		"kubernetes_network_config":  "KubernetesNetworkConfig",
		"logging":                    "Logging",
		"name":                       "Name",
		"open_id_connect_issuer_url": "OpenIdConnectIssuerUrl",
		"provider":                   "Provider",
		"public_access_cidrs":        "PublicAccessCidrs",
		"resources":                  "Resources",
		"resources_vpc_config":       "ResourcesVpcConfig",
		"role_arn":                   "RoleArn",
		"security_group_ids":         "SecurityGroupIds",
		"service_ipv_4_cidr":         "ServiceIpv4Cidr",
		"subnet_ids":                 "SubnetIds",
		"tags":                       "Tags",
		"type":                       "Type",
		"value":                      "Value",
		"version":                    "Version",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
