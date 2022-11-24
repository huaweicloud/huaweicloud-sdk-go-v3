package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新用户请求体
type UserSettingDto struct {

	// 允许同时运行的作业数
	JobQuota int32 `json:"job_quota"`

	// 作业执行超时时长，单位天
	JobTimeout int32 `json:"job_timeout"`

	// 作业的CPU资源配额，单位核
	CpuQuota int32 `json:"cpu_quota"`

	// 作业的内存资源配额，单位GB
	MemQuota int32 `json:"mem_quota"`

	// 用户可创建项目数配额
	ProjectsPerUser *int32 `json:"projects_per_user,omitempty"`
}

func (o UserSettingDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserSettingDto struct{}"
	}

	return strings.Join([]string{"UserSettingDto", string(data)}, " ")
}
