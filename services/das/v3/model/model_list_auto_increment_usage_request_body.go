package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListAutoIncrementUsageRequestBody struct {

	// 自增 ID 使用比例过滤值，只返回超过该比例的自增 ID 使用数据。取值为 0~1 的小数，默认为0
	Ratio string `json:"ratio"`

	// 是否获取实时数据： true：实时查询实例上数据并计算。最小查询时间粒度为 10 分钟，即有 10 分钟内的数据时，即使传递 true 也不进行实时查询。 false：当有近两小时的数据时，返回该数据，否则查询实例上最新数据。
	RealTime bool `json:"real_time"`

	// 偏移值
	Offsite *string `json:"offsite,omitempty"`

	// 每次获取的数量
	Limit *string `json:"limit,omitempty"`

	// 数据库名。传入此参数时，查询指定数据库中表自增 ID 使用情况，不传入时查询实例上所有数据库的表自增 ID 使用情况。
	Database *[]string `json:"database,omitempty"`
}

func (o ListAutoIncrementUsageRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAutoIncrementUsageRequestBody struct{}"
	}

	return strings.Join([]string{"ListAutoIncrementUsageRequestBody", string(data)}, " ")
}
