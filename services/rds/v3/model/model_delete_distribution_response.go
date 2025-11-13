package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDistributionResponse Response Object
type DeleteDistributionResponse struct {

	// 任务id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDistributionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDistributionResponse struct{}"
	}

	return strings.Join([]string{"DeleteDistributionResponse", string(data)}, " ")
}
