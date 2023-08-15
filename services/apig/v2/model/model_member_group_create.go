package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MemberGroupCreate struct {

	// VPC通道后端服务器组名称
	MemberGroupName string `json:"member_group_name"`

	// VPC通道后端服务器组描述
	MemberGroupRemark *string `json:"member_group_remark,omitempty"`

	// VPC通道后端服务器组权重值。  当前服务器组存在服务器且此权重值存在时，自动使用此权重值分配权重。
	MemberGroupWeight *int32 `json:"member_group_weight,omitempty"`

	// VPC通道后端服务器组的字典编码  支持英文，数字，特殊字符（-_.）  暂不支持
	DictCode *string `json:"dict_code,omitempty"`

	// VPC通道后端服务器组的版本，仅VPC通道类型为微服务时支持。
	MicroserviceVersion *string `json:"microservice_version,omitempty"`

	// VPC通道后端服务器组的端口号，仅VPC通道类型为微服务时支持。端口号为0时后端服务器组下的所有地址沿用原来负载端口继承逻辑。
	MicroservicePort *int32 `json:"microservice_port,omitempty"`

	// VPC通道后端服务器组的标签，仅VPC通道类型为微服务时支持。
	MicroserviceLabels *[]MicroserviceLabel `json:"microservice_labels,omitempty"`
}

func (o MemberGroupCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MemberGroupCreate struct{}"
	}

	return strings.Join([]string{"MemberGroupCreate", string(data)}, " ")
}
