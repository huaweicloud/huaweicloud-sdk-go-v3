package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGaussMySqlDatabaseRequest Request Object
type CreateGaussMySqlDatabaseRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *CreateGaussMySqlDatabaseRequestBody `json:"body,omitempty"`
}

func (o CreateGaussMySqlDatabaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGaussMySqlDatabaseRequest struct{}"
	}

	return strings.Join([]string{"CreateGaussMySqlDatabaseRequest", string(data)}, " ")
}
