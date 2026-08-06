package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHistoryTransactionSwitchNewResponse Response Object
type ShowHistoryTransactionSwitchNewResponse struct {

	// 开关状态
	SwitchOn *bool `json:"switch_on,omitempty"`

	// 保存时长
	RetentionHours *int64 `json:"retention_hours,omitempty"`

	// 是否可以开启
	CanOpen *bool `json:"can_open,omitempty"`

	// 无法开启原因
	CantOpenMsg    *string `json:"cant_open_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowHistoryTransactionSwitchNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHistoryTransactionSwitchNewResponse struct{}"
	}

	return strings.Join([]string{"ShowHistoryTransactionSwitchNewResponse", string(data)}, " ")
}
