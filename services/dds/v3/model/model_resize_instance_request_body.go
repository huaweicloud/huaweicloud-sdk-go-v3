package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResizeInstanceRequestBody struct {
	Resize *ResizeInstanceOption `json:"resize"`

	// **参数解释：** 变更包年包月实例规格时可指定，表示是否自动从账户中支付，此字段不影响自动续订的支付方式。 **约束限制：** 对于降低规格场景，该字段无效。 **取值范围：** 对于扩大规格场景： - true，表示自动从账户中支付。 - false，表示手动从账户中支付，默认为该方式。 **默认取值：** false。
	IsAutoPay *bool `json:"is_auto_pay,omitempty"`

	// **参数解释：** 标识是否进行强制规格变更操作。 **约束限制：** 集群的dds mongos节点和只读节点不支持强制规格变更。 **取值范围：** 对于规格变更： - true，表示执行强制规格变更。 - 不传此参数，表示执行正常规格变更。 **默认取值：** 不传此参数。
	IsForceResize *bool `json:"is_force_resize,omitempty"`
}

func (o ResizeInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"ResizeInstanceRequestBody", string(data)}, " ")
}
