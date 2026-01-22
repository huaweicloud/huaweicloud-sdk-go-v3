package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteDomainSetsDto struct {

	// **参数解释**： 防护对象ID，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志ID，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)获得，通过返回值中的data.records.protect_objects.object_id（.表示各对象之间层级的区分）获得，注意type为0的为互联网边界防护对象ID，type为1的为VPC边界防护对象ID。 **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	ObjectId *string `json:"object_id,omitempty"`

	// **参数解释**： 域名组ID列表，域名组ID，可通过[查询域名组列表接口](ListDomainSets.xml)查询获得，通过返回值中的data.records.set_id（.表示各对象之间层级的区分）获得。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	SetIds *[]string `json:"set_ids,omitempty"`
}

func (o BatchDeleteDomainSetsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteDomainSetsDto struct{}"
	}

	return strings.Join([]string{"BatchDeleteDomainSetsDto", string(data)}, " ")
}
