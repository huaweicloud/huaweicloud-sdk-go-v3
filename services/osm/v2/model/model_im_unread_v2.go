package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImUnreadV2 struct {
	// 状态

	Status *int32 `json:"status,omitempty"`
	// 工单id

	IncidentId *string `json:"incident_id,omitempty"`
	// 未读数量

	UnreadNum *int32 `json:"unread_num,omitempty"`
}

func (o ImUnreadV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImUnreadV2 struct{}"
	}

	return strings.Join([]string{"ImUnreadV2", string(data)}, " ")
}
