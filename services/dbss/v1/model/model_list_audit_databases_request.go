package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAuditDatabasesRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 实例状态 ON ：开启 OFF ： 关闭
	Status *string `json:"status,omitempty"`

	// 偏移量
	Offset *string `json:"offset,omitempty"`

	// 查询记录数
	Limit *string `json:"limit,omitempty"`
}

func (o ListAuditDatabasesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditDatabasesRequest struct{}"
	}

	return strings.Join([]string{"ListAuditDatabasesRequest", string(data)}, " ")
}
