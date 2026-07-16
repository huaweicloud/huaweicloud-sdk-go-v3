package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPodIdentityAssociationResponse Response Object
type ShowPodIdentityAssociationResponse struct {

	// **参数解释：** pod-identity关联的uid。 **约束限制：** 该值不可修改 **取值范围：** 不涉及 **默认取值：** 无
	Uid *string `json:"uid,omitempty"`

	// **参数解释：** pod-identity关联所属的集群id。 **约束限制：** 该值不可修改 **取值范围：** 不涉及 **默认取值：** 无
	ClusterId *string `json:"clusterId,omitempty"`

	// **参数解释：** pod-identity关联所要绑定的serviceaccount所属的命名空间。 **约束限制：** 该值不可修改 **取值范围：** 不涉及 **默认取值：** 无
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释：** pod-identity关联所要绑定的serviceaccount名称。 **约束限制：** 该值不可修改 **取值范围：** 不涉及 **默认取值：** 无
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// **参数解释：** pod-identity关联所要绑定的委托名称。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 无
	AgencyName *string `json:"agencyName,omitempty"`

	// **参数解释：** pod-identity关联的资源标签列表。 **约束限制：** 不涉及
	Tags *[]ResourceTag `json:"tags,omitempty"`

	// **参数解释：** pod-identity关联创建时间。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 无
	CreatedAt *string `json:"createdAt,omitempty"`

	// **参数解释：** pod-identity关联最近一次更新时间。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 无
	ModifiedAt     *string `json:"modifiedAt,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPodIdentityAssociationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPodIdentityAssociationResponse struct{}"
	}

	return strings.Join([]string{"ShowPodIdentityAssociationResponse", string(data)}, " ")
}
