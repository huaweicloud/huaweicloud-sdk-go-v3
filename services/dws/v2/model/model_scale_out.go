package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScaleOut **参数解释**： 扩容集群详情。 **取值范围**： 不涉及。
type ScaleOut struct {

	// **参数解释**： 扩容节点数。 **取值范围**： 大于等于3。
	Count int32 `json:"count"`

	// **参数解释**： 子网ID。 **取值范围**： 同VPC下有效的子网ID。
	SubnetId *string `json:"subnet_id,omitempty"`
}

func (o ScaleOut) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScaleOut struct{}"
	}

	return strings.Join([]string{"ScaleOut", string(data)}, " ")
}
