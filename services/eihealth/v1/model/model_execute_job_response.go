package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteJobResponse Response Object
type ExecuteJobResponse struct {

	// 作业id
	Id *string `json:"id,omitempty"`

	// 筛选后的app集合
	AppInfos *[]AppFilterDto `json:"app_infos,omitempty"`

	JobInfo        *JobFilterDto `json:"job_info,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ExecuteJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteJobResponse struct{}"
	}

	return strings.Join([]string{"ExecuteJobResponse", string(data)}, " ")
}
