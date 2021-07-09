package model

import (
	"encoding/json"

	"strings"
)

type MysqlProxyComputeFlavor struct {
	// Proxy规格id。

	Id *string `json:"id,omitempty"`
	// Proxy规格码。

	SpecCode *string `json:"spec_code,omitempty"`
	// CPU大小。例如：1表示1U。

	Vcpus *string `json:"vcpus,omitempty"`
	// 内存大小，单位为GB。

	Ram *string `json:"ram,omitempty"`
	// 数据库类型。

	DbType *string `json:"db_type,omitempty"`
	// 其中key是可用区编号，value是规格所在az的状态。

	AzStatus *interface{} `json:"az_status,omitempty"`
	// Region状态。

	RegionStatus *string `json:"region_status,omitempty"`
}

func (o MysqlProxyComputeFlavor) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MysqlProxyComputeFlavor struct{}"
	}

	return strings.Join([]string{"MysqlProxyComputeFlavor", string(data)}, " ")
}
