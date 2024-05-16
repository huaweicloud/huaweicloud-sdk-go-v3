package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSecurityAssignedQueueRequest Request Object
type DeleteSecurityAssignedQueueRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 需要删除的当前空间队列资源id。
	Id string `json:"id"`
}

func (o DeleteSecurityAssignedQueueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecurityAssignedQueueRequest struct{}"
	}

	return strings.Join([]string{"DeleteSecurityAssignedQueueRequest", string(data)}, " ")
}
