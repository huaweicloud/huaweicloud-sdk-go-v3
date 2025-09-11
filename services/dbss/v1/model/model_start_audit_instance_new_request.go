package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartAuditInstanceNewRequest Request Object
type StartAuditInstanceNewRequest struct {
	Body *ServerIdBean `json:"body,omitempty"`
}

func (o StartAuditInstanceNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartAuditInstanceNewRequest struct{}"
	}

	return strings.Join([]string{"StartAuditInstanceNewRequest", string(data)}, " ")
}
