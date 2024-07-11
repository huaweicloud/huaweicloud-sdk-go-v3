package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CopyApplicationResponse Response Object
type CopyApplicationResponse struct {
	Result *AppBaseInfo `json:"result,omitempty"`

	// 请求成功失败状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CopyApplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyApplicationResponse struct{}"
	}

	return strings.Join([]string{"CopyApplicationResponse", string(data)}, " ")
}
