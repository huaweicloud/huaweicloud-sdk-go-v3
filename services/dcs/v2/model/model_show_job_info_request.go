package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobInfoRequest Request Object
type ShowJobInfoRequest struct {

	// Job任务ID。
	JobId string `json:"job_id"`
}

func (o ShowJobInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowJobInfoRequest", string(data)}, " ")
}
