package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CatalogueBatchPojo 批量编辑目录pojo
type CatalogueBatchPojo struct {

	// 是否显示
	IsDisplay *bool `json:"is_display,omitempty"`

	// 目录编码
	Id *string `json:"id,omitempty"`

	// 二级目录中文别名
	SecondAliasZh *string `json:"second_alias_zh,omitempty"`

	// 二级目录中文别名
	SecondAliasEn *string `json:"second_alias_en,omitempty"`
}

func (o CatalogueBatchPojo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CatalogueBatchPojo struct{}"
	}

	return strings.Join([]string{"CatalogueBatchPojo", string(data)}, " ")
}
