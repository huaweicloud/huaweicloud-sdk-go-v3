package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddressGroupVo struct {

	// 地址组id
	SetId *string `json:"set_id,omitempty"`

	// 地址组名称
	Name *string `json:"name,omitempty"`
}

func (o AddressGroupVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddressGroupVo struct{}"
	}

	return strings.Join([]string{"AddressGroupVo", string(data)}, " ")
}
