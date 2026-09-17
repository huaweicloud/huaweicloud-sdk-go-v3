package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetMissingIndexSwitchNewResponse Response Object
type SetMissingIndexSwitchNewResponse struct {

	// 是否成功
	Success *bool `json:"success,omitempty"`

	// 状态值（1：成功，2：失败无需轮询，3：失败需要轮询）
	Status *int32 `json:"status,omitempty"`

	// 错误信息
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetMissingIndexSwitchNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetMissingIndexSwitchNewResponse struct{}"
	}

	return strings.Join([]string{"SetMissingIndexSwitchNewResponse", string(data)}, " ")
}
