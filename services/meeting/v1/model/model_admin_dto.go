package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 管理员信息
type AdminDto struct {
	// 用户账号，帐号只能包含大小写字母、数字、_、-、.、@符号，不能为纯数字和@后面带.号

	Account string `json:"account"`
	// 名称

	Name string `json:"name"`
	// 若携带则以前台携带为准，否则后台默认生成,密码必须满足: - 1、6-32位 - 2、不能和账号的正序和倒序一致 - 3、至少包含两种字符类型：小写字母、大写字母、数字、特殊字符（` ~ ! @ # $ % ^ & * ( ) - _ = + \\ | [ { } ] ; : \\\" ,' < . > / ?

	Pwd string `json:"pwd"`
	// 邮箱，管理员手机和邮箱必填其一，否则无法重置密码。如果企业短信开关关闭，则邮箱必填。格式必须满足(^$|^[\\\\w-+]+(\\\\.[\\\\w-+]+)*@[\\\\w-]+(\\\\.[\\\\w-]+)*(\\\\.[\\\\w-]{1,})$)

	Email *string `json:"email,omitempty"`
	// 手机号，必须加上国家码，例如中国大陆手机+86xxxxxxx，当填写手机号时 “country”参数必填,手机格式必须满足(^$|^[+]?[0-9]+$)

	Phone *string `json:"phone,omitempty"`
	// 若smsNumber为手机号,则需带上手机号所属的国家。 例如国家为中国大陆则country参数取值为chinaPR 国家和国家码的对应关系请参考：https://support.huaweicloud.com/api-meeting/meeting_21_0109.html

	Country *string `json:"country,omitempty"`
}

func (o AdminDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AdminDto struct{}"
	}

	return strings.Join([]string{"AdminDto", string(data)}, " ")
}
