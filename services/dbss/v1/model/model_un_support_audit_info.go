package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UnSupportAuditInfo struct {

	// 实例版本
	AuditVersion *string `json:"audit_version,omitempty"`

	// 实例名称
	ServerName *string `json:"server_name,omitempty"`
}

func (o UnSupportAuditInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnSupportAuditInfo struct{}"
	}

	return strings.Join([]string{"UnSupportAuditInfo", string(data)}, " ")
}
