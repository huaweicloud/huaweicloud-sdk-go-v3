package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProtectionStatisticsResponse Response Object
type ShowProtectionStatisticsResponse struct {

	// **参数解释**： 已防护服务器数量 **取值范围**： 最小值0，最大值2147483647
	CountProtectedServer *int32 `json:"count_protected_server,omitempty"`

	// **参数解释**： 未防护服务器数量 **取值范围**： 最小值0，最大值2147483647
	CountNeedProtectedServer *int32 `json:"count_need_protected_server,omitempty"`

	// **参数解释**： 防护事件 **取值范围**： 最小值0，最大值2147483647
	CountProtectionEvent *int32 `json:"count_protection_event,omitempty"`
	HttpStatusCode       int    `json:"-"`
}

func (o ShowProtectionStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProtectionStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowProtectionStatisticsResponse", string(data)}, " ")
}
