package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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

	// 防护对象id，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志id，可通过调用查询防火墙实例接口获得，注意type为0的为互联网边界防护对象id，type为1的为VPC边界防护对象id。具体可参考APIExlorer和帮助中心FAQ。
	ObjectId *string `json:"object_id,omitempty"`

	// 地址组类型，0表示自定义地址组，1表示预定义地址组
	AddressSetType *int32 `json:"address_set_type,omitempty"`
}

func (o AddressSetListResponseDtoDataRecords) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddressSetListResponseDtoDataRecords struct{}"
	}

	return strings.Join([]string{"AddressSetListResponseDtoDataRecords", string(data)}, " ")
}
