package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAuditLogResponse Response Object
type DeleteAuditLogResponse struct {

	// 任务ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteAuditLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAuditLogResponse struct{}"
	}

	return strings.Join([]string{"DeleteAuditLogResponse", string(data)}, " ")
}
