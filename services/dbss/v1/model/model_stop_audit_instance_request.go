package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopAuditInstanceRequest Request Object
type StopAuditInstanceRequest struct {
	Body *ServerIdBean `json:"body,omitempty"`
}

func (o StopAuditInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopAuditInstanceRequest struct{}"
	}

	return strings.Join([]string{"StopAuditInstanceRequest", string(data)}, " ")
}
