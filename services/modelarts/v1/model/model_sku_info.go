package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SkuInfo 计费信息。
type SkuInfo struct {

	// 计费码。
	Code *string `json:"code,omitempty"`

	// 计费时期。
	Period *string `json:"period,omitempty"`

	// 查询次数。
	QueriesLimit *int64 `json:"queries_limit,omitempty"`

	// 价格。
	Price *float32 `json:"price,omitempty"`
}

func (o SkuInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SkuInfo struct{}"
	}

	return strings.Join([]string{"SkuInfo", string(data)}, " ")
}
