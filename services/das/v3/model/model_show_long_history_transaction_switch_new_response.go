package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLongHistoryTransactionSwitchNewResponse Response Object
type ShowLongHistoryTransactionSwitchNewResponse struct {

	// 开关状态
	SwitchOn *bool `json:"switch_on,omitempty"`

	// 长事务阈值
	Threshold *int64 `json:"threshold,omitempty"`

	// 是否可以开启
	CanOpen *bool `json:"can_open,omitempty"`

	// 无法开启原因
	CantOpenMsg    *string `json:"cant_open_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowLongHistoryTransactionSwitchNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLongHistoryTransactionSwitchNewResponse struct{}"
	}

	return strings.Join([]string{"ShowLongHistoryTransactionSwitchNewResponse", string(data)}, " ")
}
