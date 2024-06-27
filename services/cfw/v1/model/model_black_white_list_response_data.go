package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BlackWhiteListResponseData 查询黑白名单返回值
type BlackWhiteListResponseData struct {

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示个数，范围为1-1024
	Limit *int32 `json:"limit,omitempty"`

	// 查询的总条数
	Total *int32 `json:"total,omitempty"`

	// 黑白名单记录
	Records *[]BlackWhiteListResponseDataRecords `json:"records,omitempty"`
}

func (o BlackWhiteListResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BlackWhiteListResponseData struct{}"
	}

	return strings.Join([]string{"BlackWhiteListResponseData", string(data)}, " ")
}
