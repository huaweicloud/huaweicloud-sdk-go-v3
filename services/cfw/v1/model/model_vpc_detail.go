package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VpcDetail struct {

	// id
	Id *string `json:"id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// vpc cidr
	Cidr *string `json:"cidr,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`
}

func (o VpcDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcDetail struct{}"
	}

	return strings.Join([]string{"VpcDetail", string(data)}, " ")
}
