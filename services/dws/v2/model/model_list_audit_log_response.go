package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditLogResponse Response Object
type ListAuditLogResponse struct {

	// **参数解释**： 审计日志列表。 **取值范围**： 不涉及。
	Records *[]AuditDumpRecord `json:"records,omitempty"`

	// **参数解释**： 集群ID。 **取值范围**： 36位UUID。
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**： 总数。 **取值范围**： 不涉及。
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
