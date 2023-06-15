package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDatabaseAuthorityRequest struct {

	// cluster_id
	ClusterId string `json:"cluster_id"`

	// 对象类型 [DATABASE | SCHEMA | TABLE | VIEW | COLUMN| FUNCTION|| SEQUENCE| NODEGROUP]
	Type string `json:"type"`

	// 对象名称
	Name []string `json:"name"`

	// 数据库名
	Database string `json:"database"`

	// 模式名，对象类型为TABLE、VIEW、COLUMN、FUNCTION、SEQUENCE时必选
	Schema *string `json:"schema,omitempty"`

	// 表名，对象类型为COLUMN时必选
	Table *string `json:"table,omitempty"`
}

func (o ShowDatabaseAuthorityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDatabaseAuthorityRequest struct{}"
	}

	return strings.Join([]string{"ShowDatabaseAuthorityRequest", string(data)}, " ")
}
