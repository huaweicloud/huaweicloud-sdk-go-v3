package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UnsupportedAuditInfoResponse struct {

	// 审计信息
	AuditInfos *[]UnsupportedAuditInfo `json:"audit_infos,omitempty"`

	// 支持的版本
	SupportVersion *string `json:"support_version,omitempty"`
}

func (o UnsupportedAuditInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnsupportedAuditInfoResponse struct{}"
	}

	return strings.Join([]string{"UnsupportedAuditInfoResponse", string(data)}, " ")
}
