package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDecoyPortStatisticsResponse Response Object
type ShowDecoyPortStatisticsResponse struct {

	// **参数解释**： 开启防护的主机数量 **取值范围**： 最小值0，最大值2147483647
	HostNum *int32 `json:"host_num,omitempty"`

	// **参数解释**： 策略数量 **取值范围**： 最小值0，最大值2147483647
	PolicyNum *int32 `json:"policy_num,omitempty"`

	// **参数解释**： 未处理的事件数量 **取值范围**： 最小值0，最大值2147483647
	EventNum *int32 `json:"event_num,omitempty"`

	// **参数解释**： linux的推荐端口配置 **取值范围**： 不涉及
	LinuxPropose *[]DecoyPortBasePortStruct `json:"linux_propose,omitempty"`

	// **参数解释**： windows的推荐端口配置 **取值范围**： 不涉及
	WindowsPropose *[]DecoyPortBasePortStruct `json:"windows_propose,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ShowDecoyPortStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDecoyPortStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowDecoyPortStatisticsResponse", string(data)}, " ")
}
