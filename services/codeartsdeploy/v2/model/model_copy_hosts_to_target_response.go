package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CopyHostsToTargetResponse Response Object
type CopyHostsToTargetResponse struct {

	// 请求成功失败状态
	Status *string `json:"status,omitempty"`

	// 返回结果
	Result         *bool `json:"result,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CopyHostsToTargetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyHostsToTargetResponse struct{}"
	}

	return strings.Join([]string{"CopyHostsToTargetResponse", string(data)}, " ")
}
