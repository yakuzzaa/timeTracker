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
        "/users": {
            "post": {
                "description": "Создает нового пользователя на основе переданных данных в теле запроса",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Создание пользователя",
                "parameters": [
                    {
                        "description": "Запрос",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/serializers.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/serializers.CreateUserResponse"
                        }
                    }
                }
            }
        },
        "/users/info": {
            "get": {
                "description": "Возвращает список пользователей с возможностью фильтрации и пагинации",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Получение списка пользователей",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id пользователя",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Серия паспорта",
                        "name": "passportSeries",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Номер паспорта",
                        "name": "passportNumber",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Имя",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Фамилия",
                        "name": "surname",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Отчество",
                        "name": "patronymic",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Адрес",
                        "name": "address",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Страница",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Кол-во записей на странице",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializers.GetUsersResponse"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "put": {
                "description": "Обновляет информацию о пользователе с указанным идентификатором",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Обновление данных пользователя",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id пользователя",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Запрос",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/serializers.UpdateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializers.UpdateUserResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Удаляет пользователя с указанным идентификатором",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Удаление пользователя",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id пользователя",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializers.DeleteUserResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "serializers.CreateUserRequest": {
            "type": "object",
            "required": [
                "passport"
            ],
            "properties": {
                "passport": {
                    "type": "string"
                }
            }
        },
        "serializers.CreateUserResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "serializers.DeleteUserResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        },
        "serializers.GetUsersResponse": {
            "type": "object",
            "properties": {
                "info": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/serializers.User"
                    }
                }
            }
        },
        "serializers.UpdateUserRequest": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "passportNumber": {
                    "type": "string"
                },
                "passportSeries": {
                    "type": "string"
                },
                "patronymic": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        },
        "serializers.UpdateUserResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        },
        "serializers.User": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "passportNumber": {
                    "type": "string"
                },
                "passportSeries": {
                    "type": "string"
                },
                "patronymic": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "TimeTracker API",
	Description:      "API for time tracking app",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
