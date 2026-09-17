package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetHistoryTransactionSwitchNewResponse Response Object
type SetHistoryTransactionSwitchNewResponse struct {

	// 是否成功
	Success *bool `json:"success,omitempty"`

	// 状态值
	Status *int32 `json:"status,omitempty"`

	// 错误信息
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetHistoryTransactionSwitchNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetHistoryTransactionSwitchNewResponse struct{}"
	}

	return strings.Join([]string{"SetHistoryTransactionSwitchNewResponse", string(data)}, " ")
}
