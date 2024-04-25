package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMetricsDataResponse Response Object
type ListMetricsDataResponse struct {

	// 响应码
	Code *int32 `json:"code,omitempty"`

	// 响应信息
	Msg *string `json:"msg,omitempty"`

	// 指标采集数据列表。
	Data *[]map[string]interface{} `json:"data,omitempty"`

	// 总列表大小。
	Count          *int64 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListMetricsDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMetricsDataResponse struct{}"
	}

	return strings.Join([]string{"ListMetricsDataResponse", string(data)}, " ")
}
