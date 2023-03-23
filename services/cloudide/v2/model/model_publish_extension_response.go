package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type PublishExtensionResponse struct {

	// 返回值
	Result *interface{} `json:"result,omitempty"`

	// 状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o PublishExtensionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishExtensionResponse struct{}"
	}

	return strings.Join([]string{"PublishExtensionResponse", string(data)}, " ")
}
