package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletionTaskId 删除任务ID。
type DeletionTaskId struct {
}

func (o DeletionTaskId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletionTaskId struct{}"
	}

	return strings.Join([]string{"DeletionTaskId", string(data)}, " ")
}
