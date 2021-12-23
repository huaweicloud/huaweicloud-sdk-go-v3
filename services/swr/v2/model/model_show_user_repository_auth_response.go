package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowUserRepositoryAuthResponse struct {
	// id

	Id *int32 `json:"id,omitempty"`
	// 组织名称。小写字母开头，后面跟小写字母、数字、小数点、下划线或中划线（其中下划线最多允许连续两个，小数点、下划线、中划线不能直接相连），小写字母或数字结尾，1-64个字符。

	Name *string `json:"name,omitempty"`

	SelfAuth *UserAuth `json:"self_auth,omitempty"`
	// 其他用户的权限

	OthersAuths    *[]UserAuth `json:"others_auths,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowUserRepositoryAuthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserRepositoryAuthResponse struct{}"
	}

	return strings.Join([]string{"ShowUserRepositoryAuthResponse", string(data)}, " ")
}
