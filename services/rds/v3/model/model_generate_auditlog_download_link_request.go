/*
 * RDS
 *
 * API v3
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type GenerateAuditlogDownloadLinkRequest struct {
	// 审计日志ID列表，限制50条以内。
	Ids []string `json:"ids"`
}

func (o GenerateAuditlogDownloadLinkRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"GenerateAuditlogDownloadLinkRequest", string(data)}, " ")
}
