package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConfigMapModelBoxDto struct {
	Configmap *ConfigMap `json:"configmap"`

	// 工作空间ID，默认为注册账号/子账号的default工作空间，可通过专业版HiLens控制台展开工作空间列表获取到工作空间ID
	WorkspaceId *string `json:"workspace_id,omitempty"`
}

func (o ConfigMapModelBoxDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigMapModelBoxDto struct{}"
	}

	return strings.Join([]string{"ConfigMapModelBoxDto", string(data)}, " ")
}
