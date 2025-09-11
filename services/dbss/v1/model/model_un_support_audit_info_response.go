package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UnSupportAuditInfoResponse struct {

	// 审计信息
	AuditInfos *[]UnSupportAuditInfo `json:"audit_infos,omitempty"`

	// 支持的版本
	SupportVersion *string `json:"support_version,omitempty"`
}

func (o UnSupportAuditInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnSupportAuditInfoResponse struct{}"
	}

	return strings.Join([]string{"UnSupportAuditInfoResponse", string(data)}, " ")
}
