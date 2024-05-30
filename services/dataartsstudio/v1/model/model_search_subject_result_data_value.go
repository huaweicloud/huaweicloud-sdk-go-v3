package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchSubjectResultDataValue value，统一的返回结果的外层数据结构。
type SearchSubjectResultDataValue struct {

	// 总量。
	Total *int32 `json:"total,omitempty"`

	// 查询到的审批单对象（CatalogVO）数组。
	Records *[]CatalogVo `json:"records,omitempty"`
}

func (o SearchSubjectResultDataValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchSubjectResultDataValue struct{}"
	}

	return strings.Join([]string{"SearchSubjectResultDataValue", string(data)}, " ")
}
