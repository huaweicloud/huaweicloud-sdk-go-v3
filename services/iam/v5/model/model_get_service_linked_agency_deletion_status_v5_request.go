package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetServiceLinkedAgencyDeletionStatusV5Request Request Object
type GetServiceLinkedAgencyDeletionStatusV5Request struct {

	// 删除任务ID。
	DeletionTaskId string `json:"deletion_task_id"`
}

func (o GetServiceLinkedAgencyDeletionStatusV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetServiceLinkedAgencyDeletionStatusV5Request struct{}"
	}

	return strings.Join([]string{"GetServiceLinkedAgencyDeletionStatusV5Request", string(data)}, " ")
}
