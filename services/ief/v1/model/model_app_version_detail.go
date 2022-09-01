package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// app详情
type AppVersionDetail struct {

	// 应用版本ID
	Id *string `json:"id,omitempty" xml:"id"`

	// 应用版本号
	Version *string `json:"version,omitempty" xml:"version"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at"`

	// 更新时间 只有更新后才会出现该字段
	UpdatedAt *string `json:"updated_at,omitempty" xml:"updated_at"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty" xml:"project_id"`

	// 镜像存储地址
	ImageUrl *string `json:"image_url,omitempty" xml:"image_url"`

	// 环境变量
	Envs []Env `json:"envs" xml:"envs"`

	// 卷配置
	Volumes []Volumes `json:"volumes" xml:"volumes"`

	Configs *AppConfigs `json:"configs,omitempty" xml:"configs"`

	Resources *Resources `json:"resources,omitempty" xml:"resources"`

	// 架构
	Arch *string `json:"arch,omitempty" xml:"arch"`

	// 启动命令
	Command *[]string `json:"command,omitempty" xml:"command"`

	// 参数
	Args *[]string `json:"args,omitempty" xml:"args"`

	LivenessProbe *ProbeDetail `json:"liveness_probe,omitempty" xml:"liveness_probe"`

	ReadinessProbe *ProbeDetail `json:"readiness_probe,omitempty" xml:"readiness_probe"`
}

func (o AppVersionDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppVersionDetail struct{}"
	}

	return strings.Join([]string{"AppVersionDetail", string(data)}, " ")
}
