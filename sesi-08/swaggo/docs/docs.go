// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "soberkoder@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/orders": {
            "get": {
                "description": "Get details of all orders",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Get details of all orders",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.Order"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new order with theb input paylod",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Create a new order",
                "parameters": [
                    {
                        "description": "Create order",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Order"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Order"
                        }
                    }
                }
            }
        },
        "/orders/{orderId}": {
            "get": {
                "description": "Get details of order by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Get details of order by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "orderId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Order"
                        }
                    }
                }
            },
            "put": {
                "description": "Update data order where orderId",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Update data order where orderId",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "orderId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Order"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete data order where orderId",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Delete data order where orderId",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "orderId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        }
    },
    "definitions": {
        "main.Item": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string",
                    "example": "A random description"
                },
                "itemId": {
                    "type": "string",
                    "example": "A1B2C3"
                },
                "quantity": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "main.Order": {
            "type": "object",
            "properties": {
                "customerName": {
                    "type": "string",
                    "example": "Leo Messi"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.Item"
                    }
                },
                "orderId": {
                    "type": "string",
                    "example": "1"
                },
                "orderedAt": {
                    "type": "string",
                    "example": "2019-11-09T21:21:46+00:00"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
