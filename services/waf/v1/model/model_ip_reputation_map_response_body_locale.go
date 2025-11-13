package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IpReputationMapResponseBodyLocale **参数解释：** 威胁情报控制防护选项里，各个选项的描述 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type IpReputationMapResponseBodyLocale struct {

	// **参数解释：** 鹏博士，一家提供互联网数据中心、云计算等服务的企业 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	DrPeng *string `json:"Dr.Peng,omitempty"`

	// **参数解释：** 谷歌公司，全球知名的科技企业，提供搜索引擎、云计算等服务 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Google *string `json:"Google,omitempty"`

	// **参数解释：** 腾讯，中国知名的互联网企业，提供社交、游戏、金融等服务 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Tencent *string `json:"Tencent,omitempty"`

	// **参数解释：** 美团网，中国领先的生活服务电子商务平台 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	MeiTuan *string `json:"MeiTuan,omitempty"`

	// **参数解释：** 微软公司，全球知名的科技企业，提供操作系统、办公软件等服务 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Microsoft *string `json:"Microsoft,omitempty"`

	// **参数解释：** 阿里云，阿里巴巴集团旗下的云计算品牌 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	AliCloud *string `json:"AliCloud,omitempty"`

	// **参数解释：** 亚马逊，全球知名的电子商务和云计算企业 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Amazon *string `json:"Amazon,omitempty"`

	// **参数解释：** 世纪互联，中国领先的电信中立互联网基础设施服务提供商 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Vnet *string `json:"VNET,omitempty"`

	// **参数解释：** 华为，全球知名的通信技术企业 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Hw *string `json:"HW,omitempty"`
}

func (o IpReputationMapResponseBodyLocale) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpReputationMapResponseBodyLocale struct{}"
	}

	return strings.Join([]string{"IpReputationMapResponseBodyLocale", string(data)}, " ")
}
