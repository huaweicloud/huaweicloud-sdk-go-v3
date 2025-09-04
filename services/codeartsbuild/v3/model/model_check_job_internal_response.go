package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckJobInternalResponse Response Object
type CheckJobInternalResponse struct {

	// 是否已开启内网安全访问
	Result *bool `json:"result,omitempty"`

	// 返回错误信息
	Error *string `json:"error,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CheckJobInternalResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckJobInternalResponse struct{}"
	}

	return strings.Join([]string{"CheckJobInternalResponse", string(data)}, " ")
}
