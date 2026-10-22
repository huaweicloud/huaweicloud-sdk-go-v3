package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HyperGroup struct {

	// **参数解释**：超节点组id，对于A5超节点为必填字段。 **约束限制**：不涉及。 **取值范围**：必须是UUID格式的字符串。 **默认取值**：不涉及。
	Id string `json:"id"`
}

func (o HyperGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HyperGroup struct{}"
	}

	return strings.Join([]string{"HyperGroup", string(data)}, " ")
}
