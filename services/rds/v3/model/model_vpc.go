package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Vpc Vpc信息
type Vpc struct {

	// VPC 的 ID
	Id *string `json:"id,omitempty"`

	// VPC 的名字
	Name *string `json:"name,omitempty"`

	// VPC 下的可用子网列表
	Subnets *[]Subnet `json:"subnets,omitempty"`
}

func (o Vpc) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Vpc struct{}"
	}

	return strings.Join([]string{"Vpc", string(data)}, " ")
}
