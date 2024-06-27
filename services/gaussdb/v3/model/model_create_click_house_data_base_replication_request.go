package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClickHouseDataBaseReplicationRequest Request Object
type CreateClickHouseDataBaseReplicationRequest struct {

	// ClickHouse实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`

	Body *CreateChDatabaseReplicationRequestBody `json:"body,omitempty"`
}

func (o CreateClickHouseDataBaseReplicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClickHouseDataBaseReplicationRequest struct{}"
	}

	return strings.Join([]string{"CreateClickHouseDataBaseReplicationRequest", string(data)}, " ")
}
