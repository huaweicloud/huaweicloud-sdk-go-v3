package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHarvestTaskResponse Response Object
type DeleteHarvestTaskResponse struct {

	// 任务ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteHarvestTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHarvestTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteHarvestTaskResponse", string(data)}, " ")
}
