package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteServiceLinkedAgencyV5Response Response Object
type DeleteServiceLinkedAgencyV5Response struct {

	// 删除任务ID。
	DeletionTaskId *string `json:"deletion_task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteServiceLinkedAgencyV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServiceLinkedAgencyV5Response struct{}"
	}

	return strings.Join([]string{"DeleteServiceLinkedAgencyV5Response", string(data)}, " ")
}
