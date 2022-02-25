package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StreamDetail struct {
	// 采样开始时间。日期格式按照ISO8601表示法，并使用UTC时间。 格式为：YYYY-MM-DDThh:mm:ssZ。

	StartTime string `json:"start_time"`
	// 采样结束时间。日期格式按照ISO8601表示法，并使用UTC时间。 格式为：YYYY-MM-DDThh:mm:ssZ。

	EndTime string `json:"end_time"`
	// 流监控数据列表。

	DataList []int64 `json:"data_list"`
}

func (o StreamDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StreamDetail struct{}"
	}

	return strings.Join([]string{"StreamDetail", string(data)}, " ")
}
