package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListJobInstancesResponse struct {
	Total *int32 `json:"total,omitempty" xml:"total"`

	Instances      *[]JobInstance `json:"instances,omitempty" xml:"instances"`
	HttpStatusCode int            `json:"-"`
}

func (o ListJobInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListJobInstancesResponse", string(data)}, " ")
}
