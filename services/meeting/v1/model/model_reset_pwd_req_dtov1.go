package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResetPwdReqDtov1 struct {

	// 用户身份信息（手机号码或邮箱帐号或用户真实帐号）。
	User string `json:"user"`

	// 用户新的登录密码。 密码要求： * 长度范围要求8~32 * 至少包含大小写字母、数字 * 不能包含3个以上重复字符 * 密码不能包含与其对应的用户名（不区分大小写）以及逆序的用户名（不区分大小写）
	NewPwd string `json:"newPwd"`

	// * 1：临时密码，重置完密码后登录Web Portal根据配置可能需要强制修改密码 * 非1：正式密码，重置完密码后登录Web Portal不需要强制修改密码
	PassWordType *int32 `json:"passWordType,omitempty"`
}

func (o ResetPwdReqDtov1) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetPwdReqDtov1 struct{}"
	}

	return strings.Join([]string{"ResetPwdReqDtov1", string(data)}, " ")
}
