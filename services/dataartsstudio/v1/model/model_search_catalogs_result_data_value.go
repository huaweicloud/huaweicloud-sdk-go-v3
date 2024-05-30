package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchCatalogsResultDataValue value，统一的返回结果的外层数据结构。
type SearchCatalogsResultDataValue struct {

	// 总量。
	Total *int32 `json:"total,omitempty"`

	// 流程架构详情。
	Records *[]BizCatalogVo `json:"records,omitempty"`
}

func (o SearchCatalogsResultDataValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchCatalogsResultDataValue struct{}"
	}

	return strings.Join([]string{"SearchCatalogsResultDataValue", string(data)}, " ")
}
