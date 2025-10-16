package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifyInstancePortRequestBody struct {

	// **参数解释**: 实例设置的数据库端口号。 **约束限制**: 不涉及。 **取值范围**: 1024~39998（其中2378、2379、2380、4999、5000、5999、6000、60001、8097、8098、12016、12017、20049、20050、21731、21732、32122、32123和32124被系统占用不可设置）。 **默认取值**: 不涉及。
	Port int32 `json:"port"`
}

func (o ModifyInstancePortRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyInstancePortRequestBody struct{}"
	}

	return strings.Join([]string{"ModifyInstancePortRequestBody", string(data)}, " ")
}
