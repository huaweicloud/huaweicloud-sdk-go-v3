package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 批量操作失败的返回结果
type BatchFailure struct {
	// 操作失败的API ID

	ApiId *string `json:"api_id,omitempty"`
	// 操作失败的APi名称

	ApiName *string `json:"api_name,omitempty"`
	// 操作失败的错误码

	ErrorCode *string `json:"error_code,omitempty"`
	// 操作失败的错误信息

	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o BatchFailure) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchFailure struct{}"
	}

	return strings.Join([]string{"BatchFailure", string(data)}, " ")
}
