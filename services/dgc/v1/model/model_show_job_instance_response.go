package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowJobInstanceResponse struct {
	Status *string `json:"status,omitempty" xml:"status"`

	PlanTime *int32 `json:"planTime,omitempty" xml:"planTime"`

	StartTime *int32 `json:"startTime,omitempty" xml:"startTime"`

	EndTime *int32 `json:"endTime,omitempty" xml:"endTime"`

	ExecuteTime *int32 `json:"executeTime,omitempty" xml:"executeTime"`

	InstancesId *string `json:"instancesId,omitempty" xml:"instancesId"`

	Total *int32 `json:"total,omitempty" xml:"total"`

	Nodes          *[]NodeInstance `json:"nodes,omitempty" xml:"nodes"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowJobInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobInstanceResponse struct{}"
	}

	return strings.Join([]string{"ShowJobInstanceResponse", string(data)}, " ")
}
