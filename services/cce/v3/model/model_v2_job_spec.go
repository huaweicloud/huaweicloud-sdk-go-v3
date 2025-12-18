package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type V2JobSpec struct {

	// **参数解释**： Job 类型 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Type *string `json:"type,omitempty"`

	// **参数解释**： 集群ID **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Clusteruid *string `json:"clusteruid,omitempty"`

	// **参数解释**： 资源ID **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Resourceid *string `json:"resourceid,omitempty"`

	// **参数解释**： 资源名称 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Resourcename *string `json:"resourcename,omitempty"`

	// **参数解释**： Job的扩容参数 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Extendparam map[string]string `json:"extendparam,omitempty"`

	// **参数解释**： 子Job详情列表 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Subjobs *[]V2Job `json:"subjobs,omitempty"`
}

func (o V2JobSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "V2JobSpec struct{}"
	}

	return strings.Join([]string{"V2JobSpec", string(data)}, " ")
}
