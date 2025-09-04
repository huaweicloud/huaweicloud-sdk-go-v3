package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TopUaSummary top ua 详情数据
type TopUaSummary struct {

	// UA值。
	Ua *string `json:"ua,omitempty"`

	// 对应查询类型的值。（流量单位：Byte）
	Value *int64 `json:"value,omitempty"`

	// 该ua的流量(或请求数)占当前查询条件下总流量(或请求数)的比例。保留4位小数
	Ratio *float64 `json:"ratio,omitempty"`
}

func (o TopUaSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopUaSummary struct{}"
	}

	return strings.Join([]string{"TopUaSummary", string(data)}, " ")
}
