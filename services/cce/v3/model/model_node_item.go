package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodeItem struct {

	// **参数解释**： 节点ID，节点ID获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Uid string `json:"uid"`
}

func (o NodeItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeItem struct{}"
	}

	return strings.Join([]string{"NodeItem", string(data)}, " ")
}
