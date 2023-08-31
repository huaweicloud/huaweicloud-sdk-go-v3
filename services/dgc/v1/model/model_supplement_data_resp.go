package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SupplementDataResp struct {
	EndDate float32 `json:"endDate,omitempty"`

	JobList *[]string `json:"jobList,omitempty"`

	Name *string `json:"name,omitempty"`

	Parallel *int32 `json:"parallel,omitempty"`

	StartDate float32 `json:"startDate,omitempty"`

	Status *string `json:"status,omitempty"`

	SubmittedDate float32 `json:"submittedDate,omitempty"`

	SupplementDataInstanceTime *interface{} `json:"supplement_data_instance_time,omitempty"`

	SupplementDataRunTime *interface{} `json:"supplement_data_run_time,omitempty"`

	Type *int32 `json:"type,omitempty"`

	UserName *string `json:"userName,omitempty"`
}

func (o SupplementDataResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupplementDataResp struct{}"
	}

	return strings.Join([]string{"SupplementDataResp", string(data)}, " ")
}
