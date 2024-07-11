package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHostFromEnvironmentResponse Response Object
type DeleteHostFromEnvironmentResponse struct {

	// 请求成功失败状态
	Status *string `json:"status,omitempty"`

	// 环境id
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteHostFromEnvironmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHostFromEnvironmentResponse struct{}"
	}

	return strings.Join([]string{"DeleteHostFromEnvironmentResponse", string(data)}, " ")
}
