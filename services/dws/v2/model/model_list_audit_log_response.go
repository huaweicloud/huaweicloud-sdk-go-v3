package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAuditLogResponse struct {

	// 审计日志列表。
	Records *[]AuditDumpRecord `json:"records,omitempty"`

	// 项目ID。
	ClusterId *string `json:"cluster_id,omitempty"`

	// 总数。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAuditLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditLogResponse struct{}"
	}

	return strings.Join([]string{"ListAuditLogResponse", string(data)}, " ")
}
