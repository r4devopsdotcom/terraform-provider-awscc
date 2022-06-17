// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package connect

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_connect_user", userDataSourceType)
}

// userDataSourceType returns the Terraform awscc_connect_user data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Connect::User resource type.
func userDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"directory_user_id": {
			// Property: DirectoryUserId
			// CloudFormation resource type schema:
			// {
			//   "description": "The identifier of the user account in the directory used for identity management.",
			//   "type": "string"
			// }
			Description: "The identifier of the user account in the directory used for identity management.",
			Type:        types.StringType,
			Computed:    true,
		},
		"hierarchy_group_arn": {
			// Property: HierarchyGroupArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The identifier of the hierarchy group for the user.",
			//   "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/agent-group/[-a-zA-Z0-9]*$",
			//   "type": "string"
			// }
			Description: "The identifier of the hierarchy group for the user.",
			Type:        types.StringType,
			Computed:    true,
		},
		"identity_info": {
			// Property: IdentityInfo
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The information about the identity of the user.",
			//   "properties": {
			//     "Email": {
			//       "description": "The email address. If you are using SAML for identity management and include this parameter, an error is returned.",
			//       "type": "string"
			//     },
			//     "FirstName": {
			//       "description": "The first name. This is required if you are using Amazon Connect or SAML for identity management.",
			//       "type": "string"
			//     },
			//     "LastName": {
			//       "description": "The last name. This is required if you are using Amazon Connect or SAML for identity management.",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The information about the identity of the user.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"email": {
						// Property: Email
						Description: "The email address. If you are using SAML for identity management and include this parameter, an error is returned.",
						Type:        types.StringType,
						Computed:    true,
					},
					"first_name": {
						// Property: FirstName
						Description: "The first name. This is required if you are using Amazon Connect or SAML for identity management.",
						Type:        types.StringType,
						Computed:    true,
					},
					"last_name": {
						// Property: LastName
						Description: "The last name. This is required if you are using Amazon Connect or SAML for identity management.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"instance_arn": {
			// Property: InstanceArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The identifier of the Amazon Connect instance.",
			//   "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$",
			//   "type": "string"
			// }
			Description: "The identifier of the Amazon Connect instance.",
			Type:        types.StringType,
			Computed:    true,
		},
		"password": {
			// Property: Password
			// CloudFormation resource type schema:
			// {
			//   "description": "The password for the user account. A password is required if you are using Amazon Connect for identity management. Otherwise, it is an error to include a password.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The password for the user account. A password is required if you are using Amazon Connect for identity management. Otherwise, it is an error to include a password.",
			Type:        types.StringType,
			Computed:    true,
		},
		"phone_config": {
			// Property: PhoneConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The phone settings for the user.",
			//   "properties": {
			//     "AfterContactWorkTimeLimit": {
			//       "description": "The After Call Work (ACW) timeout setting, in seconds.",
			//       "minimum": 0,
			//       "type": "integer"
			//     },
			//     "AutoAccept": {
			//       "description": "The Auto accept setting.",
			//       "type": "boolean"
			//     },
			//     "DeskPhoneNumber": {
			//       "description": "The phone number for the user's desk phone.",
			//       "type": "string"
			//     },
			//     "PhoneType": {
			//       "description": "The phone type.",
			//       "enum": [
			//         "SOFT_PHONE",
			//         "DESK_PHONE"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "PhoneType"
			//   ],
			//   "type": "object"
			// }
			Description: "The phone settings for the user.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"after_contact_work_time_limit": {
						// Property: AfterContactWorkTimeLimit
						Description: "The After Call Work (ACW) timeout setting, in seconds.",
						Type:        types.Int64Type,
						Computed:    true,
					},
					"auto_accept": {
						// Property: AutoAccept
						Description: "The Auto accept setting.",
						Type:        types.BoolType,
						Computed:    true,
					},
					"desk_phone_number": {
						// Property: DeskPhoneNumber
						Description: "The phone number for the user's desk phone.",
						Type:        types.StringType,
						Computed:    true,
					},
					"phone_type": {
						// Property: PhoneType
						Description: "The phone type.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"routing_profile_arn": {
			// Property: RoutingProfileArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The identifier of the routing profile for the user.",
			//   "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/routing-profile/[-a-zA-Z0-9]*$",
			//   "type": "string"
			// }
			Description: "The identifier of the routing profile for the user.",
			Type:        types.StringType,
			Computed:    true,
		},
		"security_profile_arns": {
			// Property: SecurityProfileArns
			// CloudFormation resource type schema:
			// {
			//   "description": "One or more security profile arns for the user",
			//   "insertionOrder": false,
			//   "items": {
			//     "description": "The identifier of the security profile for the user.",
			//     "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/security-profile/[-a-zA-Z0-9]*$",
			//     "type": "string"
			//   },
			//   "maxItems": 10,
			//   "minItems": 1,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "One or more security profile arns for the user",
			Type:        types.SetType{ElemType: types.StringType},
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "One or more tags.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is maximum of 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
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
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "One or more tags.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is maximum of 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"user_arn": {
			// Property: UserArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) for the user.",
			//   "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/agent/[-a-zA-Z0-9]*$",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) for the user.",
			Type:        types.StringType,
			Computed:    true,
		},
		"username": {
			// Property: Username
			// CloudFormation resource type schema:
			// {
			//   "description": "The user name for the account.",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "pattern": "[a-zA-Z0-9\\_\\-\\.\\@]+",
			//   "type": "string"
			// }
			Description: "The user name for the account.",
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
		Description: "Data Source schema for AWS::Connect::User",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Connect::User").WithTerraformTypeName("awscc_connect_user")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"after_contact_work_time_limit": "AfterContactWorkTimeLimit",
		"auto_accept":                   "AutoAccept",
		"desk_phone_number":             "DeskPhoneNumber",
		"directory_user_id":             "DirectoryUserId",
		"email":                         "Email",
		"first_name":                    "FirstName",
		"hierarchy_group_arn":           "HierarchyGroupArn",
		"identity_info":                 "IdentityInfo",
		"instance_arn":                  "InstanceArn",
		"key":                           "Key",
		"last_name":                     "LastName",
		"password":                      "Password",
		"phone_config":                  "PhoneConfig",
		"phone_type":                    "PhoneType",
		"routing_profile_arn":           "RoutingProfileArn",
		"security_profile_arns":         "SecurityProfileArns",
		"tags":                          "Tags",
		"user_arn":                      "UserArn",
		"username":                      "Username",
		"value":                         "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
