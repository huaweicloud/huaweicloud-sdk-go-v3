package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RebootAuditInstanceRequest Request Object
type RebootAuditInstanceRequest struct {
	Body *ServerIdBean `json:"body,omitempty"`
}

func (o RebootAuditInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RebootAuditInstanceRequest struct{}"
	}

	return strings.Join([]string{"RebootAuditInstanceRequest", string(data)}, " ")
}
