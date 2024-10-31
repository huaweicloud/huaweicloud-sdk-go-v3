package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServiceItemListResponseDtoData 服务组成员列表
type ServiceItemListResponseDtoData struct {

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示个数，范围为1-1024
	Limit *int32 `json:"limit,omitempty"`

	// 服务组成员总数
	Total *int32 `json:"total,omitempty"`

	// 服务组id
	SetId *string `json:"set_id,omitempty"`

	// 记录
	Records *[]ServiceItemListResponseDtoDataRecords `json:"records,omitempty"`
}

func (o ServiceItemListResponseDtoData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceItemListResponseDtoData struct{}"
	}

	return strings.Join([]string{"ServiceItemListResponseDtoData", string(data)}, " ")
}
