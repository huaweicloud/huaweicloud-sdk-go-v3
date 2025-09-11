package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AuditScopeRequestNew struct {

	// 审计范围ID列表
	Ids []string `json:"ids"`
}

func (o AuditScopeRequestNew) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuditScopeRequestNew struct{}"
	}

	return strings.Join([]string{"AuditScopeRequestNew", string(data)}, " ")
}
