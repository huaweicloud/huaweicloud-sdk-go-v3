package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteCrossCloudDisasterStartSimulationResponse Response Object
type ExecuteCrossCloudDisasterStartSimulationResponse struct {

	// 任务id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteCrossCloudDisasterStartSimulationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteCrossCloudDisasterStartSimulationResponse struct{}"
	}

	return strings.Join([]string{"ExecuteCrossCloudDisasterStartSimulationResponse", string(data)}, " ")
}
