package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IpReputationMapResponseBodyIpReputationMap **参数解释：** 威胁情报控制防护选项的内容类型 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type IpReputationMapResponseBodyIpReputationMap struct {

	// **参数解释：** 威胁情报控制的内容类型 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Idc *[]string `json:"idc,omitempty"`
}

func (o IpReputationMapResponseBodyIpReputationMap) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpReputationMapResponseBodyIpReputationMap struct{}"
	}

	return strings.Join([]string{"IpReputationMapResponseBodyIpReputationMap", string(data)}, " ")
}
