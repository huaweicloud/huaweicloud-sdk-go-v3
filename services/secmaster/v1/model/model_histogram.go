package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Histogram struct {

	// 当前时间段内的数据条数
	Count int64 `json:"count"`

	// 开始时间点
	From int64 `json:"from"`

	// 结束时间点
	To int64 `json:"to"`
}

func (o Histogram) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Histogram struct{}"
	}

	return strings.Join([]string{"Histogram", string(data)}, " ")
}
