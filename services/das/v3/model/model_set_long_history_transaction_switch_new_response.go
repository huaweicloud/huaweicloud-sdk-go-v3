package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetLongHistoryTransactionSwitchNewResponse Response Object
type SetLongHistoryTransactionSwitchNewResponse struct {

	// 是否成功
	Success *bool `json:"success,omitempty"`

	// 状态值
	Status *int32 `json:"status,omitempty"`

	// 错误信息
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetLongHistoryTransactionSwitchNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetLongHistoryTransactionSwitchNewResponse struct{}"
	}

	return strings.Join([]string{"SetLongHistoryTransactionSwitchNewResponse", string(data)}, " ")
}
