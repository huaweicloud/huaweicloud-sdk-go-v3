package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecollectIndexUsageNewResponse Response Object
type RecollectIndexUsageNewResponse struct {

	// 是否成功
	Success *bool `json:"success,omitempty"`

	// 状态值（1：成功，2：失败）
	Status *int32 `json:"status,omitempty"`

	// 错误文案，只有在状态为2时才显示
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RecollectIndexUsageNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecollectIndexUsageNewResponse struct{}"
	}

	return strings.Join([]string{"RecollectIndexUsageNewResponse", string(data)}, " ")
}
