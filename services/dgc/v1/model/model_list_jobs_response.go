package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListJobsResponse struct {
	Total *int32 `json:"total,omitempty" xml:"total"`

	Jobs           *[]JobInfo `json:"jobs,omitempty" xml:"jobs"`
	HttpStatusCode int        `json:"-"`
}

func (o ListJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobsResponse struct{}"
	}

	return strings.Join([]string{"ListJobsResponse", string(data)}, " ")
}
