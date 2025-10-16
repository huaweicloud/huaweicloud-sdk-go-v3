package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PremiumWafAgencyRequest 创建独享引擎代理请求体列表
type PremiumWafAgencyRequest struct {

	// 具体代理策略名
	RoleNameList *[]string `json:"role_name_list,omitempty"`
}

func (o PremiumWafAgencyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PremiumWafAgencyRequest struct{}"
	}

	return strings.Join([]string{"PremiumWafAgencyRequest", string(data)}, " ")
}
