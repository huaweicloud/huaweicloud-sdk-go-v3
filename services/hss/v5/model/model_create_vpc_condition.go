package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateVpcCondition struct {

	// **参数解释** 主机id列表 **约束限制**: 不涉及 **取值范围** 字符长度1-256位 **默认取值**: 不涉及
	HostIdList *[]string `json:"host_id_list,omitempty"`

	// **参数解释** vpc列表 **约束限制**: 不涉及 **取值范围** 字符长度1-256位 **默认取值**: 不涉及
	VpcList *[]VpcAndSubNet `json:"vpc_list,omitempty"`
}

func (o CreateVpcCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpcCondition struct{}"
	}

	return strings.Join([]string{"CreateVpcCondition", string(data)}, " ")
}
