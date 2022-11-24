package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ResetGaussMySqlDatabasePasswordRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *ResetDatabasePasswordRequest `json:"body,omitempty"`
}

func (o ResetGaussMySqlDatabasePasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetGaussMySqlDatabasePasswordRequest struct{}"
	}

	return strings.Join([]string{"ResetGaussMySqlDatabasePasswordRequest", string(data)}, " ")
}
