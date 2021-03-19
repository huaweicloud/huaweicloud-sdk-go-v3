package model

import (
	"encoding/json"

	"strings"
)

type ProduceAuditlogLinksRequestBody struct {
	// 审计日志ID列表，限制50条以内。

	Ids []string `json:"ids"`
}

func (o ProduceAuditlogLinksRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ProduceAuditlogLinksRequestBody struct{}"
	}

	return strings.Join([]string{"ProduceAuditlogLinksRequestBody", string(data)}, " ")
}
