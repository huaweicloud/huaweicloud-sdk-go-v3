package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDistributionResponse Response Object
type CreateDistributionResponse struct {

	// 任务id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDistributionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDistributionResponse struct{}"
	}

	return strings.Join([]string{"CreateDistributionResponse", string(data)}, " ")
}
