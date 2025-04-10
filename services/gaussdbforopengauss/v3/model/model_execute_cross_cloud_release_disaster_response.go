package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteCrossCloudReleaseDisasterResponse Response Object
type ExecuteCrossCloudReleaseDisasterResponse struct {

	// 任务id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteCrossCloudReleaseDisasterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteCrossCloudReleaseDisasterResponse struct{}"
	}

	return strings.Join([]string{"ExecuteCrossCloudReleaseDisasterResponse", string(data)}, " ")
}
