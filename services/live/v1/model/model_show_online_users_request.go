package model

import (
	"encoding/json"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"

	"strings"
)

// Request Object
type ShowOnlineUsersRequest struct {
	// 直播播放域名

	Domain string `json:"domain"`
	// 应用名称。 默认为“live”，若您需要自定义应用名称，请先提交工单申请。

	AppName *string `json:"app_name,omitempty"`
	// 流名称

	StreamName *string `json:"stream_name,omitempty"`
	// 查询开始时间，UTC时间，格式：yyyy-mm-ddThh:mm:ssZ。无开始时间表示查询最近统计周期在线人数数据

	StartTime *sdktime.SdkTime `json:"start_time,omitempty"`
	// 查询结束时间，UTC时间，格式：yyyy-MM-ddTHH:mm:ssZ。  - start_time与end_time均不存在时，服务端从最近一个统计周期的数据里查询。 - start_time存在、end_time不存在时，end_time取当前时间。 - start_time不存在、end_time存在时，请求非法。 - 只能查询最近三个月内的数据，start_time和end_time的跨度不能大于30天。

	EndTime *sdktime.SdkTime `json:"end_time,omitempty"`
	// 统计周期。 单位：分钟

	Step *int32 `json:"step,omitempty"`
}

func (o ShowOnlineUsersRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowOnlineUsersRequest struct{}"
	}

	return strings.Join([]string{"ShowOnlineUsersRequest", string(data)}, " ")
}
