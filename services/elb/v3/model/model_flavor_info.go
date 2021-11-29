package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 规格内容信息。
type FlavorInfo struct {
	// 并发数。

	Connection int32 `json:"connection"`
	// 新建数。

	Cps int32 `json:"cps"`
	// 7层每秒查询数。

	Qps *int32 `json:"qps,omitempty"`
	// 带宽。

	Bandwidth *int32 `json:"bandwidth,omitempty"`
	// flavor对应的lcu数量。

	Lcu *int32 `json:"lcu,omitempty"`
}

func (o FlavorInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorInfo struct{}"
	}

	return strings.Join([]string{"FlavorInfo", string(data)}, " ")
}
