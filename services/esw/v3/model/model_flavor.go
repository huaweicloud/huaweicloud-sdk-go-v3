package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Flavor 配额信息
type Flavor struct {

	// - 参数解释：ESW实例规格名称。 - 约束限制：不涉及。 - 取值范围：  - l2cg.small.ha  - l2cg.medium.ha  - l2cg.large.ha - 默认取值：不涉及。
	Name string `json:"name"`

	// - 参数解释：ESW实例规格的唯一标识。 - 约束限制：不涉及。 - 取值范围：1-3。 - 默认取值：不涉及。
	Id string `json:"id"`

	// - 参数解释：该规格ESW实例可承载的二层连接数量。 - 约束限制：不涉及。 - 取值范围：1、3、6。 - 默认取值：不涉及。
	Connections int32 `json:"connections"`

	// - 参数解释：该规格ESW实例可承载的最大带宽。 - 约束限制：单位:Gbit/s。 - 取值范围：3、5、10。 - 默认取值：不涉及。
	Bandwidth int32 `json:"bandwidth"`

	// - 参数解释：该规格ESW实例可承载的最大发包数。 - 约束限制：不涉及。 - 取值范围：500000、1000000、2000000。 - 默认取值：不涉及。
	Pps int32 `json:"pps"`

	// - 参数解释：可选用该规格的可用区列表。 - 约束限制：不涉及。 - 取值范围：不涉及。 - 默认取值：不涉及。
	AvailableZones []string `json:"available_zones"`
}

func (o Flavor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Flavor struct{}"
	}

	return strings.Join([]string{"Flavor", string(data)}, " ")
}
