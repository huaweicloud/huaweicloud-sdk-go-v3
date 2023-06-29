package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateExtensionAuthorizationResponse Response Object
type CreateExtensionAuthorizationResponse struct {

	// 返回值
	Result *bool `json:"result,omitempty"`

	// 状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateExtensionAuthorizationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateExtensionAuthorizationResponse struct{}"
	}

	return strings.Join([]string{"CreateExtensionAuthorizationResponse", string(data)}, " ")
}
