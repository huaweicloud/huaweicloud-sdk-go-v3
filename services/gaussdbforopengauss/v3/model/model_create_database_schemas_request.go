package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateDatabaseSchemasRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty" xml:"X-Language"`

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	Body *GaussDBforOpenGaussDatabaseSchemaReq `json:"body,omitempty" xml:"body"`
}

func (o CreateDatabaseSchemasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatabaseSchemasRequest struct{}"
	}

	return strings.Join([]string{"CreateDatabaseSchemasRequest", string(data)}, " ")
}
