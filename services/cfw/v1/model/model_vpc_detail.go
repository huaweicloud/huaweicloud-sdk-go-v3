package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VpcDetail struct {

	// **参数解释**： 创建引流VPC产生的随机UUID **取值范围**： 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**： 引流VPC名称 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 虚拟私有云下可用子网的范围 **取值范围**： 10.0.0.0/8~24 172.16.0.0/12~24 192.168.0.0/16~24
	Cidr *string `json:"cidr,omitempty"`
}

func (o VpcDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcDetail struct{}"
	}

	return strings.Join([]string{"VpcDetail", string(data)}, " ")
}
