package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConsistencyResultRequest Request Object
type ShowConsistencyResultRequest struct {

	// 任务id
	TaskId string `json:"task_id"`
}

func (o ShowConsistencyResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConsistencyResultRequest struct{}"
	}

	return strings.Join([]string{"ShowConsistencyResultRequest", string(data)}, " ")
}
