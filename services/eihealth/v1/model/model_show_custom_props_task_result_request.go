package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCustomPropsTaskResultRequest Request Object
type ShowCustomPropsTaskResultRequest struct {

	// 自定义属性任务ID
	TaskId string `json:"task_id"`
}

func (o ShowCustomPropsTaskResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCustomPropsTaskResultRequest struct{}"
	}

	return strings.Join([]string{"ShowCustomPropsTaskResultRequest", string(data)}, " ")
}
