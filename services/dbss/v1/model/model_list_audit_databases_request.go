package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditDatabasesRequest Request Object
type ListAuditDatabasesRequest struct {

	// 实例ID。可在查询实例列表接口的ID字段获取。
	InstanceId string `json:"instance_id"`

	// 实例状态 - ON :开启 - OFF : 关闭
	Status *string `json:"status,omitempty"`

	// 偏移量，从第一条数据偏移offset条数据后开始查询，默认为0。
	Offset *string `json:"offset,omitempty"`

	// 查询记录数，默认为100。
	Limit *string `json:"limit,omitempty"`
}

func (o ListAuditDatabasesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditDatabasesRequest struct{}"
	}

	return strings.Join([]string{"ListAuditDatabasesRequest", string(data)}, " ")
}
