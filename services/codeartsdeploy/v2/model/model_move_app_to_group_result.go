package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MoveAppToGroupResult 移动应用至指定分组结果
type MoveAppToGroupResult struct {

	// 是否失败
	Code *string `json:"code,omitempty"`

	// 应用id
	ApplicationId *string `json:"application_id,omitempty"`

	// 应用名称
	ApplicationName *string `json:"application_name,omitempty"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o MoveAppToGroupResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MoveAppToGroupResult struct{}"
	}

	return strings.Join([]string{"MoveAppToGroupResult", string(data)}, " ")
}
