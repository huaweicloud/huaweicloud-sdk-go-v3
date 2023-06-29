package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ErrorSample 错误样例
type ErrorSample struct {

	// 检测源描述。
	Source string `json:"source"`

	// 此错误共计次数。
	Count *int32 `json:"count,omitempty"`

	// 错误数据和错误提示消息。
	Msg *string `json:"msg,omitempty"`
}

func (o ErrorSample) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ErrorSample struct{}"
	}

	return strings.Join([]string{"ErrorSample", string(data)}, " ")
}
