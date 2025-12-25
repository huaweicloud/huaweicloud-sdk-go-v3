package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Department 资产所属部门
type Department struct {

	// 资产所属部门名称id，若不填写为default
	Id *string `json:"id,omitempty"`

	// 资产所属部门名称
	Name *string `json:"name,omitempty"`
}

func (o Department) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Department struct{}"
	}

	return strings.Join([]string{"Department", string(data)}, " ")
}
