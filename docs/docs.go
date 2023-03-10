// Code generated by swaggo/swag. DO NOT EDIT
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
        "/beers": {
            "get": {
                "description": "Get a list of all beers",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Beers"
                ],
                "summary": "Get a list of all beers",
                "responses": {
                    "200": {
                        "description": "Successful operation",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/controllers.BeerResponse"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controllers.BeerErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new beer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Beers"
                ],
                "summary": "Create a new beer",
                "parameters": [
                    {
                        "description": "Beer input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.BeerInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful operation",
                        "schema": {
                            "$ref": "#/definitions/controllers.BeerResponse"
                        }
                    },
                    "400": {
                        "description": "Ensure input is correct!",
                        "schema": {
                            "$ref": "#/definitions/controllers.BeerErrorResponse"
                        }
                    }
                }
            }
        },
        "/beers/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Beers"
                ],
                "summary": "Get a beer by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Beer ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.BeerResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controllers.BeerErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Beers"
                ],
                "summary": "Delete a beer by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Beer ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.DeletedBeerResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.BeerErrorResponse"
                        }
                    }
                }
            },
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Beers"
                ],
                "summary": "Update a beer by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Beer ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Beer input payload",
                        "name": "beer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.BeerInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.BeerResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.BeerErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controllers.BeerErrorResponse"
                        }
                    }
                }
            }
        },
        "/breweries": {
            "get": {
                "description": "Get a list of all breweries",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Breweries"
                ],
                "summary": "Get a list of all breweries",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/controllers.BreweryResponse"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controllers.BreweryErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Breweries"
                ],
                "summary": "Create a brewery",
                "parameters": [
                    {
                        "description": "Brewery input payload",
                        "name": "brewery",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.BreweryInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.BreweryResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.BreweryErrorResponse"
                        }
                    }
                }
            }
        },
        "/breweries/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Breweries"
                ],
                "summary": "Get a single brewery by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Brewery ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.BreweryResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controllers.BreweryErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a brewery only if there are no beers associated with it",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Breweries"
                ],
                "summary": "Delete a brewery",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Brewery ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.DeletedBreweryResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.BreweryErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controllers.BreweryErrorResponse"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update a brewery by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Breweries"
                ],
                "summary": "Update a brewery",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Brewery ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Brewery Payload",
                        "name": "brewery",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.BreweryInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.BreweryResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.BreweryErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controllers.BreweryErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.BeerErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "controllers.BeerInput": {
            "type": "object",
            "required": [
                "beername",
                "brewery"
            ],
            "properties": {
                "beername": {
                    "type": "string"
                },
                "brewery": {
                    "$ref": "#/definitions/controllers.BreweryInput"
                }
            }
        },
        "controllers.BeerResponse": {
            "type": "object",
            "properties": {
                "brewery": {
                    "$ref": "#/definitions/controllers.BreweryInput"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "controllers.BreweryErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "controllers.BreweryInput": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "controllers.BreweryResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "controllers.DeletedBeerResponse": {
            "type": "object",
            "properties": {
                "deleted": {
                    "type": "boolean"
                }
            }
        },
        "controllers.DeletedBreweryResponse": {
            "type": "object",
            "properties": {
                "deleted": {
                    "type": "boolean"
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
