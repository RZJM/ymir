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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/users/{userId}/repo/{repoId}/branch/{branchId}/assets": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Query single or set of assets.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Repo ID",
                        "name": "repoId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Branch ID",
                        "name": "branchId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Offset, default is 0",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "limit, default is 20",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "e.g. class_ids=1,3,7",
                        "name": "class_ids",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "e.g. current_asset_id=xxxyyyzzz",
                        "name": "current_asset_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "e.g. cm_types=FN,TP,MTP,IGNORED",
                        "name": "cm_types",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "ck pairs, e.g. cks=xxx,xxx:,xxx:yyy",
                        "name": "cks",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "tag pairs, e.g. cks=xxx,xxx:,xxx:yyy",
                        "name": "tags",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "'code': 0, 'msg': 'Success', 'Success': true, 'Data': constants.QueryAssetsResult",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/users/{userId}/repo/{repoId}/branch/{branchId}/dataset_meta_count": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Query dataset info, lightweight api.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Repo ID",
                        "name": "repoId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Branch ID",
                        "name": "branchId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "'code': 0, 'msg': 'Success', 'Success': true, 'Data': constants.QueryDatasetStatsResult",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/users/{userId}/repo/{repoId}/branch/{branchId}/dataset_stats": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Query dataset Stats.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Repo ID",
                        "name": "repoId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Branch ID",
                        "name": "branchId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "e.g. class_ids=1,3,7",
                        "name": "class_ids",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "'code': 0, 'msg': 'Success', 'Success': true, 'Data': constants.QueryDatasetStatsResult",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/users/{userId}/repo/{repoId}/dataset_duplication": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Query dataset dups.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Repo ID",
                        "name": "repoId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "e.g. candidate_dataset_ids=xxx,yyy",
                        "name": "candidate_dataset_ids",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "'code': 0, 'msg': 'Success', 'Success': true, 'Data': 'duplication: 50, total_count: {xxx: 100, yyy: 200}'",
                        "schema": {
                            "type": "string"
                        }
                    }
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