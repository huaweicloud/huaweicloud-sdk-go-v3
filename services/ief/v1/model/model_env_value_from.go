package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 环境变量可使用配置项和密钥导入
type EnvValueFrom struct {
	Secret *Secrets `json:"secret,omitempty"`

	Configmap *ConfigsMap `json:"configmap,omitempty"`
}

func (o EnvValueFrom) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnvValueFrom struct{}"
	}

	return strings.Join([]string{"EnvValueFrom", string(data)}, " ")
}
