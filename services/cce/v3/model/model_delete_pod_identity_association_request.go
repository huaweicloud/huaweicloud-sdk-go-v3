package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePodIdentityAssociationRequest Request Object
type DeletePodIdentityAssociationRequest struct {

	// 集群ID，获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。
	ClusterId string `json:"cluster_id"`

	// **参数解释**： Pod-identity关联ID，获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	AssociationId string `json:"association_id"`
}

func (o DeletePodIdentityAssociationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePodIdentityAssociationRequest struct{}"
	}

	return strings.Join([]string{"DeletePodIdentityAssociationRequest", string(data)}, " ")
}
