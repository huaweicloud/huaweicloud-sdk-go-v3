package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusteringJobResponse Response Object
type ShowClusteringJobResponse struct {
	BasicInfo *DrugJobDto `json:"basic_info,omitempty"`

	File *DrugFile `json:"file,omitempty"`

	// 部分失败原因和数量。
	FailedReasons *[]FailedReasonRecord `json:"failed_reasons,omitempty"`

	JobResult      *JobResult `json:"job_result,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ShowClusteringJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusteringJobResponse struct{}"
	}

	return strings.Join([]string{"ShowClusteringJobResponse", string(data)}, " ")
}
