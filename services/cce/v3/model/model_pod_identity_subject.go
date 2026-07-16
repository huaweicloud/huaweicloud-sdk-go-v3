package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PodIdentitySubject 委托凭据所属的ServiceAccount归属信息
type PodIdentitySubject struct {

	// **参数解释**： ServiceAccount所属的命名空间 **约束限制**： 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释：** ServiceAccount名称 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ServiceAccount *string `json:"serviceAccount,omitempty"`
}

func (o PodIdentitySubject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PodIdentitySubject struct{}"
	}

	return strings.Join([]string{"PodIdentitySubject", string(data)}, " ")
}
