package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddressItemListResponseDtoData 查询地址组成员返回数据
type AddressItemListResponseDtoData struct {

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示个数，范围为1-1024
	Limit *int32 `json:"limit,omitempty"`

	// 地址组成员总数
	Total *int32 `json:"total,omitempty"`

	// 地址组id
	SetId *string `json:"set_id,omitempty"`

	// 地址组成员记录列表
	Records *[]AddressItemListResponseDtoDataRecords `json:"records,omitempty"`
}

func (o AddressItemListResponseDtoData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddressItemListResponseDtoData struct{}"
	}

	return strings.Join([]string{"AddressItemListResponseDtoData", string(data)}, " ")
}
