package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TopIpSummary top ip 详情数据
type TopIpSummary struct {

	// IP值。
	Ip *string `json:"ip,omitempty"`

	// 对应查询类型的值。（流量单位：Byte）
	Value *int64 `json:"value,omitempty"`

	// 该IP的流量(或请求数)占当前查询条件下总流量(或请求数)的比例。保留4位小数
	Ratio *float64 `json:"ratio,omitempty"`
}

func (o TopIpSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopIpSummary struct{}"
	}

	return strings.Join([]string{"TopIpSummary", string(data)}, " ")
}
