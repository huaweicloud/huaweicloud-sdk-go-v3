package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 请求限速配置。
type RequestLimitRules struct {

	// 配置开关（on/off）
	Status string `json:"status"`

	// 限速方式，目前只支持按流量大小限速，size:大小。
	Type string `json:"type"`

	// 限速条件,type=size,limit_rate_after=50表示从传输表示传输50个字节后开始限速且限速值为limit_rate_value， 单位byte，取值范围：0-1073741824。
	LimitRateAfter *int64 `json:"limit_rate_after,omitempty"`

	// 限速值,单位Bps，取值范围 0-104857600
	LimitRateValue *int32 `json:"limit_rate_value,omitempty"`
}

func (o RequestLimitRules) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RequestLimitRules struct{}"
	}

	return strings.Join([]string{"RequestLimitRules", string(data)}, " ")
}
