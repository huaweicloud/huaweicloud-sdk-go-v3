package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeTypes 节点类型对象
type NodeTypes struct {

	// **参数解释**： 规格名称。 **取值范围**： 字母、数字、小数点、下划线、短横线。
	SpecName string `json:"spec_name"`

	// **参数解释**： 规格详细信息。 **取值范围**： 不涉及。
	Detail []Detail `json:"detail"`

	// **参数解释**： 规格ID。 **取值范围**： 一般为UUID。
	Id string `json:"id"`

	// **参数解释**： 产品类型。 **取值范围**： - dws：云数仓。 - hybrid：实时数仓。 - stream：IoT数仓。
	DatastoreType string `json:"datastore_type"`

	// **参数解释**： 架构类型。 **取值范围**： - x86； - arm；
	Architecture string `json:"architecture"`

	// **参数解释**： 支持的可用区及状态信息。 **取值范围**： 不涉及。
	AvailableZones []NodeTypeAvailableZones `json:"available_zones"`

	// **参数解释**： 内存大小。单位：GB。 **取值范围**： 大于0的正整数。
	Ram int32 `json:"ram"`

	// **参数解释**： CPU数量。 **取值范围**： 大于0的正整数。
	Vcpus int32 `json:"vcpus"`

	// **参数解释**： 内核版本信息。 **取值范围**： 不涉及。
	Datastores []NodeTypeDatastores `json:"datastores"`

	Volume *VolumeResp `json:"volume"`

	// **参数解释**： 弹性弹性容量规格的规格容量信息 **取值范围**： 如果规格为弹性容量规格，则该属性为规格典配的弹性容量信息，包括存储类型、最小容量、最大容量以及步长信息，如果为固定存储规格，则该属性为null。
	ElasticVolumeSpecs []NodeTypeElasticVolumeSpecs `json:"elastic_volume_specs"`
}

func (o NodeTypes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeTypes struct{}"
	}

	return strings.Join([]string{"NodeTypes", string(data)}, " ")
}
