package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UnsupportedAuditInfo struct {

	// 实例版本
	AuditVersion *string `json:"audit_version,omitempty"`

	// 实例名称
	ServerName *string `json:"server_name,omitempty"`
}

func (o UnsupportedAuditInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnsupportedAuditInfo struct{}"
	}

	return strings.Join([]string{"UnsupportedAuditInfo", string(data)}, " ")
}
