package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddressGroupVo struct {

	// 地址组类型，0表示自定义地址组，1表示WAF回源IP地址组，2表示DDoS回源IP地址组，3表示NAT64转换地址组
	AddressSetType *int32 `json:"address_set_type,omitempty"`

	// 地址组名称
	Name *string `json:"name,omitempty"`

	// 地址组id
	SetId *string `json:"set_id,omitempty"`
}

func (o AddressGroupVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddressGroupVo struct{}"
	}

	return strings.Join([]string{"AddressGroupVo", string(data)}, " ")
}
