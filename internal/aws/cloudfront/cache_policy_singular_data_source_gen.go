// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_cloudfront_cache_policy", cachePolicyDataSourceType)
}

// cachePolicyDataSourceType returns the Terraform awscc_cloudfront_cache_policy data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::CloudFront::CachePolicy resource type.
func cachePolicyDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"cache_policy_config": {
			// Property: CachePolicyConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "Comment": {
			//       "type": "string"
			//     },
			//     "DefaultTTL": {
			//       "minimum": 0,
			//       "type": "number"
			//     },
			//     "MaxTTL": {
			//       "minimum": 0,
			//       "type": "number"
			//     },
			//     "MinTTL": {
			//       "minimum": 0,
			//       "type": "number"
			//     },
			//     "Name": {
			//       "type": "string"
			//     },
			//     "ParametersInCacheKeyAndForwardedToOrigin": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "CookiesConfig": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "CookieBehavior": {
			//               "pattern": "",
			//               "type": "string"
			//             },
			//             "Cookies": {
			//               "items": {
			//                 "type": "string"
			//               },
			//               "type": "array",
			//               "uniqueItems": false
			//             }
			//           },
			//           "required": [
			//             "CookieBehavior"
			//           ],
			//           "type": "object"
			//         },
			//         "EnableAcceptEncodingBrotli": {
			//           "type": "boolean"
			//         },
			//         "EnableAcceptEncodingGzip": {
			//           "type": "boolean"
			//         },
			//         "HeadersConfig": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "HeaderBehavior": {
			//               "pattern": "",
			//               "type": "string"
			//             },
			//             "Headers": {
			//               "items": {
			//                 "type": "string"
			//               },
			//               "type": "array",
			//               "uniqueItems": false
			//             }
			//           },
			//           "required": [
			//             "HeaderBehavior"
			//           ],
			//           "type": "object"
			//         },
			//         "QueryStringsConfig": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "QueryStringBehavior": {
			//               "pattern": "",
			//               "type": "string"
			//             },
			//             "QueryStrings": {
			//               "items": {
			//                 "type": "string"
			//               },
			//               "type": "array",
			//               "uniqueItems": false
			//             }
			//           },
			//           "required": [
			//             "QueryStringBehavior"
			//           ],
			//           "type": "object"
			//         }
			//       },
			//       "required": [
			//         "EnableAcceptEncodingGzip",
			//         "HeadersConfig",
			//         "CookiesConfig",
			//         "QueryStringsConfig"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "required": [
			//     "Name",
			//     "MinTTL",
			//     "MaxTTL",
			//     "DefaultTTL",
			//     "ParametersInCacheKeyAndForwardedToOrigin"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"comment": {
						// Property: Comment
						Type:     types.StringType,
						Computed: true,
					},
					"default_ttl": {
						// Property: DefaultTTL
						Type:     types.Float64Type,
						Computed: true,
					},
					"max_ttl": {
						// Property: MaxTTL
						Type:     types.Float64Type,
						Computed: true,
					},
					"min_ttl": {
						// Property: MinTTL
						Type:     types.Float64Type,
						Computed: true,
					},
					"name": {
						// Property: Name
						Type:     types.StringType,
						Computed: true,
					},
					"parameters_in_cache_key_and_forwarded_to_origin": {
						// Property: ParametersInCacheKeyAndForwardedToOrigin
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"cookies_config": {
									// Property: CookiesConfig
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"cookie_behavior": {
												// Property: CookieBehavior
												Type:     types.StringType,
												Computed: true,
											},
											"cookies": {
												// Property: Cookies
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"enable_accept_encoding_brotli": {
									// Property: EnableAcceptEncodingBrotli
									Type:     types.BoolType,
									Computed: true,
								},
								"enable_accept_encoding_gzip": {
									// Property: EnableAcceptEncodingGzip
									Type:     types.BoolType,
									Computed: true,
								},
								"headers_config": {
									// Property: HeadersConfig
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"header_behavior": {
												// Property: HeaderBehavior
												Type:     types.StringType,
												Computed: true,
											},
											"headers": {
												// Property: Headers
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"query_strings_config": {
									// Property: QueryStringsConfig
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"query_string_behavior": {
												// Property: QueryStringBehavior
												Type:     types.StringType,
												Computed: true,
											},
											"query_strings": {
												// Property: QueryStrings
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
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"last_modified_time": {
			// Property: LastModifiedTime
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
		Description: "Data Source schema for AWS::CloudFront::CachePolicy",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFront::CachePolicy").WithTerraformTypeName("awscc_cloudfront_cache_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"cache_policy_config":           "CachePolicyConfig",
		"comment":                       "Comment",
		"cookie_behavior":               "CookieBehavior",
		"cookies":                       "Cookies",
		"cookies_config":                "CookiesConfig",
		"default_ttl":                   "DefaultTTL",
		"enable_accept_encoding_brotli": "EnableAcceptEncodingBrotli",
		"enable_accept_encoding_gzip":   "EnableAcceptEncodingGzip",
		"header_behavior":               "HeaderBehavior",
		"headers":                       "Headers",
		"headers_config":                "HeadersConfig",
		"id":                            "Id",
		"last_modified_time":            "LastModifiedTime",
		"max_ttl":                       "MaxTTL",
		"min_ttl":                       "MinTTL",
		"name":                          "Name",
		"parameters_in_cache_key_and_forwarded_to_origin": "ParametersInCacheKeyAndForwardedToOrigin",
		"query_string_behavior":                           "QueryStringBehavior",
		"query_strings":                                   "QueryStrings",
		"query_strings_config":                            "QueryStringsConfig",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
