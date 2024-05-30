package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTableModelRelationsResultDataValue value，统一的返回结果的外层数据结构。
type ListTableModelRelationsResultDataValue struct {

	// 总量。
	Total *int32 `json:"total,omitempty"`

	// RelationVO信息。
	Records *[]ListTableModelRelationsResultDataValueRecords `json:"records,omitempty"`
}

func (o ListTableModelRelationsResultDataValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTableModelRelationsResultDataValue struct{}"
	}

	return strings.Join([]string{"ListTableModelRelationsResultDataValue", string(data)}, " ")
}
