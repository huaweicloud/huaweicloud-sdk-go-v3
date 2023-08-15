package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GetUserRsp struct {

	// 用户id
	Id *string `json:"id,omitempty"`

	// 用户名，长度4~31之间，首位不能为数字，特殊字符只能包含下划线“_”、中划线“-”和空格
	Name *string `json:"name,omitempty"`

	// 角色类型：管理员(ADMIN)、操作者(OPERATOR)
	Role *string `json:"role,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 用户邮箱，需符合邮箱格式
	Email *string `json:"email,omitempty"`

	// 用户手机号，纯数字，长度小于等于32位。必须与国家码同时存在。
	Phone *string `json:"phone,omitempty"`

	// 国家码。中国大陆为“0086”
	Areacode *string `json:"areacode,omitempty"`

	// 是否domain用户
	IsDomainOwner *bool `json:"is_domain_owner,omitempty"`

	// 用户创建时间，UTC时间
	CreateTime *string `json:"create_time,omitempty"`

	// 是否需要修改密码
	PwdStatus *bool `json:"pwd_status,omitempty"`

	// 更新时间，UTC时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 来源，PLATFORM或者IAM
	Source *string `json:"source,omitempty"`
}

func (o GetUserRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetUserRsp struct{}"
	}

	return strings.Join([]string{"GetUserRsp", string(data)}, " ")
}
