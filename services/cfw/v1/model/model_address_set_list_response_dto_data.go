package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询地址组列表返回数据
type AddressSetListResponseDtoData struct {

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示个数
	Limit *int32 `json:"limit,omitempty"`

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 地址组列表
	Records *[]AddressSetListResponseDtoDataRecords `json:"records,omitempty"`
}

func (o AddressSetListResponseDtoData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddressSetListResponseDtoData struct{}"
	}

	return strings.Join([]string{"AddressSetListResponseDtoData", string(data)}, " ")
}
