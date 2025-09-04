package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TopOriginUrlSummary top origin url 详情数据
type TopOriginUrlSummary struct {

	// 回源url名称
	OriginUrl *string `json:"origin_url,omitempty"`

	// 对应查询类型的值。（流量单位：Byte）
	Value *int64 `json:"value,omitempty"`

	// 该origin url的流量(或请求数)占当前查询条件下总流量(或请求数)的比例。保留4位小数
	Ratio *float64 `json:"ratio,omitempty"`
}

func (o TopOriginUrlSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopOriginUrlSummary struct{}"
	}

	return strings.Join([]string{"TopOriginUrlSummary", string(data)}, " ")
}
