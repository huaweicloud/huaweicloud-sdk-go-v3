package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobInstanceResponse Response Object
type ShowJobInstanceResponse struct {
	Status *string `json:"status,omitempty"`

	PlanTime *int64 `json:"planTime,omitempty"`

	StartTime *int64 `json:"startTime,omitempty"`

	EndTime *int64 `json:"endTime,omitempty"`

	ExecuteTime *int64 `json:"executeTime,omitempty"`

	InstancesId *string `json:"instancesId,omitempty"`

	Total *int32 `json:"total,omitempty"`

	Nodes          *[]NodeInstance `json:"nodes,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowJobInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobInstanceResponse struct{}"
	}

	return strings.Join([]string{"ShowJobInstanceResponse", string(data)}, " ")
}
