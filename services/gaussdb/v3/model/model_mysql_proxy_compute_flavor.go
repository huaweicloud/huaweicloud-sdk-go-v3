package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MysqlProxyComputeFlavor struct {

	// CPU大小。例如：1表示1U。
	Vcpus string `json:"vcpus" xml:"vcpus"`

	// 内存大小，单位为GB。
	Ram string `json:"ram" xml:"ram"`

	// 数据库类型。
	DbType string `json:"db_type" xml:"db_type"`

	// Proxy规格id。
	Id string `json:"id" xml:"id"`

	// Proxy规格码。
	SpecCode string `json:"spec_code" xml:"spec_code"`

	// 其中key是可用区编号，value是规格所在az的状态。
	AzStatus *interface{} `json:"az_status" xml:"az_status"`
}

func (o MysqlProxyComputeFlavor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MysqlProxyComputeFlavor struct{}"
	}

	return strings.Join([]string{"MysqlProxyComputeFlavor", string(data)}, " ")
}
