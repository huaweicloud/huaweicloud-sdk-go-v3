package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EipInfo **参数解释：** 负载均衡器绑定的EIP信息 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type EipInfo struct {

	// **参数解释：** 弹性IP的ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	EipId *string `json:"eip_id,omitempty"`

	// **参数解释：** 弹性IP的IP地址 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	EipAddress *string `json:"eip_address,omitempty"`

	// **参数解释：** IP版本号 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	IpVersion *int32 `json:"ip_version,omitempty"`
}

func (o EipInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EipInfo struct{}"
	}

	return strings.Join([]string{"EipInfo", string(data)}, " ")
}
