package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchChangeNodeToPeriodReqBody struct {

	// **参数解释**： API类型 **约束限制**： 该值不可修改 **取值范围**： 不涉及 **默认取值**： Node
	Kind *string `json:"kind,omitempty"`

	// **参数解释**： API版本 **约束限制**： 该值不可修改 **取值范围**： 不涉及 **默认取值**： v3
	ApiVersion *string `json:"apiVersion,omitempty"`

	// **参数解释**： 要进行按需转包的CCE节点ID列表，示例如下： ``` \"nodeList\": [\"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx\", \"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx\"] ``` **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	NodeList []string `json:"nodeList"`

	PeriodOrderParam *PeriodOrderParam `json:"periodOrderParam"`

	// **参数解释**： 需要一起转包周期的资源类型列表，示例如下： ``` \"includeResources\": [\"eip\"] ``` **约束限制**： 当前仅支持eip（弹性公网IP）资源 **取值范围**： - \"eip\"：弹性公网IP **默认取值**： 不涉及
	IncludeResources *[]string `json:"includeResources,omitempty"`
}

func (o BatchChangeNodeToPeriodReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchChangeNodeToPeriodReqBody struct{}"
	}

	return strings.Join([]string{"BatchChangeNodeToPeriodReqBody", string(data)}, " ")
}
