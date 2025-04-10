package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCrossCloudConstructDisasterResponse Response Object
type CreateCrossCloudConstructDisasterResponse struct {

	// 任务id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateCrossCloudConstructDisasterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCrossCloudConstructDisasterResponse struct{}"
	}

	return strings.Join([]string{"CreateCrossCloudConstructDisasterResponse", string(data)}, " ")
}
