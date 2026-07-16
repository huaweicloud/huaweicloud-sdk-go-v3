package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PodIdentityAssociationUpdate pod-identity关联更新请求参数
type PodIdentityAssociationUpdate struct {

	// **参数解释：** pod-identity关联所要绑定的委托名称，委托可以是一般委托或信任委托。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 无
	AgencyName string `json:"agencyName"`
}

func (o PodIdentityAssociationUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PodIdentityAssociationUpdate struct{}"
	}

	return strings.Join([]string{"PodIdentityAssociationUpdate", string(data)}, " ")
}
