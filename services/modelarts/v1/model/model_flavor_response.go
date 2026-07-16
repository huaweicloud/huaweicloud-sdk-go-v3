package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlavorResponse 训练作业、算法的规格信息。
type FlavorResponse struct {

	// **参数解释**：训练作业选择的资源池ID。 **取值范围**：不涉及。
	PoolId *string `json:"pool_id,omitempty"`

	// 资源规格的ID。
	FlavorId *string `json:"flavor_id,omitempty"`

	// 资源规格的名称。
	FlavorName *string `json:"flavor_name,omitempty"`

	// 资源规格的最大节点数。
	MaxNum *int32 `json:"max_num,omitempty"`

	// 资源规格的类型。可选值如下： - CPU - GPU - [Ascend](tag:hc,hk,fcs_super)
	FlavorType *string `json:"flavor_type,omitempty"`

	Billing *BillingInfo `json:"billing,omitempty"`

	FlavorInfo *FlavorInfoResponse `json:"flavor_info,omitempty"`

	// 其他规格属性。
	Attributes map[string]string `json:"attributes,omitempty"`
}

func (o FlavorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorResponse struct{}"
	}

	return strings.Join([]string{"FlavorResponse", string(data)}, " ")
}
