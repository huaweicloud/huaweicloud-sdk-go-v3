package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAddressSetDto 更新地址组信息
type UpdateAddressSetDto struct {

	// 地址组名称
	Name *string `json:"name,omitempty"`

	// 地址组描述
	Description *string `json:"description,omitempty"`
}

func (o UpdateAddressSetDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAddressSetDto struct{}"
	}

	return strings.Join([]string{"UpdateAddressSetDto", string(data)}, " ")
}
