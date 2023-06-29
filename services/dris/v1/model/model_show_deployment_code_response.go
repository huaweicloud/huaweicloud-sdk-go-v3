package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDeploymentCodeResponse Response Object
type ShowDeploymentCodeResponse struct {

	// **参数说明**：生成的安装命令。
	Cmd            *string `json:"cmd,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDeploymentCodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeploymentCodeResponse struct{}"
	}

	return strings.Join([]string{"ShowDeploymentCodeResponse", string(data)}, " ")
}
