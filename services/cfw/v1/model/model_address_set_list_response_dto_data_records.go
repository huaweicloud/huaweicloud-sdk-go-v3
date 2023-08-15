package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddressSetListResponseDtoDataRecords 列表信息
type AddressSetListResponseDtoDataRecords struct {

	// 地址组id
	SetId *string `json:"set_id,omitempty"`

	// 引用次数
	RefCount *int32 `json:"ref_count,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 地址组名称
	Name *string `json:"name,omitempty"`

	// 地址类型0 ipv4,1 ipv6
	AddressType *int32 `json:"address_type,omitempty"`
}

func (o AddressSetListResponseDtoDataRecords) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddressSetListResponseDtoDataRecords struct{}"
	}

	return strings.Join([]string{"AddressSetListResponseDtoDataRecords", string(data)}, " ")
}
