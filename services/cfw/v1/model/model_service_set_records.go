package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServiceSetRecords struct {

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示个数
	Limit *int32 `json:"limit,omitempty"`

	// 查询总数
	Total *int32 `json:"total,omitempty"`

	// 服务组列表
	Records *[]ServiceSet `json:"records,omitempty"`
}

func (o ServiceSetRecords) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceSetRecords struct{}"
	}

	return strings.Join([]string{"ServiceSetRecords", string(data)}, " ")
}
