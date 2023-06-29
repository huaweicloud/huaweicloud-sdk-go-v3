package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeTypes 节点类型对象
type NodeTypes struct {

	// 节点类型名称。
	SpecName string `json:"spec_name"`

	// 节点类型详细。
	Detail []Detail `json:"detail"`

	// 节点类型ID。
	Id string `json:"id"`

	// 产品类型 - dws：云数仓。 - hybrid：实时数仓。 - stream：IoT数仓。
	DatastoreType string `json:"datastore_type"`

	// 支持的可用区及状态信息。
	AvailableZones []NodeTypeAvailableZones `json:"available_zones"`

	// 内存大小。
	Ram int32 `json:"ram"`

	// CPU数量。
	Vcpus int32 `json:"vcpus"`

	// 内核版本信息。
	Datastores []NodeTypeDatastores `json:"datastores"`

	Volume *VolumeResp `json:"volume"`

	// 如果规格为弹性容量规格，则该属性为规格典配的弹性容量信息，包括存储类型、最小容量、最大容量以及步长信息，如果为固定存储规格，则该属性为null。
	ElasticVolumeSpecs []NodeTypeElasticVolumeSpecs `json:"elastic_volume_specs"`
}

func (o NodeTypes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeTypes struct{}"
	}

	return strings.Join([]string{"NodeTypes", string(data)}, " ")
}
