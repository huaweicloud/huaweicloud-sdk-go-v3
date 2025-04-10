package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteCrossCloudDisasterEndSimulationResponse Response Object
type ExecuteCrossCloudDisasterEndSimulationResponse struct {

	// 任务id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteCrossCloudDisasterEndSimulationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteCrossCloudDisasterEndSimulationResponse struct{}"
	}

	return strings.Join([]string{"ExecuteCrossCloudDisasterEndSimulationResponse", string(data)}, " ")
}
