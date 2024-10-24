package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDdmFlavorsRequest Request Object
type ListDdmFlavorsRequest struct {

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0。取值必须为数字，且不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询个数上限值。取值范围：1~128。不传该参数时，默认值为10。
	Limit *int32 `json:"limit,omitempty"`

	// 引擎ID,通过查询DDM引擎信息接口获取，引擎ID与引擎版本至少指定一个
	EngineId *string `json:"engine_id,omitempty"`

	// 引擎版本,通过查询DDM引擎信息接口获取，引擎ID与引擎版本至少指定一个
	EngineVersion *string `json:"engine_version,omitempty"`

	// 可用区，多个用\",\"分割，如cn-southwest-244a,cn-southwest-244b。请参见地区和终端节点(https://console.huaweicloud.com/apiexplorer/#/endpoint/DDM)。
	AvailableZones *string `json:"available_zones,omitempty"`
}

func (o ListDdmFlavorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDdmFlavorsRequest struct{}"
	}

	return strings.Join([]string{"ListDdmFlavorsRequest", string(data)}, " ")
}
