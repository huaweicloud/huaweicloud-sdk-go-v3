package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteClickHouseDataBaseReplicationRequest Request Object
type DeleteClickHouseDataBaseReplicationRequest struct {

	// ClickHouse实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	// 数据库名。
	DatabaseName string `json:"database_name"`

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o DeleteClickHouseDataBaseReplicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteClickHouseDataBaseReplicationRequest struct{}"
	}

	return strings.Join([]string{"DeleteClickHouseDataBaseReplicationRequest", string(data)}, " ")
}
