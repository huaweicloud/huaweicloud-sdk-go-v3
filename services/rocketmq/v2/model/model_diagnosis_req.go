package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DiagnosisReq struct {

	// **参数解释**： 消费组名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	GroupName string `json:"group_name"`

	// **参数解释**： 节点ID列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	NodeIdList *[]string `json:"node_id_list,omitempty"`
}

func (o DiagnosisReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiagnosisReq struct{}"
	}

	return strings.Join([]string{"DiagnosisReq", string(data)}, " ")
}
