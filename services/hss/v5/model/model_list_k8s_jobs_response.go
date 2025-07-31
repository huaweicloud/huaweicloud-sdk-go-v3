package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListK8sJobsResponse Response Object
type ListK8sJobsResponse struct {

	// 普通任务总量
	TotalNum *int64 `json:"total_num,omitempty"`

	// 最近更新时间
	LastUpdateTime *int64 `json:"last_update_time,omitempty"`

	// 普通任务基本信息列表
	JobInfoList    *[]ServerlessJobInfo `json:"job_info_list,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListK8sJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListK8sJobsResponse struct{}"
	}

	return strings.Join([]string{"ListK8sJobsResponse", string(data)}, " ")
}
