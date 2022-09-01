package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询资源返回的异常信息
type QueryError struct {

	// 错误编码。
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code"`

	// 错误信息。
	ErrorMessage *string `json:"error_message,omitempty" xml:"error_message"`
}

func (o QueryError) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryError struct{}"
	}

	return strings.Join([]string{"QueryError", string(data)}, " ")
}
