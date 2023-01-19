package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TrendQueryData struct {

	// 查询结果。
	Result *string `json:"result,omitempty"`

	// 时间戳。
	Timestamp *int64 `json:"timestamp,omitempty"`
}

func (o TrendQueryData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrendQueryData struct{}"
	}

	return strings.Join([]string{"TrendQueryData", string(data)}, " ")
}
