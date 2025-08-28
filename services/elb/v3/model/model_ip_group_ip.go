package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IpGroupIp **参数解释**：IP地址组中IP列表的IP地址信息。  **约束限制**：不涉及
type IpGroupIp struct {

	// **参数解释**：需要从IP地址组中删除的IP地址，可以是单个IP地址、IP地址段和连续IP地址范围。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	Ip string `json:"ip"`
}

func (o IpGroupIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpGroupIp struct{}"
	}

	return strings.Join([]string{"IpGroupIp", string(data)}, " ")
}
