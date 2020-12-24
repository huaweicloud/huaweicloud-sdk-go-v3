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

// Response Object
type ListAuditlogsResponse struct {
	Auditlogs *[]Auditlog `json:"auditlogs,omitempty"`
	// 总记录数。
	TotalRecord    *int32 `json:"total_record,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAuditlogsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListAuditlogsResponse", string(data)}, " ")
}
