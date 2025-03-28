package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFlavorsDetailsRequest Request Object
type ListFlavorsDetailsRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// 数据库版本号。
	Version *string `json:"version,omitempty"`

	// 规格编码
	SpecCode *string `json:"spec_code,omitempty"`

	// 实例类型  集中式centralization_standard  分布式enterprise
	HaMode *string `json:"ha_mode,omitempty"`

	// 查询记录数。默认为100，不能为负数，最小值为1，最大值为100。
	Limit *int32 `json:"limit,omitempty"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListFlavorsDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlavorsDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListFlavorsDetailsRequest", string(data)}, " ")
}
