package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBuildJobRequestBody 创建构建任务接口请求体
type CreateBuildJobRequestBody struct {

	// 使用机器的架构
	Arch string `json:"arch"`

	// 构建任务所在项目的ID
	ProjectId string `json:"project_id"`

	// 任务名称
	JobName string `json:"job_name"`

	// 是否自动更新子模块
	AutoUpdateSubModule *string `json:"auto_update_sub_module,omitempty"`

	// 执行机规格
	Flavor *string `json:"flavor,omitempty"`

	// 构建执行参数列表
	Parameters *[]CreateBuildJobParameter `json:"parameters,omitempty"`

	// 任务分组id
	GroupId *string `json:"group_id,omitempty"`

	Timeout *CreateBuildTimeout `json:"timeout,omitempty"`

	// 构建执行SCM
	Scms *[]CreateBuildJobScm `json:"scms,omitempty"`

	// 构建执行的步骤
	Steps []CreateBuildJobSteps `json:"steps"`

	// host类型
	HostType *string `json:"host_type,omitempty"`

	// 构建的配置类型
	BuildConfigType *string `json:"build_config_type,omitempty"`

	// 提交代码触发构建开关
	BuildIfCodeUpdated *bool `json:"build_if_code_updated,omitempty"`

	// 定时任务触发器集合
	Triggers *[]Trigger `json:"triggers,omitempty"`
}

func (o CreateBuildJobRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBuildJobRequestBody struct{}"
	}

	return strings.Join([]string{"CreateBuildJobRequestBody", string(data)}, " ")
}
