// Code generated by generators/resource/main.go; DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_apigateway_method", methodResourceType)
}

// methodResourceType returns the Terraform awscc_apigateway_method resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::ApiGateway::Method resource type.
func methodResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"api_key_required": {
			// Property: ApiKeyRequired
			// CloudFormation resource type schema:
			// {
			//   "description": "Indicates whether the method requires clients to submit a valid API key.",
			//   "type": "boolean"
			// }
			Description: "Indicates whether the method requires clients to submit a valid API key.",
			Type:        types.BoolType,
			Optional:    true,
		},
		"authorization_scopes": {
			// Property: AuthorizationScopes
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of authorization scopes configured on the method.",
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "A list of authorization scopes configured on the method.",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
		},
		"authorization_type": {
			// Property: AuthorizationType
			// CloudFormation resource type schema:
			// {
			//   "description": "The method's authorization type.",
			//   "enum": [
			//     "NONE",
			//     "AWS_IAM",
			//     "CUSTOM",
			//     "COGNITO_USER_POOLS"
			//   ],
			//   "type": "string"
			// }
			Description: "The method's authorization type.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"NONE",
					"AWS_IAM",
					"CUSTOM",
					"COGNITO_USER_POOLS",
				}),
			},
		},
		"authorizer_id": {
			// Property: AuthorizerId
			// CloudFormation resource type schema:
			// {
			//   "description": "The identifier of the authorizer to use on this method.",
			//   "type": "string"
			// }
			Description: "The identifier of the authorizer to use on this method.",
			Type:        types.StringType,
			Optional:    true,
		},
		"http_method": {
			// Property: HttpMethod
			// CloudFormation resource type schema:
			// {
			//   "description": "The backend system that the method calls when it receives a request.",
			//   "type": "string"
			// }
			Description: "The backend system that the method calls when it receives a request.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"integration": {
			// Property: Integration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The backend system that the method calls when it receives a request.",
			//   "properties": {
			//     "CacheKeyParameters": {
			//       "description": "A list of request parameters whose values API Gateway caches.",
			//       "items": {
			//         "type": "string"
			//       },
			//       "type": "array",
			//       "uniqueItems": true
			//     },
			//     "CacheNamespace": {
			//       "description": "An API-specific tag group of related cached parameters.",
			//       "type": "string"
			//     },
			//     "ConnectionId": {
			//       "description": "The ID of the VpcLink used for the integration when connectionType=VPC_LINK, otherwise undefined.",
			//       "type": "string"
			//     },
			//     "ConnectionType": {
			//       "description": "The type of the network connection to the integration endpoint.",
			//       "enum": [
			//         "INTERNET",
			//         "VPC_LINK"
			//       ],
			//       "type": "string"
			//     },
			//     "ContentHandling": {
			//       "description": "Specifies how to handle request payload content type conversions.",
			//       "enum": [
			//         "CONVERT_TO_BINARY",
			//         "CONVERT_TO_TEXT"
			//       ],
			//       "type": "string"
			//     },
			//     "Credentials": {
			//       "description": "The credentials that are required for the integration.",
			//       "type": "string"
			//     },
			//     "IntegrationHttpMethod": {
			//       "description": "The integration's HTTP method type.",
			//       "type": "string"
			//     },
			//     "IntegrationResponses": {
			//       "description": "The response that API Gateway provides after a method's backend completes processing a request.",
			//       "items": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "ContentHandling": {
			//             "description": "Specifies how to handle request payload content type conversions.",
			//             "enum": [
			//               "CONVERT_TO_BINARY",
			//               "CONVERT_TO_TEXT"
			//             ],
			//             "type": "string"
			//           },
			//           "ResponseParameters": {
			//             "additionalProperties": false,
			//             "description": "The response parameters from the backend response that API Gateway sends to the method response.",
			//             "patternProperties": {
			//               "": {
			//                 "type": "string"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "ResponseTemplates": {
			//             "additionalProperties": false,
			//             "description": "The templates that are used to transform the integration response body. Specify templates as key-value pairs (string-to-string mappings), with a content type as the key and a template as the value.",
			//             "patternProperties": {
			//               "": {
			//                 "type": "string"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "SelectionPattern": {
			//             "description": "A regular expression that specifies which error strings or status codes from the backend map to the integration response.",
			//             "type": "string"
			//           },
			//           "StatusCode": {
			//             "description": "The status code that API Gateway uses to map the integration response to a MethodResponse status code.",
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "StatusCode"
			//         ],
			//         "type": "object"
			//       },
			//       "type": "array",
			//       "uniqueItems": true
			//     },
			//     "PassthroughBehavior": {
			//       "description": "Indicates when API Gateway passes requests to the targeted backend.",
			//       "enum": [
			//         "WHEN_NO_MATCH",
			//         "WHEN_NO_TEMPLATES",
			//         "NEVER"
			//       ],
			//       "type": "string"
			//     },
			//     "RequestParameters": {
			//       "additionalProperties": false,
			//       "description": "The request parameters that API Gateway sends with the backend request.",
			//       "patternProperties": {
			//         "": {
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "RequestTemplates": {
			//       "additionalProperties": false,
			//       "description": "A map of Apache Velocity templates that are applied on the request payload.",
			//       "patternProperties": {
			//         "": {
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "TimeoutInMillis": {
			//       "description": "Custom timeout between 50 and 29,000 milliseconds.",
			//       "maximum": 29000,
			//       "minimum": 50,
			//       "type": "integer"
			//     },
			//     "Type": {
			//       "description": "The type of backend that your method is running.",
			//       "enum": [
			//         "AWS",
			//         "AWS_PROXY",
			//         "HTTP",
			//         "HTTP_PROXY",
			//         "MOCK"
			//       ],
			//       "type": "string"
			//     },
			//     "Uri": {
			//       "description": "The Uniform Resource Identifier (URI) for the integration.",
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "Type"
			//   ],
			//   "type": "object"
			// }
			Description: "The backend system that the method calls when it receives a request.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"cache_key_parameters": {
						// Property: CacheKeyParameters
						Description: "A list of request parameters whose values API Gateway caches.",
						Type:        types.ListType{ElemType: types.StringType},
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.UniqueItems(),
						},
					},
					"cache_namespace": {
						// Property: CacheNamespace
						Description: "An API-specific tag group of related cached parameters.",
						Type:        types.StringType,
						Optional:    true,
					},
					"connection_id": {
						// Property: ConnectionId
						Description: "The ID of the VpcLink used for the integration when connectionType=VPC_LINK, otherwise undefined.",
						Type:        types.StringType,
						Optional:    true,
					},
					"connection_type": {
						// Property: ConnectionType
						Description: "The type of the network connection to the integration endpoint.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"INTERNET",
								"VPC_LINK",
							}),
						},
					},
					"content_handling": {
						// Property: ContentHandling
						Description: "Specifies how to handle request payload content type conversions.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"CONVERT_TO_BINARY",
								"CONVERT_TO_TEXT",
							}),
						},
					},
					"credentials": {
						// Property: Credentials
						Description: "The credentials that are required for the integration.",
						Type:        types.StringType,
						Optional:    true,
					},
					"integration_http_method": {
						// Property: IntegrationHttpMethod
						Description: "The integration's HTTP method type.",
						Type:        types.StringType,
						Optional:    true,
					},
					"integration_responses": {
						// Property: IntegrationResponses
						Description: "The response that API Gateway provides after a method's backend completes processing a request.",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"content_handling": {
									// Property: ContentHandling
									Description: "Specifies how to handle request payload content type conversions.",
									Type:        types.StringType,
									Optional:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringInSlice([]string{
											"CONVERT_TO_BINARY",
											"CONVERT_TO_TEXT",
										}),
									},
								},
								"response_parameters": {
									// Property: ResponseParameters
									Description: "The response parameters from the backend response that API Gateway sends to the method response.",
									// Pattern: ""
									Type:     types.MapType{ElemType: types.StringType},
									Optional: true,
								},
								"response_templates": {
									// Property: ResponseTemplates
									Description: "The templates that are used to transform the integration response body. Specify templates as key-value pairs (string-to-string mappings), with a content type as the key and a template as the value.",
									// Pattern: ""
									Type:     types.MapType{ElemType: types.StringType},
									Optional: true,
								},
								"selection_pattern": {
									// Property: SelectionPattern
									Description: "A regular expression that specifies which error strings or status codes from the backend map to the integration response.",
									Type:        types.StringType,
									Optional:    true,
								},
								"status_code": {
									// Property: StatusCode
									Description: "The status code that API Gateway uses to map the integration response to a MethodResponse status code.",
									Type:        types.StringType,
									Required:    true,
								},
							},
							tfsdk.ListNestedAttributesOptions{},
						),
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.UniqueItems(),
						},
					},
					"passthrough_behavior": {
						// Property: PassthroughBehavior
						Description: "Indicates when API Gateway passes requests to the targeted backend.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"WHEN_NO_MATCH",
								"WHEN_NO_TEMPLATES",
								"NEVER",
							}),
						},
					},
					"request_parameters": {
						// Property: RequestParameters
						Description: "The request parameters that API Gateway sends with the backend request.",
						// Pattern: ""
						Type:     types.MapType{ElemType: types.StringType},
						Optional: true,
					},
					"request_templates": {
						// Property: RequestTemplates
						Description: "A map of Apache Velocity templates that are applied on the request payload.",
						// Pattern: ""
						Type:     types.MapType{ElemType: types.StringType},
						Optional: true,
					},
					"timeout_in_millis": {
						// Property: TimeoutInMillis
						Description: "Custom timeout between 50 and 29,000 milliseconds.",
						Type:        types.Int64Type,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntBetween(50, 29000),
						},
					},
					"type": {
						// Property: Type
						Description: "The type of backend that your method is running.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"AWS",
								"AWS_PROXY",
								"HTTP",
								"HTTP_PROXY",
								"MOCK",
							}),
						},
					},
					"uri": {
						// Property: Uri
						Description: "The Uniform Resource Identifier (URI) for the integration.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Optional: true,
		},
		"method_responses": {
			// Property: MethodResponses
			// CloudFormation resource type schema:
			// {
			//   "description": "The responses that can be sent to the client who calls the method.",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "ResponseModels": {
			//         "additionalProperties": false,
			//         "description": "The resources used for the response's content type. Specify response models as key-value pairs (string-to-string maps), with a content type as the key and a Model resource name as the value.",
			//         "patternProperties": {
			//           "": {
			//             "type": "string"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "ResponseParameters": {
			//         "additionalProperties": false,
			//         "description": "Response parameters that API Gateway sends to the client that called a method. Specify response parameters as key-value pairs (string-to-Boolean maps), with a destination as the key and a Boolean as the value.",
			//         "patternProperties": {
			//           "": {
			//             "type": "boolean"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "StatusCode": {
			//         "description": "The method response's status code, which you map to an IntegrationResponse.",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "StatusCode"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "The responses that can be sent to the client who calls the method.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"response_models": {
						// Property: ResponseModels
						Description: "The resources used for the response's content type. Specify response models as key-value pairs (string-to-string maps), with a content type as the key and a Model resource name as the value.",
						// Pattern: ""
						Type:     types.MapType{ElemType: types.StringType},
						Optional: true,
					},
					"response_parameters": {
						// Property: ResponseParameters
						Description: "Response parameters that API Gateway sends to the client that called a method. Specify response parameters as key-value pairs (string-to-Boolean maps), with a destination as the key and a Boolean as the value.",
						// Pattern: ""
						Type:     types.MapType{ElemType: types.BoolType},
						Optional: true,
					},
					"status_code": {
						// Property: StatusCode
						Description: "The method response's status code, which you map to an IntegrationResponse.",
						Type:        types.StringType,
						Required:    true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.UniqueItems(),
			},
		},
		"operation_name": {
			// Property: OperationName
			// CloudFormation resource type schema:
			// {
			//   "description": "A friendly operation name for the method.",
			//   "type": "string"
			// }
			Description: "A friendly operation name for the method.",
			Type:        types.StringType,
			Optional:    true,
		},
		"request_models": {
			// Property: RequestModels
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The resources that are used for the request's content type. Specify request models as key-value pairs (string-to-string mapping), with a content type as the key and a Model resource name as the value.",
			//   "patternProperties": {
			//     "": {
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The resources that are used for the request's content type. Specify request models as key-value pairs (string-to-string mapping), with a content type as the key and a Model resource name as the value.",
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Optional: true,
		},
		"request_parameters": {
			// Property: RequestParameters
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The request parameters that API Gateway accepts. Specify request parameters as key-value pairs (string-to-Boolean mapping), with a source as the key and a Boolean as the value.",
			//   "patternProperties": {
			//     "": {
			//       "type": "boolean"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The request parameters that API Gateway accepts. Specify request parameters as key-value pairs (string-to-Boolean mapping), with a source as the key and a Boolean as the value.",
			// Pattern: ""
			Type:     types.MapType{ElemType: types.BoolType},
			Optional: true,
		},
		"request_validator_id": {
			// Property: RequestValidatorId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the associated request validator.",
			//   "type": "string"
			// }
			Description: "The ID of the associated request validator.",
			Type:        types.StringType,
			Optional:    true,
		},
		"resource_id": {
			// Property: ResourceId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of an API Gateway resource.",
			//   "type": "string"
			// }
			Description: "The ID of an API Gateway resource.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"rest_api_id": {
			// Property: RestApiId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the RestApi resource in which API Gateway creates the method.",
			//   "type": "string"
			// }
			Description: "The ID of the RestApi resource in which API Gateway creates the method.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
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
		Description: "Resource Type definition for AWS::ApiGateway::Method",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApiGateway::Method").WithTerraformTypeName("awscc_apigateway_method")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"api_key_required":        "ApiKeyRequired",
		"authorization_scopes":    "AuthorizationScopes",
		"authorization_type":      "AuthorizationType",
		"authorizer_id":           "AuthorizerId",
		"cache_key_parameters":    "CacheKeyParameters",
		"cache_namespace":         "CacheNamespace",
		"connection_id":           "ConnectionId",
		"connection_type":         "ConnectionType",
		"content_handling":        "ContentHandling",
		"credentials":             "Credentials",
		"http_method":             "HttpMethod",
		"integration":             "Integration",
		"integration_http_method": "IntegrationHttpMethod",
		"integration_responses":   "IntegrationResponses",
		"method_responses":        "MethodResponses",
		"operation_name":          "OperationName",
		"passthrough_behavior":    "PassthroughBehavior",
		"request_models":          "RequestModels",
		"request_parameters":      "RequestParameters",
		"request_templates":       "RequestTemplates",
		"request_validator_id":    "RequestValidatorId",
		"resource_id":             "ResourceId",
		"response_models":         "ResponseModels",
		"response_parameters":     "ResponseParameters",
		"response_templates":      "ResponseTemplates",
		"rest_api_id":             "RestApiId",
		"selection_pattern":       "SelectionPattern",
		"status_code":             "StatusCode",
		"timeout_in_millis":       "TimeoutInMillis",
		"type":                    "Type",
		"uri":                     "Uri",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
