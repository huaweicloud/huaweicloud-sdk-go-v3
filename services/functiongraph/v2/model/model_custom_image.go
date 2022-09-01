package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomImage struct {

	// 是否启用
	Enabled *bool `json:"enabled,omitempty" xml:"enabled"`

	// 镜像地址
	Image *string `json:"image,omitempty" xml:"image"`

	// 启动容器镜像的命令
	Command *string `json:"command,omitempty" xml:"command"`

	// 启动容器镜像的命令行参数
	Args *string `json:"args,omitempty" xml:"args"`

	// 镜像容器工作目录
	WorkingDir *string `json:"working_dir,omitempty" xml:"working_dir"`

	// 镜像容器的用户id
	Uid *string `json:"uid,omitempty" xml:"uid"`

	// 镜像容器的用户组id
	Gid *string `json:"gid,omitempty" xml:"gid"`
}

func (o CustomImage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomImage struct{}"
	}

	return strings.Join([]string{"CustomImage", string(data)}, " ")
}
