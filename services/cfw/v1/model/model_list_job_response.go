package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobResponse Response Object
type ListJobResponse struct {

	// 错误码
	ErrorCode string `json:"error_code"`

	// 错误描述
	ErrorMsg string `json:"error_msg"`

	// 执行结果
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobResponse struct{}"
	}

	return strings.Join([]string{"ListJobResponse", string(data)}, " ")
}
