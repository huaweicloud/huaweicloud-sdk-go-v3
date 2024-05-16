package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecurityAssignedQueueRequest Request Object
type UpdateSecurityAssignedQueueRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 分配给当前空间的队列资源id。
	Id string `json:"id"`

	Body *QueueSrcAssignUpdateDto `json:"body,omitempty"`
}

func (o UpdateSecurityAssignedQueueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityAssignedQueueRequest struct{}"
	}

	return strings.Join([]string{"UpdateSecurityAssignedQueueRequest", string(data)}, " ")
}
