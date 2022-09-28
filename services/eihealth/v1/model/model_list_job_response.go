package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListJobResponse struct {

	// 作业列表
	Jobs *[]JobListDto `json:"jobs,omitempty"`

	// 作业总数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobResponse struct{}"
	}

	return strings.Join([]string{"ListJobResponse", string(data)}, " ")
}
