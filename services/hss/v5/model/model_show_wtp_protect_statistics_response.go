package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWtpProtectStatisticsResponse Response Object
type ShowWtpProtectStatisticsResponse struct {

	// **参数解释**: 防护服务器总数，包含防护状态是防护中、部分防护、防护暂停、防护失败、防护中断和开启中共6种状态的服务器总数 **取值范围**: 最小值0，最大值1000000
	ProtectHostNum *int32 `json:"protect_host_num,omitempty"`

	// **参数解释**: 防护状态为防护中的服务器总数 **取值范围**: 最小值0，最大值1000000
	ProtectSuccessHostNum *int32 `json:"protect_success_host_num,omitempty"`

	// **参数解释**: 防护状态为防护失败的服务器总数 **取值范围**: 最小值0，最大值1000000
	ProtectFailHostNum *int32 `json:"protect_fail_host_num,omitempty"`

	// **参数解释**: 近168小时防护事件数 **取值范围**: 最小值0，最大值50000000
	AntiTamperingNum *int32 `json:"anti_tampering_num,omitempty"`
	HttpStatusCode   int    `json:"-"`
}

func (o ShowWtpProtectStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWtpProtectStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowWtpProtectStatisticsResponse", string(data)}, " ")
}
