package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobsRequestBody job_ids列表
type JobsRequestBody struct {

	// job_ids列表
	JobIds []string `json:"job_ids"`
}

func (o JobsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobsRequestBody struct{}"
	}

	return strings.Join([]string{"JobsRequestBody", string(data)}, " ")
}
