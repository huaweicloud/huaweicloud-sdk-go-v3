package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchFailure struct {

	// 发布或下线失败的API ID
	ApiId *string `json:"api_id,omitempty" xml:"api_id"`

	// 发布或下线失败的APi名称
	ApiName *string `json:"api_name,omitempty" xml:"api_name"`

	// 发布或下线失败的错误码
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code"`

	// 发布或下线失败的错误信息
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg"`
}

func (o BatchFailure) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchFailure struct{}"
	}

	return strings.Join([]string{"BatchFailure", string(data)}, " ")
}
