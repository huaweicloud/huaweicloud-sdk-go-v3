package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDataTransformationResponse Response Object
type DeleteDataTransformationResponse struct {

	// UUID
	DataTransformationId *string `json:"data_transformation_id,omitempty"`

	// 毫秒时间戳
	DeleteTime *int64 `json:"delete_time,omitempty"`

	ProcessStatus  *JobProcessStatus `json:"process_status,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o DeleteDataTransformationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDataTransformationResponse struct{}"
	}

	return strings.Join([]string{"DeleteDataTransformationResponse", string(data)}, " ")
}
