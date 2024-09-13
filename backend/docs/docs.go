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
        "/api/ping": {
            "get": {
                "description": "Возвращает \"ok\" с кодом состояния 200",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health"
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/tenders": {
            "get": {
                "description": "Возвращает список всех тендеров из базы данных",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tenders"
                ],
                "summary": "Получение списка тендеров",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Tender"
                            }
                        }
                    },
                    "500": {
                        "description": "Error fetching data",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/tenders/new": {
            "post": {
                "description": "Создает новый тендер и возвращает его",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tenders"
                ],
                "summary": "Создание нового тендера",
                "parameters": [
                    {
                        "description": "Тело запроса для создания тендера",
                        "name": "tender",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Tender"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешное создание тендера",
                        "schema": {
                            "$ref": "#/definitions/models.Tender"
                        }
                    }
                }
            }
        },
        "/api/tenders/{tenderId}/edit": {
            "patch": {
                "description": "Обновляет параметры существующего тендера и возвращает обновленный тендер",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tenders"
                ],
                "summary": "Обновление существующего тендера",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID тендера",
                        "name": "tenderId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Тело запроса для обновления тендера",
                        "name": "tender",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/TenderHandlers.UpdateTenderRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешное обновление тендера",
                        "schema": {
                            "$ref": "#/definitions/models.Tender"
                        }
                    },
                    "400": {
                        "description": "Неверный запрос",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Тендер не найден",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/tenders/{tenderId}/rollback/{version}": {
            "put": {
                "description": "Откатывает параметры тендера к указанной версии",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tenders"
                ],
                "summary": "Откат тендера",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID тендера",
                        "name": "tenderId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Версия тендера",
                        "name": "version",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешный откат тендера",
                        "schema": {
                            "$ref": "#/definitions/models.Tender"
                        }
                    },
                    "400": {
                        "description": "Неверный запрос",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Тендер или версия не найдены",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "TenderHandlers.UpdateTenderRequest": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.Tender": {
            "description": "Tender содержит информацию о тендере",
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "description": "ID пользователя",
                    "type": "string"
                },
                "creator_username": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "organization_id": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/models.TenderStatus"
                },
                "tender_version": {
                    "$ref": "#/definitions/models.TenderVersion"
                },
                "updated_at": {
                    "type": "string"
                },
                "version": {
                    "type": "integer"
                }
            }
        },
        "models.TenderStatus": {
            "description": "TenderStatus содержит возможные статусы тендера",
            "type": "string",
            "enum": [
                "CREATED",
                "PUBLISHED",
                "CLOSED"
            ],
            "x-enum-varnames": [
                "TenderCreated",
                "TenderPublished",
                "TenderClosed"
            ]
        },
        "models.TenderVersion": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "tender_id": {
                    "type": "string"
                },
                "version": {
                    "type": "integer"
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
	Title:            "TenderAPI",
	Description:      "API Server for TodoList Application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
