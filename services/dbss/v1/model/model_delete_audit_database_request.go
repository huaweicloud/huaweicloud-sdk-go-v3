package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAuditDatabaseRequest Request Object
type DeleteAuditDatabaseRequest struct {

	// 实例ID。可在查询实例列表接口的ID字段获取。
	InstanceId string `json:"instance_id"`

	// 数据库ID，可在查询数据库列表接口ID字段获取。
	DbId string `json:"db_id"`
}

func (o DeleteAuditDatabaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAuditDatabaseRequest struct{}"
	}

	return strings.Join([]string{"DeleteAuditDatabaseRequest", string(data)}, " ")
}
