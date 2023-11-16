package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBuildJobRequestBody 更新构建作业配置接口请求体
type UpdateBuildJobRequestBody struct {

	// 使用机器的架构
	Arch string `json:"arch"`

	// 构建任务所在项目的ID
	ProjectId string `json:"project_id"`

	// 任务名称
	JobName string `json:"job_name"`

	// 构建任务ID
	JobId string `json:"job_id"`

	// 是否自动更新子模块
	AutoUpdateSubModule *string `json:"auto_update_sub_module,omitempty"`

	// 执行机规格
	Flavor *string `json:"flavor,omitempty"`

	// 构建执行参数列表
	Parameters *[]UpdateBuildJobParameter `json:"parameters,omitempty"`

	// 构建执行SCM
	Scms *[]UpdateBuildJobScm `json:"scms,omitempty"`

	// 构建执行的步骤
	Steps []UpdateBuildJobSteps `json:"steps"`

	// host类型
	HostType *string `json:"host_type,omitempty"`

	// 构建的配置类型
	BuildConfigType *string `json:"build_config_type,omitempty"`
}

func (o UpdateBuildJobRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBuildJobRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateBuildJobRequestBody", string(data)}, " ")
}
