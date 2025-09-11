package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RebootAuditInstanceNewRequest Request Object
type RebootAuditInstanceNewRequest struct {
	Body *ServerIdBean `json:"body,omitempty"`
}

func (o RebootAuditInstanceNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RebootAuditInstanceNewRequest struct{}"
	}

	return strings.Join([]string{"RebootAuditInstanceNewRequest", string(data)}, " ")
}
