package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowTrafficRequest struct {
	// 播放域名，不指定域名表示查询租户所有域名汇总流量

	Domain *string `json:"domain,omitempty"`
	// 查询起始时间，UTC时间，格式：yyyy-MM-ddTHH:mm:ssZ

	StartTime *string `json:"start_time,omitempty"`
	// 查询结束时间，UTC时间，格式：yyyy-MM-ddTHH:mm:ssZ。  - start_time与end_time均不存在时，服务端从最近一个统计周期的数据里查询。 - start_time存在、end_time不存在时，end_time取当前时间。 - start_time不存在、end_time存在时，请求非法。 - 只能查询最近三个月内的数据，start_time和end_time的跨度不能大于30天。

	EndTime *string `json:"end_time,omitempty"`
	// 统计周期，单位分钟

	Step *int32 `json:"step,omitempty"`
}

func (o ShowTrafficRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowTrafficRequest struct{}"
	}

	return strings.Join([]string{"ShowTrafficRequest", string(data)}, " ")
}
