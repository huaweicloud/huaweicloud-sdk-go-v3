package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEnvironmentDetailRequest Request Object
type ShowEnvironmentDetailRequest struct {

	// 应用id
	ApplicationId string `json:"application_id"`

	// 环境id
	EnvironmentId string `json:"environment_id"`
}

func (o ShowEnvironmentDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEnvironmentDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowEnvironmentDetailRequest", string(data)}, " ")
}
