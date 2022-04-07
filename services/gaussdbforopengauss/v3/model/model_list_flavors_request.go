package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListFlavorsRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`
	// 数据库版本号。

	Version *string `json:"version,omitempty"`
	// 规格编码

	SpecCode *string `json:"spec_code,omitempty"`
	// 实例类型  集中式centralization_standard  分布式enterprise

	HaMode *string `json:"ha_mode,omitempty"`
	// 查询记录数。

	Limit *int32 `json:"limit,omitempty"`
	// 偏移量。

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListFlavorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlavorsRequest struct{}"
	}

	return strings.Join([]string{"ListFlavorsRequest", string(data)}, " ")
}
