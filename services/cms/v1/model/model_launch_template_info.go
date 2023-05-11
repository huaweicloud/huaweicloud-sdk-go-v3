package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 启动模板，创建虚拟机时，使用到安全组、网络、镜像、flavor等信息
type LaunchTemplateInfo struct {

	// 模板的ID,唯一标识一个模板
	LaunchTemplateId string `json:"launch_template_id"`

	// 模板的版本，模板更新，都会增加版本
	Version string `json:"version"`
}

func (o LaunchTemplateInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LaunchTemplateInfo struct{}"
	}

	return strings.Join([]string{"LaunchTemplateInfo", string(data)}, " ")
}