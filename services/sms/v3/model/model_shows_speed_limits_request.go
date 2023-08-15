package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowsSpeedLimitsRequest Request Object
type ShowsSpeedLimitsRequest struct {

	// 查询限速信息的任务ID
	TaskId string `json:"task_id"`
}

func (o ShowsSpeedLimitsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowsSpeedLimitsRequest struct{}"
	}

	return strings.Join([]string{"ShowsSpeedLimitsRequest", string(data)}, " ")
}
