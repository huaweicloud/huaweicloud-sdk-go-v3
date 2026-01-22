package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteAddressSetsDto struct {

	// **参数解释**： 防护对象ID，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)获得，通过返回值中的data.records.protect_objects.object_id（.表示各对象之间层级的区分）获得，注意type为0的为互联网边界防护对象id，type为1的为VPC边界防护对象id。 **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	ObjectId string `json:"object_id"`

	// **参数解释**： 地址组ID，可以通过调用[查询地址组列表接口]获得，通过返回值中的data.records.set_id获得 **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	SetIds []string `json:"set_ids"`
}

func (o BatchDeleteAddressSetsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAddressSetsDto struct{}"
	}

	return strings.Join([]string{"BatchDeleteAddressSetsDto", string(data)}, " ")
}
