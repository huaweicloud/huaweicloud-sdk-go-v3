package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDockerLoginResponse Response Object
type ShowDockerLoginResponse struct {

	// 临时登录指令
	LoginCmd       *string `json:"login_cmd,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDockerLoginResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDockerLoginResponse struct{}"
	}

	return strings.Join([]string{"ShowDockerLoginResponse", string(data)}, " ")
}
