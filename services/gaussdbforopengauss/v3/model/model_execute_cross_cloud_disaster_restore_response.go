package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteCrossCloudDisasterRestoreResponse Response Object
type ExecuteCrossCloudDisasterRestoreResponse struct {

	// 任务id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteCrossCloudDisasterRestoreResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteCrossCloudDisasterRestoreResponse struct{}"
	}

	return strings.Join([]string{"ExecuteCrossCloudDisasterRestoreResponse", string(data)}, " ")
}
