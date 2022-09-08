package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateUnreadNewInstantMsgV2Req struct {

	// 组id
	GroupId *string `json:"group_id,omitempty"`
}

func (o UpdateUnreadNewInstantMsgV2Req) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUnreadNewInstantMsgV2Req struct{}"
	}

	return strings.Join([]string{"UpdateUnreadNewInstantMsgV2Req", string(data)}, " ")
}
