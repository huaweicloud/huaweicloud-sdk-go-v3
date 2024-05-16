package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSecurityAssignedQueueRequest Request Object
type CreateSecurityAssignedQueueRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *QueueSrcAssignCreateDto `json:"body,omitempty"`
}

func (o CreateSecurityAssignedQueueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityAssignedQueueRequest struct{}"
	}

	return strings.Join([]string{"CreateSecurityAssignedQueueRequest", string(data)}, " ")
}
