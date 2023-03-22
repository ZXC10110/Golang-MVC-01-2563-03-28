package View

import "github.com/swaggo/swag"

const docTemplate = `{
  "openapi": "3.0.3",
  "info": {
    "title": "Golang MVC Structure",
    "version": "1.0.11"
  },
  "tags": [
    {
      "name": "Hospitals"
    },
    {
      "name": "Patients"
    },
    {
      "name": "PatientsCovidStatus"
    }
  ],
  "paths": {
    "/getAllHospitals": {
      "get": {
        "tags": [
          "Hospitals"
        ],
        "summary": "get Hospitals",
        "responses": {
          "201": {
            "description": "get Successful",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/Hospitals"
                }
              }
            }
          }
        }
      }
    },
    "/getAllPatients": {
      "get": {
        "tags": [
          "Patients"
        ],
        "summary": "get Patients",
        "responses": {
          "201": {
            "description": "get Successful",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/Patients"
                }
              }
            }
          }
        }
      }
    },
    "/getAllPatientsCovidStatus": {
      "get": {
        "tags": [
          "PatientsCovidStatus"
        ],
        "summary": "get PatientsCovidStatus",
        "responses": {
          "201": {
            "description": "get Successful",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/PatientsCovidStatus"
                }
              }
            }
          }
        }
      }
    },
    "/getCovidPatientInfo": {
      "get": {
        "tags": [
          "PatientsCovidStatus"
        ],
        "summary": "get PatientsCovidStatus",
        "responses": {
          "201": {
            "description": "get Successful",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/PatientsCovidStatus"
                }
              }
            }
          }
        }
      }
    }, 
    "/getCountPatient": {
      "get": {
        "tags": [
          "Hospitals"
        ],
        "summary": "get Patient Per Hospital",
        "responses": {
          "201": {
            "description": "get Successful",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/Hospitals"
                }
              }
            }
          }
        }
      }
    },
    "/getOrderByCountPatient": {
      "get": {
        "tags": [
          "Hospitals"
        ],
        "summary": "get Order Patient Per Hospital",
        "responses": {
          "201": {
            "description": "get Successful",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/Hospitals"
                }
              }
            }
          }
        }
      }
    }
  },
  "components": {
    "schemas": {
      "Hospitals": {
        "type": "object",
        "properties": {
          "hid": {
            "type": "String",
            "example": "123"
          }, 
          "title": {
            "type": "String",
            "example": "123"
          }
        }
      },
      "Patients": {
        "type": "object",
        "properties": {
          "hnid": {
            "type": "String",
            "example": "123"
          }, 
          "first_name": {
            "type": "String",
            "example": "123"
          }, 
          "last_name": {
            "type": "String",
            "example": "123"
          }, 
          "hid": {
            "type": "String",
            "example": "123"
          }
        }
      },
      "PatientsCovidStatus": {
        "type": "object",
        "properties": {
          "hnid": {
            "type": "String",
            "example": "123"
          }, 
          "covid_status": {
            "type": "String",
            "example": "123"
          }
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
	Title:            "Golang MVC",
	Description:      "This is a sample server todo server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
