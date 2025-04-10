package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteCrossCloudDisasterRecoveryFailoverResponse Response Object
type ExecuteCrossCloudDisasterRecoveryFailoverResponse struct {

	// 任务id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteCrossCloudDisasterRecoveryFailoverResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteCrossCloudDisasterRecoveryFailoverResponse struct{}"
	}

	return strings.Join([]string{"ExecuteCrossCloudDisasterRecoveryFailoverResponse", string(data)}, " ")
}
