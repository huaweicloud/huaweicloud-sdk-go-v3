package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUserByDomainReq 最终租户修改子用户安全信息
type UpdateUserByDomainReq struct {

	// 新密码，在8-32位之间支持用户自定义密码长度，至少包含以下四种字符中的两种： 大写字母、小写字母、数字和特殊字符。
	Password *string `json:"password,omitempty"`

	// 用户手机号，纯数字，长度小于等于32位，当且仅当重置手机号时传入空串。必须与国家码同时存在。
	Mobile *string `json:"mobile,omitempty"`

	// 国家码，当且仅当重置手机号时传入空串。中国大陆为“0086”
	Areacode *string `json:"areacode,omitempty"`

	// 用户邮箱，需符合邮箱格式
	Email *string `json:"email,omitempty"`
}

func (o UpdateUserByDomainReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserByDomainReq struct{}"
	}

	return strings.Join([]string{"UpdateUserByDomainReq", string(data)}, " ")
}
