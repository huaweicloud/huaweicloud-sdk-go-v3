package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VpcAndSubNet struct {

	// **参数解释** Vpc标识ID **约束限制**: 不涉及 **取值范围** 字符长度1-256位 **默认取值**: 不涉及
	VpcId string `json:"vpc_id"`

	// **参数解释** 子网id的列表 **约束限制**: 不涉及
	SubnetIdList []string `json:"subnet_id_list"`
}

func (o VpcAndSubNet) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcAndSubNet struct{}"
	}

	return strings.Join([]string{"VpcAndSubNet", string(data)}, " ")
}
