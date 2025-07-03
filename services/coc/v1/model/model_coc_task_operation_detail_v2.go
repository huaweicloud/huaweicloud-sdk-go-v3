package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CocTaskOperationDetailV2 struct {

	// 任务ID
	TaskId *string `json:"task_id,omitempty"`

	// 任务KEY
	Key *string `json:"key,omitempty"`
}

func (o CocTaskOperationDetailV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CocTaskOperationDetailV2 struct{}"
	}

	return strings.Join([]string{"CocTaskOperationDetailV2", string(data)}, " ")
}
