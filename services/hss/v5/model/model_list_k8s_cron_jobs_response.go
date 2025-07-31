package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListK8sCronJobsResponse Response Object
type ListK8sCronJobsResponse struct {

	// 定时任务总量
	TotalNum *int64 `json:"total_num,omitempty"`

	// 最近更新时间
	LastUpdateTime *int64 `json:"last_update_time,omitempty"`

	// 定时任务基本信息列表
	CronjobInfoList *[]ServerlessCronJobInfo `json:"cronjob_info_list,omitempty"`
	HttpStatusCode  int                      `json:"-"`
}

func (o ListK8sCronJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListK8sCronJobsResponse struct{}"
	}

	return strings.Join([]string{"ListK8sCronJobsResponse", string(data)}, " ")
}
