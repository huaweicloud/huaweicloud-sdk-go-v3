package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchSubjectNewResultDataValue value，统一的返回结果的外层数据结构。
type SearchSubjectNewResultDataValue struct {

	// 总量。
	Total *int32 `json:"total,omitempty"`

	// CatalogVO信息。
	Records *[]CatalogVo `json:"records,omitempty"`
}

func (o SearchSubjectNewResultDataValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchSubjectNewResultDataValue struct{}"
	}

	return strings.Join([]string{"SearchSubjectNewResultDataValue", string(data)}, " ")
}
