package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowJobInfosRequest struct {

	// Job ID
	JobId string `json:"job_id" xml:"job_id"`
}

func (o ShowJobInfosRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobInfosRequest struct{}"
	}

	return strings.Join([]string{"ShowJobInfosRequest", string(data)}, " ")
}
