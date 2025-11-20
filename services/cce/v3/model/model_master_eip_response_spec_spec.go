package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MasterEipResponseSpecSpec **参数解释**： 待绑定的弹性IP配置属性 **约束限制**： 不涉及
type MasterEipResponseSpecSpec struct {

	// **参数解释**： 弹性网卡ID **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Id *string `json:"id,omitempty"`

	Eip *EipSpec `json:"eip,omitempty"`

	// **参数解释**： 是否动态创建 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	IsDynamic *bool `json:"IsDynamic,omitempty"`
}

func (o MasterEipResponseSpecSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MasterEipResponseSpecSpec struct{}"
	}

	return strings.Join([]string{"MasterEipResponseSpecSpec", string(data)}, " ")
}
