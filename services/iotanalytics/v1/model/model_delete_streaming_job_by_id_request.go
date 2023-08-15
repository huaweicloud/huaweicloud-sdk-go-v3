package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteStreamingJobByIdRequest Request Object
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
