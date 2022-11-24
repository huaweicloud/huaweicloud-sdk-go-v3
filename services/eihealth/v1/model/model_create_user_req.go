package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建用户请求体
type CreateUserReq struct {

	// 用户名，长度4~31之间，首位不能为数字，特殊字符只能包含下划线“_”、中划线“-”和空格
	Name string `json:"name"`

	// 用户密码，在8-32位之间支持用户自定义密码长度，至少包含以下四种字符中的三种： 大写字母、小写字母、数字和特殊字符。
	Password string `json:"password"`

	// 角色类型：管理员(ADMIN)、操作者(OPERATOR)
	Role string `json:"role"`

	// 用户邮箱，需符合邮箱格式
	Email *string `json:"email,omitempty"`

	// 用户手机号，纯数字，长度小于等于32位。必须与国家码同时存在。
	Phone *string `json:"phone,omitempty"`

	// 国家码。中国大陆为“0086”
	Areacode *string `json:"areacode,omitempty"`

	Settings *UserSettingDto `json:"settings,omitempty"`
}

func (o CreateUserReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUserReq struct{}"
	}

	return strings.Join([]string{"CreateUserReq", string(data)}, " ")
}
