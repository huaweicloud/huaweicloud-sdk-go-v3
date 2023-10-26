package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExecuteCommandRequestBody struct {

	// 客户端ID
	ClientId *string `json:"client_id,omitempty"`

	// 命令
	Command *string `json:"command,omitempty"`

	// 数据库编号
	Database *int32 `json:"database,omitempty"`
}

func (o ExecuteCommandRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteCommandRequestBody struct{}"
	}

	return strings.Join([]string{"ExecuteCommandRequestBody", string(data)}, " ")
}
