package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopAuditInstanceNewRequest Request Object
type StopAuditInstanceNewRequest struct {
	Body *ServerIdBean `json:"body,omitempty"`
}

func (o StopAuditInstanceNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopAuditInstanceNewRequest struct{}"
	}

	return strings.Join([]string{"StopAuditInstanceNewRequest", string(data)}, " ")
}
