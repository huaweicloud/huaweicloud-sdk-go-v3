package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEnvRequest Request Object
type DeleteEnvRequest struct {

	// 环境id
	EnvironmentId string `json:"environment_id"`
}

func (o DeleteEnvRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEnvRequest struct{}"
	}

	return strings.Join([]string{"DeleteEnvRequest", string(data)}, " ")
}
