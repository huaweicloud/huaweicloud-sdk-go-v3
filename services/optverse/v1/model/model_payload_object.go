package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PayloadObject **参数解释**： 具体返回信息。 **约束限制**： 不涉及 **取值范围**： 不涉及。 **默认取值**： 不涉及
type PayloadObject struct {

	// **参数解释**： 返回信息列表。 **约束限制**： 不涉及 **取值范围**： 元素数量范围[0,100000000]。 **默认取值**： 不涉及
	List *[]interface{} `json:"list,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// **参数解释**： 返回对象信息。 **约束限制**： 不涉及 **取值范围**： 不涉及。 **默认取值**： 不涉及
	Item *interface{} `json:"item,omitempty"`
}

func (o PayloadObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PayloadObject struct{}"
	}

	return strings.Join([]string{"PayloadObject", string(data)}, " ")
}
