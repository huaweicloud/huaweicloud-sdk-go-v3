package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTableModelsResultDataValue value，统一的返回结果的外层数据结构。
type ListTableModelsResultDataValue struct {

	// 总量。
	Total *int32 `json:"total,omitempty"`

	// TableModelVO信息。
	Records *[]TableModelVo `json:"records,omitempty"`
}

func (o ListTableModelsResultDataValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTableModelsResultDataValue struct{}"
	}

	return strings.Join([]string{"ListTableModelsResultDataValue", string(data)}, " ")
}
