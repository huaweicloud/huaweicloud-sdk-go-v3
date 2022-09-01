package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowExtensionAuthorizationResponse struct {

	// 返回值
	Result *interface{} `json:"result,omitempty" xml:"result"`

	// 状态
	Status         *string `json:"status,omitempty" xml:"status"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowExtensionAuthorizationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowExtensionAuthorizationResponse struct{}"
	}

	return strings.Join([]string{"ShowExtensionAuthorizationResponse", string(data)}, " ")
}
