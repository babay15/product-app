package docs

import (
	"bytes"
	"encoding/json"
	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
	"strings"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger":"2.0",
   "info":{
      "description":"{{.Description}}",
      "title":"{{.Title}}",
      "contact":{
         
      },
      "license":{
         
      },
      "version":"{{.Version}}"
   },
   "host":"{{.Host}}",
   "basePath":"{{.BasePath}}",
   "paths":{
      "/auth/token":{
         "post":{
            "description":"Authenticates a user and provides a JWT to Authorize API calls",
            "produces":[
               "application/json"
            ],
            "tags": [
               "Authorization"
            ],
            "summary":"Provides a JSON Web Token",
            "operationId":"Authentication",
            "parameters":[
               {
                  "type":"string",
                  "description":"Username",
                  "name":"username",
                  "in":"formData",
                  "required":true
               },
               {
                  "type":"string",
                  "description":"Password",
                  "name":"password",
                  "in":"formData",
                  "required":true
               }
            ],
            "responses":{
               "200":{
                  "description":"OK",
                  "schema":{
                     "$ref":"#/definitions/dto.JWT"
                  }
               },
               "401":{
                  "description":"Unauthorized",
                  "schema":{
                     "$ref":"#/definitions/dto.Response"
                  }
               }
            }
         }
      },
      "/products/all":{
         "get":{
            "security":[
               {
                  "bearerAuth":[

                  ]
               }
            ],
            "description":"Get all the products",
            "consumes":[
               "application/json"
            ],
            "produces":[
               "application/json"
            ],
            "tags":[
               "Products"
            ],
            "summary":"List existing products",
            "responses":{
               "200":{
                  "description":"OK",
                  "schema":{
                     "type":"array",
                     "items":{
                        "$ref":"#/definitions/entity.Product"
                     }
                  }
               },
               "401":{
                  "description":"Unauthorized",
                  "schema":{
                     "$ref":"#/definitions/dto.Response"
                  }
               }
            }
         }
      },
      "/products/save":{
         "post":{
            "security":[
               {
                  "bearerAuth":[

                  ]
               }
            ],
            "description":"Create a new products",
            "consumes":[
               "application/json"
            ],
            "produces":[
               "application/json"
            ],
            "tags":[
               "Products"
            ],
            "summary":"Create new product",
            "parameters":[
               {
                  "description":"Create product",
                  "name":"product",
                  "in":"body",
                  "required":true,
                  "schema":{
                     "$ref":"#/definitions/entity.Product"
                  }
               }
            ],
            "responses":{
               "200":{
                  "description":"OK",
                  "schema":{
                     "$ref":"#/definitions/dto.Response"
                  }
               },
               "400":{
                  "description":"Bad Request",
                  "schema":{
                     "$ref":"#/definitions/dto.Response"
                  }
               },
               "401":{
                  "description":"Unauthorized",
                  "schema":{
                     "$ref":"#/definitions/dto.Response"
                  }
               }
            }
         }
      },
      "/products/update{id}":{
         "post":{
            "security":[
               {
                  "bearerAuth":[

                  ]
               }
            ],
            "description":"Update a single product",
            "consumes":[
               "application/json"
            ],
            "produces":[
               "application/json"
            ],
            "tags":[
               "Products"
            ],
            "summary":"Update product",
            "parameters":[
               {
                  "type":"integer",
                  "description":"Product ID",
                  "name":"id",
                  "in":"path",
                  "required":true
               },
               {
                  "description":"Update product",
                  "name":"product",
                  "in":"body",
                  "required":true,
                  "schema":{
                     "$ref":"#/definitions/entity.Product"
                  }
               }
            ],
            "responses":{
               "200":{
                  "description":"OK",
                  "schema":{
                     "$ref":"#/definitions/dto.Response"
                  }
               },
               "400":{
                  "description":"Bad Request",
                  "schema":{
                     "$ref":"#/definitions/dto.Response"
                  }
               },
               "401":{
                  "description":"Unauthorized",
                  "schema":{
                     "$ref":"#/definitions/dto.Response"
                  }
               }
            }
         }
      },
      "/products/delete{id}":{
         "post":{
            "security":[
               {
                  "bearerAuth":[

                  ]
               }
            ],
            "description":"Delete a product",
            "consumes":[
               "application/json"
            ],
            "produces":[
               "application/json"
            ],
            "tags":[
               "Products"
            ],
            "summary":"Delete product",
            "parameters":[
               {
                  "type":"integer",
                  "description":"Product ID",
                  "name":"id",
                  "in":"path",
                  "required":true
               },
               {
                  "description":"Delete products",
                  "name":"product",
                  "in":"body",
                  "required":true,
                  "schema":{
                     "$ref":"#/definitions/entity.Product"
                  }
               }
            ],
            "responses":{
               "200":{
                  "description":"OK",
                  "schema":{
                     "$ref":"#/definitions/dto.Response"
                  }
               },
               "400":{
                  "description":"Bad Request",
                  "schema":{
                     "$ref":"#/definitions/dto.Response"
                  }
               },
               "401":{
                  "description":"Unauthorized",
                  "schema":{
                     "$ref":"#/definitions/dto.Response"
                  }
               }
            }
         }
      },
      "/brands/all":{
         "get":{
            "security":[
               {
                  "bearerAuth":[

                  ]
               }
            ],
            "description":"Get all the brands",
            "consumes":[
               "application/json"
            ],
            "produces":[
               "application/json"
            ],
            "tags":[
               "Brands"
            ],
            "summary":"List existing brands",
            "responses":{
               "200":{
                  "description":"OK",
                  "schema":{
                     "type":"array",
                     "items":{
                        "$ref":"#/definitions/entity.Brand"
                     }
                  }
               },
               "401":{
                  "description":"Unauthorized",
                  "schema":{
                     "$ref":"#/definitions/dto.Response"
                  }
               }
            }
         }
      },
      "/brands/save":{
         "post":{
            "security":[
               {
                  "bearerAuth":[

                  ]
               }
            ],
            "description":"Create a new brand",
            "consumes":[
               "application/json"
            ],
            "produces":[
               "application/json"
            ],
            "tags":[
               "Brands"
            ],
            "summary":"Create new brands",
            "parameters":[
               {
                  "description":"Create brands",
                  "name":"brand",
                  "in":"body",
                  "required":true,
                  "schema":{
                     "$ref":"#/definitions/entity.Brand"
                  }
               }
            ],
            "responses":{
               "200":{
                  "description":"OK",
                  "schema":{
                     "$ref":"#/definitions/dto.Response"
                  }
               },
               "400":{
                  "description":"Bad Request",
                  "schema":{
                     "$ref":"#/definitions/dto.Response"
                  }
               },
               "401":{
                  "description":"Unauthorized",
                  "schema":{
                     "$ref":"#/definitions/dto.Response"
                  }
               }
            }
         }
      },
      "/brands/update{id}":{
         "post":{
            "security":[
               {
                  "bearerAuth":[

                  ]
               }
            ],
            "description":"Update a single brand",
            "consumes":[
               "application/json"
            ],
            "produces":[
               "application/json"
            ],
            "tags":[
               "Brands"
            ],
            "summary":"Update brand",
            "parameters":[
               {
                  "type":"integer",
                  "description":"Brand ID",
                  "name":"id",
                  "in":"path",
                  "required":true
               },
               {
                  "description":"Update brand",
                  "name":"brand",
                  "in":"body",
                  "required":true,
                  "schema":{
                     "$ref":"#/definitions/entity.Brand"
                  }
               }
            ],
            "responses":{
               "200":{
                  "description":"OK",
                  "schema":{
                     "$ref":"#/definitions/dto.Response"
                  }
               },
               "400":{
                  "description":"Bad Request",
                  "schema":{
                     "$ref":"#/definitions/dto.Response"
                  }
               },
               "401":{
                  "description":"Unauthorized",
                  "schema":{
                     "$ref":"#/definitions/dto.Response"
                  }
               }
            }
         }
      },
      "/brands/delete{id}":{
         "post":{
            "security":[
               {
                  "bearerAuth":[

                  ]
               }
            ],
            "description":"Delete a brand",
            "consumes":[
               "application/json"
            ],
            "produces":[
               "application/json"
            ],
            "tags":[
               "Brands"
            ],
            "summary":"Delete brand",
            "parameters":[
               {
                  "type":"integer",
                  "description":"Brand ID",
                  "name":"id",
                  "in":"path",
                  "required":true
               },
               {
                  "description":"Delete Brand",
                  "name":"brand",
                  "in":"body",
                  "required":true,
                  "schema":{
                     "$ref":"#/definitions/entity.Brand"
                  }
               }
            ],
            "responses":{
               "200":{
                  "description":"OK",
                  "schema":{
                     "$ref":"#/definitions/dto.Response"
                  }
               },
               "400":{
                  "description":"Bad Request",
                  "schema":{
                     "$ref":"#/definitions/dto.Response"
                  }
               },
               "401":{
                  "description":"Unauthorized",
                  "schema":{
                     "$ref":"#/definitions/dto.Response"
                  }
               }
            }
         }
      }
   },
   "definitions":{
      "dto.JWT":{
         "type":"object",
         "properties":{
            "token":{
               "type":"string"
            }
         }
      },
      "dto.Response":{
         "type":"object",
         "properties":{
            "message":{
               "type":"string"
            }
         }
      },
      "entity.Brand":{
         "type":"object",
         "required":[
            "brand-id",
            "brand-name"
         ],
         "properties":{
            "brand-id":{
               "type":"string"
            },
            "brand-name":{
               "type":"string"
            },
            "brand-logo":{
               "type":"string"
            }
         }
      },
      "entity.Product":{
         "type":"object",
         "required":[
            "id",
            "name",
            "color",
            "sale-price",
            "order-cost",
            "quantity",
            "type",
            "brand-id"
         ],
         "properties":{
            "author":{
               "type":"object",
               "$ref":"#/definitions/entity.Product"
            },
            "id":{
               "type":"string"
            },
            "name":{
               "type":"string"
            },
            "color":{
               "type":"string"
            },
            "description":{
               "type":"string"
            },
            "sale-price":{
               "type":"number"
            },
            "order-cost":{
               "type":"number"
            },
            "quantity":{
               "type":"integer"
            },
            "type":{
               "type":"string"
            },
            "image":{
               "type":"string"
            },
            "brand-id":{
               "type":"string"
            }
         }
      }
   },
   "securityDefinitions":{
      "bearerAuth":{
         "type":"apiKey",
         "name":"Authorization",
         "in":"header"
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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
	swag.Register(swag.Name, &s{})
}
