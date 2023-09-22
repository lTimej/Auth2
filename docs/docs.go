// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/applications": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取服务商应用",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "服务商相关接口"
                ],
                "summary": "获取服务商应用",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httpResp.Resp"
                        }
                    }
                }
            }
        },
        "/auth/applications/credentials": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取生成的应用凭证",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "服务商相关接口"
                ],
                "summary": "获取生成的应用凭证",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httpResp.Resp"
                        }
                    }
                }
            }
        },
        "/auth/applications/register": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "服务商应用注册",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "服务商相关接口"
                ],
                "summary": "服务商应用注册",
                "parameters": [
                    {
                        "description": "服务商应用注册参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.ProviderApplicationRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httpResp.Resp"
                        }
                    }
                }
            }
        },
        "/auth/odic_providers/register": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "服务商注册接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "服务商相关接口"
                ],
                "summary": "服务商注册接口",
                "parameters": [
                    {
                        "description": "服务商注册参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.OpenIDConfigRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httpResp.Resp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "httpResp.Resp": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {
                    "type": "string"
                }
            }
        },
        "service.OpenIDConfigRequest": {
            "type": "object",
            "required": [
                "client_id",
                "client_secret",
                "issuer"
            ],
            "properties": {
                "client_id": {
                    "type": "string"
                },
                "client_secret": {
                    "type": "string"
                },
                "issuer": {
                    "type": "string"
                },
                "sched_provider_code": {
                    "type": "string"
                }
            }
        },
        "service.ProviderApplicationRequest": {
            "type": "object",
            "properties": {
                "client_id": {
                    "type": "string"
                },
                "client_secret": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "provider_code": {
                    "type": "string"
                },
                "redirect_uri": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}