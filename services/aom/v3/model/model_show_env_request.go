package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEnvRequest Request Object
type ShowEnvRequest struct {

	// 环境id
	EnvironmentId string `json:"environment_id"`
}

func (o ShowEnvRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEnvRequest struct{}"
	}

	return strings.Join([]string{"ShowEnvRequest", string(data)}, " ")
}
