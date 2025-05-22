package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTargetFlavorsRequest Request Object
type ListTargetFlavorsRequest struct {

	// **参数解释**： 集群当前的规格ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	FlavorId string `json:"flavor_id"`

	// **参数解释**： 集群ID。获取方式方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 此参数不传时，可查询所有支持转换的目标规格，但是由于配额等原因，部分规格可能存在售罄无法使用。 传递集群ID时会自动关联此集群所在可用区下的配额充足的目标规格。 **取值范围**： 不涉及。 **默认取值**： null
	ClusterId *string `json:"cluster_id,omitempty"`
}

func (o ListTargetFlavorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTargetFlavorsRequest struct{}"
	}

	return strings.Join([]string{"ListTargetFlavorsRequest", string(data)}, " ")
}
