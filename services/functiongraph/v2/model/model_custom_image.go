package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CustomImage 用户容器镜像。
type CustomImage struct {

	// 是否启用
	Enabled *bool `json:"enabled,omitempty"`

	// 镜像地址
	Image *string `json:"image,omitempty"`

	// 启动容器镜像的命令
	Command *string `json:"command,omitempty"`

	// 启动容器镜像的命令行参数
	Args *string `json:"args,omitempty"`

	// 镜像容器工作目录
	WorkingDir *string `json:"working_dir,omitempty"`

	// 镜像容器的用户id
	Uid *string `json:"uid,omitempty"`

	// 镜像容器的用户组id
	Gid *string `json:"gid,omitempty"`
}

func (o CustomImage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomImage struct{}"
	}

	return strings.Join([]string{"CustomImage", string(data)}, " ")
}
