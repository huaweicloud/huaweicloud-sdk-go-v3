package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MergeErrorDto struct {

	// 合并人名字
	MergeUserName *string `json:"merge_user_name,omitempty"`

	// 合入时间
	MergeStartTime *string `json:"merge_start_time,omitempty"`

	// 错误标题
	ErrorTitle *string `json:"error_title,omitempty"`

	// 状态码
	HttpCodeStatus *string `json:"http_code_status,omitempty"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息
	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o MergeErrorDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MergeErrorDto struct{}"
	}

	return strings.Join([]string{"MergeErrorDto", string(data)}, " ")
}
