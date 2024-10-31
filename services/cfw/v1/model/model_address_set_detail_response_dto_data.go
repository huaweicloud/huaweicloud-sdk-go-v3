package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddressSetDetailResponseDtoData 查询地址组详情数据
type AddressSetDetailResponseDtoData struct {

	// 地址组id
	Id *string `json:"id,omitempty"`

	// 地址组名称
	Name *string `json:"name,omitempty"`

	// 地址组描述
	Description *string `json:"description,omitempty"`

	// 地址组类型，0表示自定义地址组，1表示WAF回源IP地址组，2表示DDoS回源IP地址组，3表示NAT64转换地址组
	AddressSetType *int32 `json:"address_set_type,omitempty"`

	// 地址类型0 ipv4，1 ipv6
	AddressType *int32 `json:"address_type,omitempty"`
}

func (o AddressSetDetailResponseDtoData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddressSetDetailResponseDtoData struct{}"
	}

	return strings.Join([]string{"AddressSetDetailResponseDtoData", string(data)}, " ")
}
