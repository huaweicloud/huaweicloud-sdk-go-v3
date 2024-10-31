package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddressGroupVo struct {

	// 地址组类型，0表示自定义地址组，1表示WAF回源IP地址组，2表示DDoS回源IP地址组，3表示NAT64转换地址组
	AddressSetType *int32 `json:"address_set_type,omitempty"`

	// 关联IP地址组名称，可通过[查询地址组列表接口](ListAddressSets.xml)查询获得，通过返回值中的data.records.name（.表示各对象之间层级的区分）获得。
	Name *string `json:"name,omitempty"`

	// 关联IP地址组ID，可通过[查询地址组列表接口](ListAddressSets.xml)查询获得，通过返回值中的data.records.set_id（.表示各对象之间层级的区分）获得。
	SetId *string `json:"set_id,omitempty"`
}

func (o AddressGroupVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddressGroupVo struct{}"
	}

	return strings.Join([]string{"AddressGroupVo", string(data)}, " ")
}
