package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreatePortalRefNonceResponse struct {

	// 用于跳转登录的nonce信息。 说明： 通过链接https://bmeeting.huaweicloud.com/?lang=zh-CN&nonce=xxxxxxxxxxxxx#/login进行免登陆跳转。
	Nonce          *string `json:"nonce,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePortalRefNonceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePortalRefNonceResponse struct{}"
	}

	return strings.Join([]string{"CreatePortalRefNonceResponse", string(data)}, " ")
}
