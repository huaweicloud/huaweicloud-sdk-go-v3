package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CocUpdateChangeRequestBodyV2TicketInfo 主单字段
type CocUpdateChangeRequestBodyV2TicketInfo struct {

	// 阶段
	Phase *string `json:"phase,omitempty"`

	// 状态
	WorkFlowStatus *string `json:"work_flow_status,omitempty"`
}

func (o CocUpdateChangeRequestBodyV2TicketInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CocUpdateChangeRequestBodyV2TicketInfo struct{}"
	}

	return strings.Join([]string{"CocUpdateChangeRequestBodyV2TicketInfo", string(data)}, " ")
}
