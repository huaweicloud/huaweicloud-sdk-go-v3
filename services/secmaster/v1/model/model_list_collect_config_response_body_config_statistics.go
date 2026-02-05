package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCollectConfigResponseBodyConfigStatistics 统计数据
type ListCollectConfigResponseBodyConfigStatistics struct {

	// 账号接入数量
	AccountNum float32 `json:"account_num,omitempty"`

	// 每日流量，单位：Byte
	DailyTraffic float32 `json:"daily_traffic,omitempty"`

	// 已接入日志数量
	LogNum float32 `json:"log_num,omitempty"`

	// 云产品总数量
	ProductAllNum float32 `json:"product_all_num,omitempty"`

	// 接入云产品数量
	ProductInNum float32 `json:"product_in_num,omitempty"`

	// 云厂商数量
	VendorNum float32 `json:"vendor_num,omitempty"`
}

func (o ListCollectConfigResponseBodyConfigStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCollectConfigResponseBodyConfigStatistics struct{}"
	}

	return strings.Join([]string{"ListCollectConfigResponseBodyConfigStatistics", string(data)}, " ")
}
