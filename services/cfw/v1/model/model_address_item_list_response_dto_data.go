package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询地址组成员返回数据
type AddressItemListResponseDtoData struct {

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示个数
	Limit *int32 `json:"limit,omitempty"`

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 地址组id
	SetId *string `json:"set_id,omitempty"`

	// 成员信息
	Records *[]AddressItemListResponseDtoDataRecords `json:"records,omitempty"`
}

func (o AddressItemListResponseDtoData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddressItemListResponseDtoData struct{}"
	}

	return strings.Join([]string{"AddressItemListResponseDtoData", string(data)}, " ")
}
