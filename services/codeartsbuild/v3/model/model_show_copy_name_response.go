package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCopyNameResponse Response Object
type ShowCopyNameResponse struct {

	// 状态
	Status *string `json:"status,omitempty"`

	// 错误
	Error *interface{} `json:"error,omitempty"`

	// 结果
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowCopyNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCopyNameResponse struct{}"
	}

	return strings.Join([]string{"ShowCopyNameResponse", string(data)}, " ")
}
