package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFlavorsResult 实例规格信息列表。
type ListFlavorsResult struct {

	// 引擎名称。
	EngineName string `json:"engine_name"`

	// 引擎版本。
	EngineVersion string `json:"engine_version"`

	// CPU核数。
	Vcpus string `json:"vcpus"`

	// 内存大小，单位为兆字节。
	Ram string `json:"ram"`

	// 资源规格编码。例如：geminidb.cassandra.8xlarge.4   - “geminidb.cassandra”表示云数据库GaussDB NoSQL的Cassandra数据库产品。   - “8xlarge.4”表示节点性能规格。
	SpecCode string `json:"spec_code"`

	// 支持该规格的可用区ID。   - 该字段已废弃，请不要使用。
	AvailabilityZone []string `json:"availability_zone"`

	// 规格在可用区内的状态，包含以下状态：   - normal，在售。   - unsupported，暂不支持该规格。   - sellout，售罄。
	AzStatus *interface{} `json:"az_status"`
}

func (o ListFlavorsResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlavorsResult struct{}"
	}

	return strings.Join([]string{"ListFlavorsResult", string(data)}, " ")
}
