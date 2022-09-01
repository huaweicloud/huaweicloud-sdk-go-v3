package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type Job struct {

	// API类型，固定值“Job”，该值不可修改。
	Kind *string `json:"kind,omitempty" xml:"kind"`

	// API版本，固定值“v3”，该值不可修改。
	ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion"`

	Metadata *JobMetadata `json:"metadata,omitempty" xml:"metadata"`

	Spec *JobSpec `json:"spec,omitempty" xml:"spec"`

	Status *JobStatus `json:"status,omitempty" xml:"status"`
}

func (o Job) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Job struct{}"
	}

	return strings.Join([]string{"Job", string(data)}, " ")
}
