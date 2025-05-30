package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AppVersionDetail app详情
type AppVersionDetail struct {

	// 应用版本ID
	Id *string `json:"id,omitempty"`

	// 应用版本号
	Version *string `json:"version,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间 只有更新后才会出现该字段
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 镜像存储地址
	ImageUrl *string `json:"image_url,omitempty"`

	// 环境变量
	Envs []Env `json:"envs"`

	// 卷配置
	Volumes []Volumes `json:"volumes"`

	Configs *AppConfigs `json:"configs,omitempty"`

	Resources *Resources `json:"resources,omitempty"`

	// 架构
	Arch *string `json:"arch,omitempty"`

	// 启动命令
	Command *[]string `json:"command,omitempty"`

	// 参数
	Args *[]string `json:"args,omitempty"`

	LivenessProbe *ProbeDetail `json:"liveness_probe,omitempty"`

	ReadinessProbe *ProbeDetail `json:"readiness_probe,omitempty"`

	// NPU类型，支持D310、D310B，支持填写： - D310：D310类型 - D310B：D310B类型 - 不填表示为D310类型。
	NpuType *string `json:"npu_type,omitempty"`
}

func (o AppVersionDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppVersionDetail struct{}"
	}

	return strings.Join([]string{"AppVersionDetail", string(data)}, " ")
}
