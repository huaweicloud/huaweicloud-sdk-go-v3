package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeploymentTemplate 部署模板信息
type DeploymentTemplate struct {
	Configs *PodConfig `json:"configs,omitempty"`

	// 部署应用列表
	Apps *[]AppDef `json:"apps,omitempty"`
}

func (o DeploymentTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeploymentTemplate struct{}"
	}

	return strings.Join([]string{"DeploymentTemplate", string(data)}, " ")
}
