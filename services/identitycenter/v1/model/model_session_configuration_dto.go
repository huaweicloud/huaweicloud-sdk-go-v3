package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SessionConfigurationDto the struct of SessionConfiguration
type SessionConfigurationDto struct {

	// 会话生效时间
	MaxAuthenticationAge string `json:"max_authentication_age"`
}

func (o SessionConfigurationDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SessionConfigurationDto struct{}"
	}

	return strings.Join([]string{"SessionConfigurationDto", string(data)}, " ")
}
