package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 错误样例
type ErrorSample struct {

	// 检测源描述。
	Source string `json:"source" xml:"source"`

	// 此错误共计次数。
	Count *int32 `json:"count,omitempty" xml:"count"`

	// 错误数据和错误提示消息。
	Msg *string `json:"msg,omitempty" xml:"msg"`
}

func (o ErrorSample) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ErrorSample struct{}"
	}

	return strings.Join([]string{"ErrorSample", string(data)}, " ")
}
