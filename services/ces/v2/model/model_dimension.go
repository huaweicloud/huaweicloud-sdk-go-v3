package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Dimension **参数解释**： 资源信息。 **取值范围**： 不涉及。
type Dimension struct {

	// **参数解释**： 资源维度名称，如：弹性云服务器，则维度为instance_id。各服务资源的资源维度名称可查看：“[服务维度名称](ces_03_0059.xml)”。 **约束限制**： 不涉及。 **取值范围**： 以字母开头，只能包含字母、数字、“_”、“-”。长度为[1,32]个字符。 **默认取值**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 资源维度值，为资源的实例ID，如：4270ff17-aba3-4138-89fa-820594c39755。 **约束限制**： 不涉及。 **取值范围**： 长度为[1,256]个字符。        **默认取值**： 不涉及。
	Value *string `json:"value,omitempty"`
}

func (o Dimension) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Dimension struct{}"
	}

	return strings.Join([]string{"Dimension", string(data)}, " ")
}
