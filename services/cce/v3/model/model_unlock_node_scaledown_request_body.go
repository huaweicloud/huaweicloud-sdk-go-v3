package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UnlockNodeScaledownRequestBody struct {

	// **参数解释**： API版本 **约束限制**： 固定值，不允许修改 **取值范围**： 不涉及 **默认取值**： v3
	ApiVersion string `json:"apiVersion"`

	// **参数解释**： API类型 **约束限制**： 固定值，不允许修改 **取值范围**： 不涉及 **默认取值**： List
	Kind string `json:"kind"`

	// **参数解释**： 需要关闭缩容保护的节点ID列表，节点ID获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	NodeList *[]string `json:"nodeList,omitempty"`
}

func (o UnlockNodeScaledownRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnlockNodeScaledownRequestBody struct{}"
	}

	return strings.Join([]string{"UnlockNodeScaledownRequestBody", string(data)}, " ")
}
