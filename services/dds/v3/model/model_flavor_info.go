package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlavorInfo 规格信息。
type FlavorInfo struct {

	// 引擎名称。
	EngineName string `json:"engine_name"`

	// 节点类型。文档数据库包含以下几种节点类型： - mongos - shard - config - replica - single - readonly
	Type string `json:"type"`

	// CPU核数。
	Vcpus string `json:"vcpus"`

	// 内存大小，单位为兆字节。
	Ram string `json:"ram"`

	// 资源规格编码。例如：dds.mongodb.c6.xlarge.2.shard。  - “dds”表示文档数据库服务产品。 - “c6.xlarge.2”表示节点性能规格，为高内存类型。 - “shard”表示节点类型。
	SpecCode string `json:"spec_code"`

	// '支持该规格的可用区ID。' 示例：[\"cn-east-2a\",\"cn-east-2b\",\"cn-east-2c\"]。
	AzStatus *interface{} `json:"az_status"`

	// 数据库版本号列表。针对DDS引擎的mongos节点，例如：{\"3.4\", \"4.0\"}
	EngineVersions []string `json:"engine_versions"`
}

func (o FlavorInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorInfo struct{}"
	}

	return strings.Join([]string{"FlavorInfo", string(data)}, " ")
}
