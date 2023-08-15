package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePortalRefNonceResponse Response Object
type CreatePortalRefNonceResponse struct {

	// 用于跳转登录的nonce信息。同一个nonce只能使用一次。 > 通过链接https://meeting.huaweicloud.com/?lang=zh-CN&nonce=xxxxxxxxxxxxx#/login进行免登陆跳转。
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
