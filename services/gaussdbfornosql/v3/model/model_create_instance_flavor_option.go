package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceFlavorOption 实例规格详情。
type CreateInstanceFlavorOption struct {

	// **参数解释：** 节点数量。 **约束限制：** 不涉及。 **取值范围：** - GeminiDB Cassandra实例的节点数量可取3~60。 - GeminiDB Mongo 4.0版本副本集实例的节点数量为3。 - GeminiDB Influx集群实例的节点数量可取3~16。 - GeminiDB Influx单节点实例的节点数量可取1。 - GeminiDB Redis实例的节点数量可取3~12。 **默认取值：** 不涉及。
	Num string `json:"num"`

	// **参数解释：** 磁盘类型。 **约束限制：** 不涉及。 **取值范围：** 取值为“ULTRAHIGH”，表示SSD盘。 **默认取值：** 不涉及。
	Storage string `json:"storage"`

	// **参数解释：** 磁盘大小。必须为整数，单位为GB。 **约束限制：** GeminiDB Cassandra，GeminiDB Mongo，GeminiDB Influx的最小磁盘容量100GB，最大磁盘容量与实例的性能规格有关。GeminiDB Redis的最大和最小磁盘容量与节点数和实例的性能规格有关。 **取值范围：** - GeminiDB Cassandra请参见数据库实例规格。 - GeminiDB Mongo请参见数据库实例规格。 - GeminiDB Influx请参见数据库实例规格。 - GeminiDB Redis请参见数据库实例规格。 **默认取值：** 不涉及。
	Size string `json:"size"`

	// -| **参数解释：** 资源规格编码。 **约束限制：** 不涉及。 **取值范围：** 获取方法请参见查询数据库规格 - QueryingInstanceSpecifications中响应参数“spec_code”的值。 **默认取值：** 不涉及。
	SpecCode string `json:"spec_code"`
}

func (o CreateInstanceFlavorOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceFlavorOption struct{}"
	}

	return strings.Join([]string{"CreateInstanceFlavorOption", string(data)}, " ")
}
