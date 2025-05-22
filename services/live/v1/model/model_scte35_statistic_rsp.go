package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Scte35StatisticRsp 获取SCTE35信号的响应体。
type Scte35StatisticRsp struct {

	// 查询到scet35信息的总个数。
	Total int32 `json:"total"`

	// 详细的scte35信号的数组。
	Scte35Info *[]Scte35InfoItem `json:"scte35_info,omitempty"`
}

func (o Scte35StatisticRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Scte35StatisticRsp struct{}"
	}

	return strings.Join([]string{"Scte35StatisticRsp", string(data)}, " ")
}
