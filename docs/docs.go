// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-06-07 22:28:45.519557 +0200 CEST m=+0.149596895

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "API for twitch logs",
        "title": "justlog API",
        "contact": {
            "name": "gempir",
            "url": "https://gempir.com",
            "email": "gempir.dev@gmail.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://github.com/gempir/justlog/blob/master/LICENSE"
        },
        "version": "1.0"
    },
    "host": "api.gempir.com",
    "basePath": "/",
    "paths": {
        "/channel/{channel}/user/{username}": {
            "get": {
                "produces": [
                    "application/json",
                    "text/plain"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Redirect to last logs of user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "channelname",
                        "name": "channel",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "username",
                        "name": "username",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "unix timestamp, limit logs by timestamps from this point",
                        "name": "from",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "unix timestamp, limit logs by timestamps to this point",
                        "name": "to",
                        "in": "query"
                    },
                    {
                        "type": "any",
                        "description": "response as json",
                        "name": "json",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "define response type only json supported currently, rest defaults to plain",
                        "name": "type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {},
                    "303": {},
                    "404": {},
                    "500": {}
                }
            }
        },
        "/channel/{channel}/user/{username}/random": {
            "get": {
                "produces": [
                    "application/json",
                    "text/plain"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get a random chat message from a user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "channelname",
                        "name": "channel",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "username",
                        "name": "username",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "any",
                        "description": "response as json",
                        "name": "json",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "define response type only json supported currently, rest defaults to plain",
                        "name": "type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.RandomQuoteJSON"
                        }
                    }
                }
            }
        },
        "/channel/{channel}/user/{username}/{year}/{month}": {
            "get": {
                "produces": [
                    "application/json",
                    "text/plain"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get logs for user by year and month",
                "parameters": [
                    {
                        "type": "string",
                        "description": "channelname",
                        "name": "channel",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "username",
                        "name": "username",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "year of logs",
                        "name": "year",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "month of logs",
                        "name": "month",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "unix timestamp, limit logs by timestamps from this point",
                        "name": "from",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "unix timestamp, limit logs by timestamps to this point",
                        "name": "to",
                        "in": "query"
                    },
                    {
                        "type": "any",
                        "description": "response as json",
                        "name": "json",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "define response type only json supported currently, rest defaults to plain",
                        "name": "type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {},
                    "500": {}
                }
            }
        },
        "/channelid/{channelid}/user/{userid}": {
            "get": {
                "produces": [
                    "application/json",
                    "text/plain"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Redirect to last logs of user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "twitch userid",
                        "name": "channelid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "twitch userid",
                        "name": "userid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "unix timestamp, limit logs by timestamps from this point",
                        "name": "from",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "unix timestamp, limit logs by timestamps to this point",
                        "name": "to",
                        "in": "query"
                    },
                    {
                        "type": "any",
                        "description": "response as json",
                        "name": "json",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "define response type only json supported currently, rest defaults to plain",
                        "name": "type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {},
                    "303": {},
                    "404": {}
                }
            }
        },
        "/channelid/{channelid}/userid/{userid}/random": {
            "get": {
                "produces": [
                    "application/json",
                    "text/plain"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get a random chat message from a user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "twitch userid",
                        "name": "channelid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "twitch userid",
                        "name": "userid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "any",
                        "description": "response as json",
                        "name": "json",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "define response type only json supported currently, rest defaults to plain",
                        "name": "type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.RandomQuoteJSON"
                        }
                    }
                }
            }
        },
        "/channelid/{channelid}/userid/{userid}/{year}/{month}": {
            "get": {
                "produces": [
                    "application/json",
                    "text/plain"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get logs for user by year and month",
                "parameters": [
                    {
                        "type": "string",
                        "description": "twitch userid",
                        "name": "channelid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "twitch userid",
                        "name": "userid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "year of logs",
                        "name": "year",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "month of logs",
                        "name": "month",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "unix timestamp, limit logs by timestamps from this point",
                        "name": "from",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "unix timestamp, limit logs by timestamps to this point",
                        "name": "to",
                        "in": "query"
                    },
                    {
                        "type": "any",
                        "description": "response as json",
                        "name": "json",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "define response type only json supported currently, rest defaults to plain",
                        "name": "type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {},
                    "500": {}
                }
            }
        },
        "/channels": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bot"
                ],
                "summary": "Get all joined channels",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.RandomQuoteJSON"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "api.RandomQuoteJSON": {
            "type": "object",
            "properties": {
                "channel": {
                    "type": "string"
                },
                "displayName": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "object",
                    "$ref": "#/definitions/api.timestamp"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "api.timestamp": {
            "type": "object"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
