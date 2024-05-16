package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSecurityAssignedQueueResponse Response Object
type CreateSecurityAssignedQueueResponse struct {
	Body           *[]QueueSrcAssignEntity `json:"body,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o CreateSecurityAssignedQueueResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityAssignedQueueResponse struct{}"
	}

	return strings.Join([]string{"CreateSecurityAssignedQueueResponse", string(data)}, " ")
}
