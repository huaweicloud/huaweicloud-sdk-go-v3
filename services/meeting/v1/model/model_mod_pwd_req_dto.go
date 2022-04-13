package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModPwdReqDto struct {
	// 帐号，必须是携带域名的帐号 maxLength: 255 minLength: 1

	Account string `json:"account"`
	// 用户旧的登录密码 maxLength: 255 minLength: 1

	OldPwd string `json:"oldPwd"`
	// 用户新的登录密码 密码要求： * 长度范围要求8~32 * 至少包含大小写字母、数字 * 旧密码和新密码不能相同 * 上次修改密码后5分钟内不能更新密码 * 不能与最近使用的旧密码相同 * 不能包含3个以上重复字符 * 密码不能包含与其对应的用户名（不区分大小写）以及逆序的用户名（不区分大小写） * 新密码与旧密码之间允许的最少不相同字符数为2个

	NewPwd string `json:"newPwd"`
}

func (o ModPwdReqDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModPwdReqDto struct{}"
	}

	return strings.Join([]string{"ModPwdReqDto", string(data)}, " ")
}
