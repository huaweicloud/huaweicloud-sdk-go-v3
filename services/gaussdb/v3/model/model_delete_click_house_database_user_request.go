package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteClickHouseDatabaseUserRequest Request Object
type DeleteClickHouseDatabaseUserRequest struct {

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`

	// ClickHouse实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	// 数据库名称。
	UserName string `json:"user_name"`
}

func (o DeleteClickHouseDatabaseUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteClickHouseDatabaseUserRequest struct{}"
	}

	return strings.Join([]string{"DeleteClickHouseDatabaseUserRequest", string(data)}, " ")
}
