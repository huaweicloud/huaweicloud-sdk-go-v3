package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type UpdateUserOption struct {
	// IAM用户访问方式。 - default：默认访问模式，编程访问和管理控制台访问。 - programmatic：编程访问。 - console：管理控制台访问。

	AccessMode *string `json:"access_mode,omitempty"`
	// 新IAM用户名，长度5~32字符之间，首位不能为数字，特殊字符只能包含下划线“_”、中划线“-”和空格。

	Name *string `json:"name,omitempty"`
	// IAM用户新密码。 - 系统默认密码最小长度为6字符，在6-32字符之间支持用户自定义密码长度。 - 至少包含以下四种字符中的两种： 大写字母、小写字母、数字和特殊字符。 - 必须满足账户设置中[密码策略](https://support.huaweicloud.com/usermanual-iam/iam_01_0607.html)的要求。 - 新密码不能与当前密码相同。

	Password *string `json:"password,omitempty"`
	// IAM用户新邮箱，需符合邮箱格式，长度小于等于255字符。

	Email *string `json:"email,omitempty"`
	// 国家码。必须与手机号同时存在。中国大陆为“0086”。

	Areacode *string `json:"areacode,omitempty"`
	// IAM用户新手机号，纯数字，长度小于等于32字符。必须与国家码同时存在。

	Phone *string `json:"phone,omitempty"`
	// 是否启用IAM用户。true为启用，false为停用，默认为true。

	Enabled *bool `json:"enabled,omitempty"`
	// IAM用户密码状态。true：需要修改密码，false：正常。

	PwdStatus *bool `json:"pwd_status,omitempty"`
	// IAM用户在外部系统中的类型。长度小于等于64字符。xuser_type如果存在，则需要与同一租户中的xaccount_type、xdomain_type校验，须与xuser_id同时存在。 >外部系统指与华为云对接的外部企业管理系统，xaccount_type、xaccount_id、xdomain_type、xdomain_id、xuser_type、xuser_id等参数值，无法在华为云获取，请咨询企业管理员。

	XuserType *string `json:"xuser_type,omitempty"`
	// IAM用户在外部系统中的ID。长度小于等于128字符，必须与xuser_type同时存在。 >外部系统指与华为云对接的外部企业管理系统，xaccount_type、xaccount_id、xdomain_type、xdomain_id、xuser_type、xuser_id等参数值，无法在华为云获取，请咨询企业管理员。

	XuserId *string `json:"xuser_id,omitempty"`
	// IAM用户新描述信息。

	Description *string `json:"description,omitempty"`
}

func (o UpdateUserOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserOption struct{}"
	}

	return strings.Join([]string{"UpdateUserOption", string(data)}, " ")
}
