package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableDefensePolicyResponse Response Object
type EnableDefensePolicyResponse struct {

	// 内部错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 内部错误描述
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 任务ID，后续可根据该ID查询本任务状态。 本字段为后续的任务审计扩展，暂时不需要，先保留。
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o EnableDefensePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableDefensePolicyResponse struct{}"
	}

	return strings.Join([]string{"EnableDefensePolicyResponse", string(data)}, " ")
}
