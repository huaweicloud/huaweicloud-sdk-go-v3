package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableDataTransformationResponse Response Object
type DisableDataTransformationResponse struct {

	// UUID
	DataTransformationId *string `json:"data_transformation_id,omitempty"`

	Status *JobStatus `json:"status,omitempty"`

	ProcessStatus  *JobProcessStatus `json:"process_status,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o DisableDataTransformationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableDataTransformationResponse struct{}"
	}

	return strings.Join([]string{"DisableDataTransformationResponse", string(data)}, " ")
}
