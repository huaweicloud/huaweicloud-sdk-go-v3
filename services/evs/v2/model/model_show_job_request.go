package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobRequest Request Object
type ShowJobRequest struct {

	// job ID。
	JobId string `json:"job_id"`
}

func (o ShowJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobRequest struct{}"
	}

	return strings.Join([]string{"ShowJobRequest", string(data)}, " ")
}
