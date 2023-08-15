package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOpenIdConnectConfigRequestBody 请求体
type CreateOpenIdConnectConfigRequestBody struct {
	OpenidConnectConfig *CreateOpenIdConnectConfig `json:"openid_connect_config"`
}

func (o CreateOpenIdConnectConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOpenIdConnectConfigRequestBody struct{}"
	}

	return strings.Join([]string{"CreateOpenIdConnectConfigRequestBody", string(data)}, " ")
}
