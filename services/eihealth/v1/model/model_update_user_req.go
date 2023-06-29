package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUserReq 修改用户基本信息
type UpdateUserReq struct {

	// 用户手机号，纯数字，长度小于等于32位。必须与国家码同时存在。
	Mobile *string `json:"mobile,omitempty"`

	// 国家码。中国大陆为“0086”
	Areacode *string `json:"areacode,omitempty"`

	// 用户邮箱，需符合邮箱格式
	Email *string `json:"email,omitempty"`

	// 预验证凭证
	Ticket *string `json:"ticket,omitempty"`
}

func (o UpdateUserReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserReq struct{}"
	}

	return strings.Join([]string{"UpdateUserReq", string(data)}, " ")
}
