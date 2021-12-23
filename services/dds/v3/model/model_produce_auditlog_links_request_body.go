package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProduceAuditlogLinksRequestBody struct {
	// 审计日志ID列表，限制50条以内。

	Ids []string `json:"ids"`
}

func (o ProduceAuditlogLinksRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProduceAuditlogLinksRequestBody struct{}"
	}

	return strings.Join([]string{"ProduceAuditlogLinksRequestBody", string(data)}, " ")
}
