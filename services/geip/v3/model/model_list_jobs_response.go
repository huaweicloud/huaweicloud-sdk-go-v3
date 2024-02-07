package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobsResponse Response Object
type ListJobsResponse struct {

	// 本次请求的编号
	RequestId *string `json:"request_id,omitempty"`

	// 响应job对象
	Jobs *[]ListJobs `json:"jobs,omitempty"`

	PageInfo *ListGlobalEipsResponseBodyPageInfo `json:"page_info,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobsResponse struct{}"
	}

	return strings.Join([]string{"ListJobsResponse", string(data)}, " ")
}
