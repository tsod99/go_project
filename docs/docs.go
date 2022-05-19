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
        "/add/group": {
            "post": {
                "description": "create new group",
                "consumes": [
                    "application/json"
                ],
                "summary": "Create new group",
                "parameters": [
                    {
                        "description": "group data",
                        "name": "group",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.Group"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": ""
                    },
                    "400": {
                        "description": ""
                    },
                    "405": {
                        "description": ""
                    },
                    "409": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/add/user": {
            "post": {
                "description": "create new user",
                "consumes": [
                    "application/json"
                ],
                "summary": "Create new user",
                "parameters": [
                    {
                        "description": "user data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": ""
                    },
                    "400": {
                        "description": ""
                    },
                    "405": {
                        "description": ""
                    },
                    "409": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/delete/group": {
            "delete": {
                "description": "delete group",
                "consumes": [
                    "application/json"
                ],
                "summary": "Delete Group",
                "parameters": [
                    {
                        "description": "delete group request",
                        "name": "delete",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.DeleteGroupRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": ""
                    },
                    "405": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/delete/user": {
            "delete": {
                "description": "delete user",
                "consumes": [
                    "application/json"
                ],
                "summary": "Delete User",
                "parameters": [
                    {
                        "description": "delete user request",
                        "name": "delete",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.DeleteUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": ""
                    },
                    "405": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/list/groups": {
            "get": {
                "description": "list groups",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "List groups",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api.Group"
                            }
                        }
                    },
                    "405": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/list/users": {
            "get": {
                "description": "list users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "List users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api.User"
                            }
                        }
                    },
                    "405": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/update/group": {
            "patch": {
                "description": "update group",
                "consumes": [
                    "application/json"
                ],
                "summary": "Update group",
                "parameters": [
                    {
                        "description": "update group request",
                        "name": "update",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.UpdateGroupRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": ""
                    },
                    "405": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/update/user": {
            "patch": {
                "description": "update user",
                "consumes": [
                    "application/json"
                ],
                "summary": "Update user",
                "parameters": [
                    {
                        "description": "update user request",
                        "name": "update",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.UpdateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": ""
                    },
                    "405": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "api.DeleteGroupRequest": {
            "type": "object",
            "properties": {
                "group_name": {
                    "type": "string"
                }
            }
        },
        "api.DeleteUserRequest": {
            "type": "object",
            "properties": {
                "username": {
                    "type": "string"
                }
            }
        },
        "api.Group": {
            "type": "object",
            "properties": {
                "group_name": {
                    "type": "string"
                }
            }
        },
        "api.UpdateGroupRequest": {
            "type": "object",
            "properties": {
                "group_name": {
                    "type": "string"
                },
                "new_group_name": {
                    "type": "string"
                }
            }
        },
        "api.UpdateUserRequest": {
            "type": "object",
            "properties": {
                "field": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "api.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "group": {
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
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "api",
	Description: "api swagger documentation",
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
