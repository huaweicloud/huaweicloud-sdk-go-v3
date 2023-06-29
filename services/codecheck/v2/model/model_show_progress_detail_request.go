package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProgressDetailRequest Request Object
type ShowProgressDetailRequest struct {

	// 任务ID
	TaskId string `json:"task_id"`
}

func (o ShowProgressDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProgressDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowProgressDetailRequest", string(data)}, " ")
}
