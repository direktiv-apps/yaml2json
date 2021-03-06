// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Converts incoming json, xml or properties to json, xml or properties.",
    "title": "yaml2json",
    "version": "1.0.0",
    "x-direktiv-meta": {
      "category": "Tools",
      "container": "direktiv/yaml2json",
      "issues": "https://github.com/direktiv-apps/yaml2json/issues",
      "license": "[Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)",
      "long-description": "This [yq](https://github.com/mikefarah/yq) based service can convert incoming json, xml, csv, tsv or properties data into the same formats.\nThe incoming data can be a string, file path or base64-encoded string. ",
      "maintainer": "[direktiv.io](https://www.direktiv.io)",
      "url": "https://github.com/direktiv-apps/yaml2json"
    }
  },
  "paths": {
    "/": {
      "post": {
        "parameters": [
          {
            "type": "string",
            "description": "direktiv action id is an UUID. \nFor development it can be set to 'development'\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          },
          {
            "type": "string",
            "description": "direktiv temp dir is the working directory for that request\nFor development it can be set to e.g. '/tmp'\n",
            "name": "Direktiv-TempDir",
            "in": "header"
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "required": [
                "data"
              ],
              "properties": {
                "data": {
                  "description": "Depending on the input-type this value can be a file path, base64 string or a plain string.",
                  "type": "string",
                  "example": "aHR0cHM6Ly93d3cueW91dHViZS5jb20vd2F0Y2g/dj1kUXc0dzlXZ1hjUQ=="
                },
                "input-format": {
                  "description": "Input format. Possible values are yaml,xml,prop.",
                  "type": "string",
                  "default": "yaml",
                  "enum": [
                    "yaml",
                    "xml",
                    "props"
                  ],
                  "example": "xml"
                },
                "input-type": {
                  "type": "string",
                  "default": "string",
                  "enum": [
                    "file",
                    "string",
                    "base64"
                  ],
                  "example": "base64"
                },
                "output-format": {
                  "description": "Output format. Possible values are yaml, json, props, csv, tsv, xml.",
                  "type": "string",
                  "default": "json",
                  "enum": [
                    "yaml",
                    "json",
                    "props",
                    "csv",
                    "tsv",
                    "xml"
                  ],
                  "example": "json"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "nice greeting",
            "schema": {
              "type": "object",
              "properties": {
                "output": {
                  "type": "object",
                  "properties": {
                    "result": {
                      "additionalProperties": false
                    },
                    "success": {
                      "type": "boolean"
                    }
                  }
                }
              }
            },
            "examples": {
              "greeting": "Hello YourName"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            },
            "headers": {
              "Direktiv-ErrorCode": {
                "type": "string"
              },
              "Direktiv-ErrorMessage": {
                "type": "string"
              }
            }
          }
        },
        "x-direktiv": {
          "cmds": [
            {
              "action": "exec",
              "exec": "/y2j.sh \"{{ .Data }}\" {{ default \"yaml\" .InputFormat }} {{ default \"json\" .OutputFormat }} {{ default \"string\" .InputType }}"
            }
          ],
          "output": "{\n  \"output\": {{ index . 0 | toJson }}\n}\n"
        },
        "x-direktiv-errors": {
          "io.direktiv.command.error": "Command execution failed",
          "io.direktiv.output.error": "Template error for output generation of the service",
          "io.direktiv.ri.error": "Can not create information object from request"
        },
        "x-direktiv-examples": [
          {
            "content": "- id: convert\n     type: action\n     action:\n       function: yaml2json\n       input:\n         data: 'hello: world'\n         output-format: json\n         input-type: string",
            "title": "String"
          },
          {
            "content": "- id: set\n    type: setter\n    variables:\n    - key: indata\n      scope: instance\n      mimeType: text/plain\n      value: 'hello: world'\n    transition: modify\n  - id: modify \n    type: action\n    action:\n      function: yaml2json\n      files:\n        - key: indata\n          scope: instance\n      input: \n        data: indata\n        output-format: json\n        input-type: file",
            "title": "File"
          },
          {
            "content": "- id: convert\n     type: action\n     action:\n       function: yaml2json\n       input:\n         data: aGVsbG86IHdvcmxk\n         output-format: xml\n         input-type: base64",
            "title": "Base64"
          }
        ],
        "x-direktiv-function": "functions:\n  - id: yaml2json\n    image: direktiv/yaml2json\n    type: knative-workflow"
      },
      "delete": {
        "parameters": [
          {
            "type": "string",
            "description": "On cancel Direktiv sends a DELETE request to\nthe action with id in the header\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          }
        ],
        "responses": {
          "200": {
            "description": ""
          }
        },
        "x-direktiv": {
          "cancel": "echo 'cancel {{ .DirektivActionID }}'"
        }
      }
    }
  },
  "definitions": {
    "direktivFile": {
      "type": "object",
      "x-go-type": {
        "import": {
          "package": "github.com/direktiv/apps/go/pkg/apps"
        },
        "type": "DirektivFile"
      }
    },
    "error": {
      "type": "object",
      "required": [
        "errorCode",
        "errorMessage"
      ],
      "properties": {
        "errorCode": {
          "type": "string"
        },
        "errorMessage": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Converts incoming json, xml or properties to json, xml or properties.",
    "title": "yaml2json",
    "version": "1.0.0",
    "x-direktiv-meta": {
      "category": "Tools",
      "container": "direktiv/yaml2json",
      "issues": "https://github.com/direktiv-apps/yaml2json/issues",
      "license": "[Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)",
      "long-description": "This [yq](https://github.com/mikefarah/yq) based service can convert incoming json, xml, csv, tsv or properties data into the same formats.\nThe incoming data can be a string, file path or base64-encoded string. ",
      "maintainer": "[direktiv.io](https://www.direktiv.io)",
      "url": "https://github.com/direktiv-apps/yaml2json"
    }
  },
  "paths": {
    "/": {
      "post": {
        "parameters": [
          {
            "type": "string",
            "description": "direktiv action id is an UUID. \nFor development it can be set to 'development'\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          },
          {
            "type": "string",
            "description": "direktiv temp dir is the working directory for that request\nFor development it can be set to e.g. '/tmp'\n",
            "name": "Direktiv-TempDir",
            "in": "header"
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "required": [
                "data"
              ],
              "properties": {
                "data": {
                  "description": "Depending on the input-type this value can be a file path, base64 string or a plain string.",
                  "type": "string",
                  "example": "aHR0cHM6Ly93d3cueW91dHViZS5jb20vd2F0Y2g/dj1kUXc0dzlXZ1hjUQ=="
                },
                "input-format": {
                  "description": "Input format. Possible values are yaml,xml,prop.",
                  "type": "string",
                  "default": "yaml",
                  "enum": [
                    "yaml",
                    "xml",
                    "props"
                  ],
                  "example": "xml"
                },
                "input-type": {
                  "type": "string",
                  "default": "string",
                  "enum": [
                    "file",
                    "string",
                    "base64"
                  ],
                  "example": "base64"
                },
                "output-format": {
                  "description": "Output format. Possible values are yaml, json, props, csv, tsv, xml.",
                  "type": "string",
                  "default": "json",
                  "enum": [
                    "yaml",
                    "json",
                    "props",
                    "csv",
                    "tsv",
                    "xml"
                  ],
                  "example": "json"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "nice greeting",
            "schema": {
              "type": "object",
              "properties": {
                "output": {
                  "type": "object",
                  "properties": {
                    "result": {
                      "additionalProperties": false
                    },
                    "success": {
                      "type": "boolean"
                    }
                  }
                }
              }
            },
            "examples": {
              "greeting": "Hello YourName"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            },
            "headers": {
              "Direktiv-ErrorCode": {
                "type": "string"
              },
              "Direktiv-ErrorMessage": {
                "type": "string"
              }
            }
          }
        },
        "x-direktiv": {
          "cmds": [
            {
              "action": "exec",
              "exec": "/y2j.sh \"{{ .Data }}\" {{ default \"yaml\" .InputFormat }} {{ default \"json\" .OutputFormat }} {{ default \"string\" .InputType }}"
            }
          ],
          "output": "{\n  \"output\": {{ index . 0 | toJson }}\n}\n"
        },
        "x-direktiv-errors": {
          "io.direktiv.command.error": "Command execution failed",
          "io.direktiv.output.error": "Template error for output generation of the service",
          "io.direktiv.ri.error": "Can not create information object from request"
        },
        "x-direktiv-examples": [
          {
            "content": "- id: convert\n     type: action\n     action:\n       function: yaml2json\n       input:\n         data: 'hello: world'\n         output-format: json\n         input-type: string",
            "title": "String"
          },
          {
            "content": "- id: set\n    type: setter\n    variables:\n    - key: indata\n      scope: instance\n      mimeType: text/plain\n      value: 'hello: world'\n    transition: modify\n  - id: modify \n    type: action\n    action:\n      function: yaml2json\n      files:\n        - key: indata\n          scope: instance\n      input: \n        data: indata\n        output-format: json\n        input-type: file",
            "title": "File"
          },
          {
            "content": "- id: convert\n     type: action\n     action:\n       function: yaml2json\n       input:\n         data: aGVsbG86IHdvcmxk\n         output-format: xml\n         input-type: base64",
            "title": "Base64"
          }
        ],
        "x-direktiv-function": "functions:\n  - id: yaml2json\n    image: direktiv/yaml2json\n    type: knative-workflow"
      },
      "delete": {
        "parameters": [
          {
            "type": "string",
            "description": "On cancel Direktiv sends a DELETE request to\nthe action with id in the header\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          }
        ],
        "responses": {
          "200": {
            "description": ""
          }
        },
        "x-direktiv": {
          "cancel": "echo 'cancel {{ .DirektivActionID }}'"
        }
      }
    }
  },
  "definitions": {
    "PostOKBodyOutput": {
      "type": "object",
      "properties": {
        "result": {
          "additionalProperties": false
        },
        "success": {
          "type": "boolean"
        }
      }
    },
    "direktivFile": {
      "type": "object",
      "x-go-type": {
        "import": {
          "package": "github.com/direktiv/apps/go/pkg/apps"
        },
        "type": "DirektivFile"
      }
    },
    "error": {
      "type": "object",
      "required": [
        "errorCode",
        "errorMessage"
      ],
      "properties": {
        "errorCode": {
          "type": "string"
        },
        "errorMessage": {
          "type": "string"
        }
      }
    }
  }
}`))
}
