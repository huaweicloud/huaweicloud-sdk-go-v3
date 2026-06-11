package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IpGroupsDetail **参数解释：** IP地址组中包含的IP或网段列表。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
type IpGroupsDetail struct {

	// **参数解释：** IP地址或网段。 **约束限制：** 支持IPv4。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Ip string `json:"ip"`

	// **参数解释：** 备注信息。 **约束限制：** 最长255字符。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Description string `json:"description"`
}

func (o IpGroupsDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpGroupsDetail struct{}"
	}

	return strings.Join([]string{"IpGroupsDetail", string(data)}, " ")
}
