package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateTokenPolicyReqBody struct {

	// 是否允许获取Token，默认为true，设置为false后将不允许获取账号下所有身份类型（IAM用户、委托、联邦用户）的Token（联邦认证获取的unscoped token不受Token策略影响）。
	TokenEnabled *bool `json:"token_enabled,omitempty"`
}

func (o UpdateTokenPolicyReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTokenPolicyReqBody struct{}"
	}

	return strings.Join([]string{"UpdateTokenPolicyReqBody", string(data)}, " ")
}
