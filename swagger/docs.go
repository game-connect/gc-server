// Package swagger GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package swagger

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/user_login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "ユーザーログイン",
                "parameters": [
                    {
                        "description": "ユーザーログイン",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/parameter.UserLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Success"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Items": {
                                            "$ref": "#/definitions/output.UserLogin"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/output.Error"
                            }
                        }
                    }
                }
            }
        },
        "/auth/user_register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "ユーザー登録",
                "parameters": [
                    {
                        "description": "ユーザー登録",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/parameter.UserRegister"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Success"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Items": {
                                            "$ref": "#/definitions/output.UserRegister"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/output.Error"
                            }
                        }
                    }
                }
            }
        },
        "/user/{user_key}/user_check": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "ユーザー取得",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "maxLength": 12,
                        "type": "string",
                        "description": "user_key",
                        "name": "user_key",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Success"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Items": {
                                            "$ref": "#/definitions/output.UserCheck"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/output.Error"
                            }
                        }
                    }
                }
            }
        },
        "/user/{user_key}/user_delete": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "ユーザー削除",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ユーザーキー",
                        "name": "user_key",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Success"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Items": {
                                            "$ref": "#/definitions/output.UserDelete"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/output.Error"
                            }
                        }
                    }
                }
            }
        },
        "/user/{user_key}/user_logout": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "ログアウト",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "maxLength": 12,
                        "type": "string",
                        "description": "user_key",
                        "name": "user_key",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Success"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Items": {
                                            "$ref": "#/definitions/output.UserLogout"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/output.Error"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "output.Error": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "output.UserCheck": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "user_key": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "output.UserDelete": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "output.UserLogin": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                },
                "user_key": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "output.UserLogout": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "output.UserRegister": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "user_key": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "parameter.UserLogin": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "parameter.UserRegister": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "response.Success": {
            "type": "object",
            "properties": {
                "items": {},
                "status": {
                    "type": "integer"
                },
                "types": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8001",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Chat Connect",
	Description:      "This is a sample swagger server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
