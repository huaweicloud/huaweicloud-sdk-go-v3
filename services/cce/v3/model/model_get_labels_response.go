package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetLabelsResponse Response Object
type GetLabelsResponse struct {

	// **参数解释**： API类型 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： Labels
	Kind *string `json:"kind,omitempty"`

	// **参数解释**： API版本 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： v3
	ApiVersion *string `json:"apiVersion,omitempty"`

	// **参数解释：** 节点标签，按节点池分类。 - key表示节点池ID，默认节点池取值为\"DefaultPool\"。 - value表示标签，key/value对格式。其中key为标签的键，value为标签键对应的值列表。  **约束限制：** 不涉及
	Spec           map[string]map[string][]string `json:"spec,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o GetLabelsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetLabelsResponse struct{}"
	}

	return strings.Join([]string{"GetLabelsResponse", string(data)}, " ")
}
