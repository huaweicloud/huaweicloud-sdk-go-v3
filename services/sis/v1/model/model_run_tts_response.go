package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RunTtsResponse struct {

	// 服务内部的令牌，可用于在日志中追溯具体流程。  在某些错误情况下可能没有此令牌字符串。
	TraceId *string `json:"trace_id,omitempty" xml:"trace_id"`

	Result         *CustomResult `json:"result,omitempty" xml:"result"`
	HttpStatusCode int           `json:"-"`
}

func (o RunTtsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunTtsResponse struct{}"
	}

	return strings.Join([]string{"RunTtsResponse", string(data)}, " ")
}
