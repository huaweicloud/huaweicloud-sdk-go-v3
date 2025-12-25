package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableDataTransformationResponse Response Object
type EnableDataTransformationResponse struct {

	// UUID
	DataTransformationId *string `json:"data_transformation_id,omitempty"`

	Status *JobStatus `json:"status,omitempty"`

	ProcessStatus  *JobProcessStatus `json:"process_status,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o EnableDataTransformationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableDataTransformationResponse struct{}"
	}

	return strings.Join([]string{"EnableDataTransformationResponse", string(data)}, " ")
}
