// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/user/info": {
            "get": {
                "description": "用户个人信息接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "用户个人信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "0": {
                        "description": "msg: Success",
                        "schema": {
                            "$ref": "#/definitions/serializer.User"
                        }
                    },
                    "50000": {
                        "description": "msg: 查询个人信息错误",
                        "schema": {
                            "$ref": "#/definitions/serializer.Response"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "用户登录接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "用户名, 密码",
                        "name": "form",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.UserLoginService"
                        }
                    }
                ],
                "responses": {
                    "0": {
                        "description": "msg: Success",
                        "schema": {
                            "$ref": "#/definitions/serializer.Response"
                        }
                    },
                    "30001": {
                        "description": "msg: 参数错误",
                        "schema": {
                            "$ref": "#/definitions/serializer.Response"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "用户注册接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "description": "用户名, 密码, 邮箱, 昵称",
                        "name": "form",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.UserRegisterService"
                        }
                    }
                ],
                "responses": {
                    "0": {
                        "description": "msg: Success",
                        "schema": {
                            "$ref": "#/definitions/serializer.Response"
                        }
                    },
                    "30001": {
                        "description": "msg: 参数错误",
                        "schema": {
                            "$ref": "#/definitions/serializer.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "serializer.Response": {
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
        "serializer.User": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "state": {
                    "type": "integer"
                },
                "uid": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "service.UserLoginService": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "service.UserRegisterService": {
            "type": "object",
            "required": [
                "email",
                "nickname",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.2.1",
	Host:        "127.0.0.1:3000",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "ACG.Fate API",
	Description: "The RESTFul API of Server",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}
