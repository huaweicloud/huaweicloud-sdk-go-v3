package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteClickHouseInstanceRequest Request Object
type DeleteClickHouseInstanceRequest struct {

	// GaussDB(for MySQL)实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	// ClickHouse实例ID，严格匹配UUID规则。
	ClickhouseInstanceId string `json:"clickhouse_instance_id"`

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o DeleteClickHouseInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteClickHouseInstanceRequest struct{}"
	}

	return strings.Join([]string{"DeleteClickHouseInstanceRequest", string(data)}, " ")
}
