package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteCrossCloudDisasterSwitchoverResponse Response Object
type ExecuteCrossCloudDisasterSwitchoverResponse struct {

	// 任务id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteCrossCloudDisasterSwitchoverResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteCrossCloudDisasterSwitchoverResponse struct{}"
	}

	return strings.Join([]string{"ExecuteCrossCloudDisasterSwitchoverResponse", string(data)}, " ")
}
