package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuthorizeCsmsAndKmsRequestBody 云堡垒机设置委托授权请求体。
type AuthorizeCsmsAndKmsRequestBody struct {
	Authorization *AgencyAuthorizeInfo `json:"authorization"`
}

func (o AuthorizeCsmsAndKmsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizeCsmsAndKmsRequestBody struct{}"
	}

	return strings.Join([]string{"AuthorizeCsmsAndKmsRequestBody", string(data)}, " ")
}
