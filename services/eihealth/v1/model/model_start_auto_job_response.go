package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartAutoJobResponse Response Object
type StartAutoJobResponse struct {

	// 作业id
	Id *string `json:"id,omitempty"`

	MaxPlatformFlavor *TaskResourceDto `json:"max_platform_flavor,omitempty"`

	// 筛选后的app集合
	AppInfos *[]AppFilterDto `json:"app_infos,omitempty"`

	JobInfo        *JobFilterDto `json:"job_info,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o StartAutoJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartAutoJobResponse struct{}"
	}

	return strings.Join([]string{"StartAutoJobResponse", string(data)}, " ")
}
