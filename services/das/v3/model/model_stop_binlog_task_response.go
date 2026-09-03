package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopBinlogTaskResponse Response Object
type StopBinlogTaskResponse struct {

	// 是否成功
	Success *bool `json:"success,omitempty"`

	// 状态。取值范围：1（成功）、2（失败）
	Status *int32 `json:"status,omitempty"`

	// 错误文案，只有在状态为2时才显示
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StopBinlogTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopBinlogTaskResponse struct{}"
	}

	return strings.Join([]string{"StopBinlogTaskResponse", string(data)}, " ")
}
