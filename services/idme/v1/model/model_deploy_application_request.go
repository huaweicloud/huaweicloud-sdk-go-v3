package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeployApplicationRequest Request Object
type DeployApplicationRequest struct {

	// 运行服务ID。
	EnvId string `json:"env_id"`

	// 待部署应用的ID。
	AppId string `json:"app_id"`

	Body *DeployApplicationRequestBody `json:"body,omitempty"`
}

func (o DeployApplicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeployApplicationRequest struct{}"
	}

	return strings.Join([]string{"DeployApplicationRequest", string(data)}, " ")
}
