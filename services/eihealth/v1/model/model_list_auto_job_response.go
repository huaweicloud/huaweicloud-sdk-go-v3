package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAutoJobResponse struct {

	// 自动作业列表
	AutoJobs *[]AutoJobListDto `json:"auto_jobs,omitempty"`

	// 作业总数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAutoJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAutoJobResponse struct{}"
	}

	return strings.Join([]string{"ListAutoJobResponse", string(data)}, " ")
}
