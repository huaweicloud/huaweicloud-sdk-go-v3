package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateSqlserverDbUserRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty" xml:"X-Language"`

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	Body *SqlserverUserForCreation `json:"body,omitempty" xml:"body"`
}

func (o CreateSqlserverDbUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSqlserverDbUserRequest struct{}"
	}

	return strings.Join([]string{"CreateSqlserverDbUserRequest", string(data)}, " ")
}
