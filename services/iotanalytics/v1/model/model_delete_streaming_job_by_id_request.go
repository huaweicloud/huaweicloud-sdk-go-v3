package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteStreamingJobByIdRequest struct {
	// 作业ID

	JobId string `json:"job_id"`
}

func (o DeleteStreamingJobByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStreamingJobByIdRequest struct{}"
	}

	return strings.Join([]string{"DeleteStreamingJobByIdRequest", string(data)}, " ")
}
