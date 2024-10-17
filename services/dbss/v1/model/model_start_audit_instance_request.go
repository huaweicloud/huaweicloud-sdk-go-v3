package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartAuditInstanceRequest Request Object
type StartAuditInstanceRequest struct {
	Body *ServerIdBean `json:"body,omitempty"`
}

func (o StartAuditInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartAuditInstanceRequest struct{}"
	}

	return strings.Join([]string{"StartAuditInstanceRequest", string(data)}, " ")
}
