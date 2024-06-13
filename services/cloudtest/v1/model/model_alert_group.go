package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlertGroup struct {
	DingTalkHookUrl *string `json:"dingTalkHookUrl,omitempty"`

	// 告警组名称
	GroupName *string `json:"group_name,omitempty"`

	// 告警组ID
	Id *string `json:"id,omitempty"`

	WeChatWorkHookUrl *string `json:"weChatWorkHookUrl,omitempty"`

	WeLinkGroupNo *string `json:"weLinkGroupNo,omitempty"`
}

func (o AlertGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertGroup struct{}"
	}

	return strings.Join([]string{"AlertGroup", string(data)}, " ")
}
