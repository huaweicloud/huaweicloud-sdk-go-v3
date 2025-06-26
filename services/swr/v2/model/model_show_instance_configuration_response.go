package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceConfigurationResponse Response Object
type ShowInstanceConfigurationResponse struct {

	// 是否开启匿名登录
	AnonymousAccess *bool `json:"anonymous_access,omitempty"`
	HttpStatusCode  int   `json:"-"`
}

func (o ShowInstanceConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceConfigurationResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceConfigurationResponse", string(data)}, " ")
}
