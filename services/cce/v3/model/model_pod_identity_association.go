package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PodIdentityAssociation pod-identity关联创建请求参数
type PodIdentityAssociation struct {

	// **参数解释：** pod-identity关联所要绑定的serviceaccount所属的命名空间。 **约束限制：** 该值不可修改 **取值范围：** 不涉及 **默认取值：** 无
	Namespace string `json:"namespace"`

	// **参数解释：** pod-identity关联所要绑定的serviceaccount名称。 **约束限制：** 同一个serviceaccount最多创建一条pod-identity关联记录，不支持创建多个 **取值范围：** 不涉及 **默认取值：** 无
	ServiceAccount string `json:"serviceAccount"`

	// **参数解释：** pod-identity关联所要绑定的委托名称，委托可以是一般委托或信任委托。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 无
	AgencyName string `json:"agencyName"`

	// **参数解释：** pod-identity关联的资源标签列表。 **约束限制：** 不涉及
	Tags *[]ResourceTag `json:"tags,omitempty"`
}

func (o PodIdentityAssociation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PodIdentityAssociation struct{}"
	}

	return strings.Join([]string{"PodIdentityAssociation", string(data)}, " ")
}
