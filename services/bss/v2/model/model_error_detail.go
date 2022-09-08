package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ErrorDetail struct {

	// 状态码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述信息。
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 出现错误的客户ID或批量处理消息的ID。
	Id *string `json:"id,omitempty"`
}

func (o ErrorDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ErrorDetail struct{}"
	}

	return strings.Join([]string{"ErrorDetail", string(data)}, " ")
}
