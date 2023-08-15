package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteFlinkJobsRequestBody
type BatchDeleteFlinkJobsRequestBody struct {

	//
	JobIds []int64 `json:"job_ids"`
}

func (o BatchDeleteFlinkJobsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteFlinkJobsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteFlinkJobsRequestBody", string(data)}, " ")
}
