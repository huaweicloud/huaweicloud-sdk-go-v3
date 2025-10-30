package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryVpcCondition struct {

	// **参数解释**: 主机id列表 **取值范围**: 字符长度1-64位
	HostIdList *[]string `json:"host_id_list,omitempty"`

	// **参数解释**: VpcId列表 **取值范围**: 字符长度1-64位
	VpcIdList *[]string `json:"vpc_id_list,omitempty"`
}

func (o QueryVpcCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryVpcCondition struct{}"
	}

	return strings.Join([]string{"QueryVpcCondition", string(data)}, " ")
}
