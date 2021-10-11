package model

import (
	"encoding/json"

	"strings"
)

type BatchFailure struct {
	// 发布或下线失败的API ID

	ApiId *string `json:"api_id,omitempty"`
	// 发布或下线失败的APi名称

	ApiName *string `json:"api_name,omitempty"`
	// 发布或下线失败的错误码

	ErrorCode *string `json:"error_code,omitempty"`
	// 发布或下线失败的错误信息

	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o BatchFailure) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchFailure struct{}"
	}

	return strings.Join([]string{"BatchFailure", string(data)}, " ")
}
