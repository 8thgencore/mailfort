// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Tom Jerry",
            "url": "https://github.com/8thgencore/mailfort",
            "email": "test@gmail.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/send-confirmation-email": {
            "post": {
                "description": "Processes a request to generate and send an email with a confirmation code.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Mail"
                ],
                "summary": "Send Confirmation Email",
                "parameters": [
                    {
                        "description": "Request body for sending confirmation email",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.sendEmailWithOTPCodeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success response",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/send-password-reset-email": {
            "post": {
                "description": "Processes a request to generate and send an email with a password reset code.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Mail"
                ],
                "summary": "Send Password Reset Email",
                "parameters": [
                    {
                        "description": "Request body for sending password reset email",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.sendEmailWithOTPCodeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success response",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "http.sendEmailWithOTPCodeRequest": {
            "type": "object",
            "required": [
                "email",
                "otp_code"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "user@example.com"
                },
                "otp_code": {
                    "type": "string",
                    "example": "123456"
                }
            }
        },
        "response.ErrorResponse": {
            "type": "object",
            "properties": {
                "messages": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "Error message 1",
                        " Error message 2"
                    ]
                },
                "success": {
                    "type": "boolean",
                    "example": false
                }
            }
        },
        "response.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string",
                    "example": "Success"
                },
                "success": {
                    "type": "boolean",
                    "example": true
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "api.example.com",
	BasePath:         "/v1",
	Schemes:          []string{"http", "https"},
	Title:            "MailFort API",
	Description:      "MailFort API is a service for handling email-related operations.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
