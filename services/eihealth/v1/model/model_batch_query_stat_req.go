package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchQueryStatReq struct {

	// 查询监控数据起始时间，UNIX时间戳，单位毫秒，不填时默认为当前时间
	FromTime *int64 `json:"from_time,omitempty"`

	// 查询数据截止时间，UNIX时间戳，单位毫秒，不填时默认为当前时间
	ToTime *int64 `json:"to_time,omitempty"`

	Period *MonitorPeriod `json:"period,omitempty"`

	Method *MonitorMethod `json:"method,omitempty"`

	// 查询的监控资源对象id集合
	ResourceIds []string `json:"resource_ids"`
}

func (o BatchQueryStatReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchQueryStatReq struct{}"
	}

	return strings.Join([]string{"BatchQueryStatReq", string(data)}, " ")
}
