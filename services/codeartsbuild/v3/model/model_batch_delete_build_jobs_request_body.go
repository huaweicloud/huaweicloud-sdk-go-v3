package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteBuildJobsRequestBody struct {

	// jobId列表
	JobIds []string `json:"job_ids"`
}

func (o BatchDeleteBuildJobsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteBuildJobsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteBuildJobsRequestBody", string(data)}, " ")
}
