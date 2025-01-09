package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateAgencyPolicyRequestBody struct {

	// 委托解绑的权限策略集合。
	UnbindRoleNames []string `json:"unbind_role_names"`

	// 委托绑定的权限策略集合。
	BindRoleNames []string `json:"bind_role_names"`
}

func (o UpdateAgencyPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAgencyPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAgencyPolicyRequestBody", string(data)}, " ")
}
