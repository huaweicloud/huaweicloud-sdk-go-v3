package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateGaussMySqlDatabaseRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty" xml:"X-Language"`

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	Body *CreateGaussMySqlDatabaseRequestBody `json:"body,omitempty" xml:"body"`
}

func (o CreateGaussMySqlDatabaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGaussMySqlDatabaseRequest struct{}"
	}

	return strings.Join([]string{"CreateGaussMySqlDatabaseRequest", string(data)}, " ")
}
