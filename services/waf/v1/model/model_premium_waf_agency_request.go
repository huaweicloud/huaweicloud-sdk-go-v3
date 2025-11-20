package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PremiumWafAgencyRequest 创建独享引擎代理策略请求体
type PremiumWafAgencyRequest struct {

	// **参数解释：** 创建独享引擎代理策略名列表 **约束限制：** 已经存在的策略，无法再次创建 **取值范围：** 列表中支持以下策略名字 - evs_to_waf_operate_policy - vpc_to_waf_operate_policy - ecs_to_waf_operate_policy - elb_to_waf_operate_policy **默认取值：** 不涉及
	RoleNameList *[]string `json:"role_name_list,omitempty"`
}

func (o PremiumWafAgencyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PremiumWafAgencyRequest struct{}"
	}

	return strings.Join([]string{"PremiumWafAgencyRequest", string(data)}, " ")
}
