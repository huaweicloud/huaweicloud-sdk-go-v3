package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowJobsResponse struct {
	// 作业数,查询单个作业时为0

	Total *int32 `json:"total,omitempty"`
	// 作业列表，请参见jobs参数说明

	Jobs *[]Job `json:"jobs,omitempty"`
	// 返回指定页号的作业

	PageNo *int32 `json:"page_no,omitempty"`
	// 每页作业数

	PageSize       *int32 `json:"page_size,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobsResponse struct{}"
	}

	return strings.Join([]string{"ShowJobsResponse", string(data)}, " ")
}
