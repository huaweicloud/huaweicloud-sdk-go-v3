package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchCustomizedFieldsResultDataValue value，统一的返回结果的外层数据结构。
type SearchCustomizedFieldsResultDataValue struct {

	// 总量。
	Total *int32 `json:"total,omitempty"`

	// CustomizedFieldsVO信息。
	Records *[]CustomizedFieldsVo `json:"records,omitempty"`
}

func (o SearchCustomizedFieldsResultDataValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchCustomizedFieldsResultDataValue struct{}"
	}

	return strings.Join([]string{"SearchCustomizedFieldsResultDataValue", string(data)}, " ")
}
