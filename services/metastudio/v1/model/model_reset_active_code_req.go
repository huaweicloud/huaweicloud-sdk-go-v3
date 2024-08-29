package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetActiveCodeReq 重置激活码请求。
type ResetActiveCodeReq struct {

	// 是否清除鉴权凭证。
	CleanAuthCredential *bool `json:"clean_auth_credential,omitempty"`
}

func (o ResetActiveCodeReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetActiveCodeReq struct{}"
	}

	return strings.Join([]string{"ResetActiveCodeReq", string(data)}, " ")
}
