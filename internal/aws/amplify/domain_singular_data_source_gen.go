// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package amplify

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_amplify_domain", domainDataSourceType)
}

// domainDataSourceType returns the Terraform awscc_amplify_domain data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Amplify::Domain resource type.
func domainDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"app_id": {
			// Property: AppId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 20,
			//   "minLength": 1,
			//   "pattern": "d[a-z0-9]+",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1000,
			//   "pattern": "(?s).*",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"auto_sub_domain_creation_patterns": {
			// Property: AutoSubDomainCreationPatterns
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "maxLength": 2048,
			//     "minLength": 1,
			//     "pattern": "(?s).+",
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Computed: true,
		},
		"auto_sub_domain_iam_role": {
			// Property: AutoSubDomainIAMRole
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1000,
			//   "pattern": "^$|^arn:.+:iam::\\d{12}:role.+",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"certificate_record": {
			// Property: CertificateRecord
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1000,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"domain_name": {
			// Property: DomainName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 255,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"domain_status": {
			// Property: DomainStatus
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"enable_auto_sub_domain": {
			// Property: EnableAutoSubDomain
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Computed: true,
		},
		"status_reason": {
			// Property: StatusReason
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1000,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"sub_domain_settings": {
			// Property: SubDomainSettings
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "BranchName": {
			//         "maxLength": 255,
			//         "minLength": 1,
			//         "pattern": "(?s).+",
			//         "type": "string"
			//       },
			//       "Prefix": {
			//         "maxLength": 255,
			//         "pattern": "(?s).*",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Prefix",
			//       "BranchName"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 255,
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"branch_name": {
						// Property: BranchName
						Type:     types.StringType,
						Computed: true,
					},
					"prefix": {
						// Property: Prefix
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
		Description: "Data Source schema for AWS::Amplify::Domain",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Amplify::Domain").WithTerraformTypeName("awscc_amplify_domain")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"app_id":                            "AppId",
		"arn":                               "Arn",
		"auto_sub_domain_creation_patterns": "AutoSubDomainCreationPatterns",
		"auto_sub_domain_iam_role":          "AutoSubDomainIAMRole",
		"branch_name":                       "BranchName",
		"certificate_record":                "CertificateRecord",
		"domain_name":                       "DomainName",
		"domain_status":                     "DomainStatus",
		"enable_auto_sub_domain":            "EnableAutoSubDomain",
		"prefix":                            "Prefix",
		"status_reason":                     "StatusReason",
		"sub_domain_settings":               "SubDomainSettings",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
