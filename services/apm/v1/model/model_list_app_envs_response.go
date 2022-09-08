package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAppEnvsResponse struct {

	// 环境信息列表
	Envs           *[]EnvNodeModel `json:"envs,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListAppEnvsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppEnvsResponse struct{}"
	}

	return strings.Join([]string{"ListAppEnvsResponse", string(data)}, " ")
}
