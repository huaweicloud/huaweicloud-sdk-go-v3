package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AuditInstanceRiskCount struct {

	// 风险数量
	Count *int64 `json:"count,omitempty"`

	// 数据库ID
	DbId *string `json:"db_id,omitempty"`

	// 数据库IP
	DbIp *string `json:"db_ip,omitempty"`

	// 数据库名称
	DbName *string `json:"db_name,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例名称
	InstanceName *string `json:"instance_name,omitempty"`
}

func (o AuditInstanceRiskCount) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuditInstanceRiskCount struct{}"
	}

	return strings.Join([]string{"AuditInstanceRiskCount", string(data)}, " ")
}
