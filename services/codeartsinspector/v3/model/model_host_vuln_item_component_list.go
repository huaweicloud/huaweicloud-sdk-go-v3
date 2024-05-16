package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HostVulnItemComponentList struct {

	// 内容名称
	ComponentName *string `json:"componentName,omitempty"`

	// 安装版本
	ComponentInstallVersion *string `json:"componentInstallVersion,omitempty"`

	// 已经修复版本
	ComponentFixedVersion *string `json:"componentFixedVersion,omitempty"`
}

func (o HostVulnItemComponentList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostVulnItemComponentList struct{}"
	}

	return strings.Join([]string{"HostVulnItemComponentList", string(data)}, " ")
}
