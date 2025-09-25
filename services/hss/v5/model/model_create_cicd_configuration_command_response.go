package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCicdConfigurationCommandResponse Response Object
type CreateCicdConfigurationCommandResponse struct {

	// **参数解释**： cicd接入配置命令 **取值范围**： 字符长度1-1024位
	CicdCommand    *string `json:"cicd_command,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateCicdConfigurationCommandResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCicdConfigurationCommandResponse struct{}"
	}

	return strings.Join([]string{"CreateCicdConfigurationCommandResponse", string(data)}, " ")
}
