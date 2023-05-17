// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/api/user/login": {
            "post": {
                "description": "登录",
                "tags": [
                    "用户信息"
                ],
                "summary": "登录",
                "parameters": [
                    {
                        "description": "用户信息",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/reqDto.NewUserLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "返回结果",
                        "schema": {
                            "type": "Object"
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "type": "Object"
                        }
                    }
                }
            }
        },
        "/api/user/{id}": {
            "get": {
                "description": "获取指定用户的信息",
                "tags": [
                    "用户信息"
                ],
                "summary": "获取用户信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "返回结果",
                        "schema": {
                            "type": "Object"
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "type": "Object"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "reqDto.Captcha": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                }
            }
        },
        "reqDto.NewUserLogin": {
            "type": "object",
            "required": [
                "method",
                "phone",
                "revoke"
            ],
            "properties": {
                "captcha": {
                    "$ref": "#/definitions/reqDto.Captcha"
                },
                "message": {
                    "$ref": "#/definitions/reqDto.TextMessage"
                },
                "method": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "revoke": {
                    "type": "boolean"
                }
            }
        },
        "reqDto.TextMessage": {
            "type": "object",
            "properties": {
                "code": {
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