package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRelationsResultDataValue value，统一的返回结果的外层数据结构。
type ListRelationsResultDataValue struct {

	// 总量。
	Total *int32 `json:"total,omitempty"`

	// RelationVO信息。
	Records *[]RelationVo `json:"records,omitempty"`
}

func (o ListRelationsResultDataValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRelationsResultDataValue struct{}"
	}

	return strings.Join([]string{"ListRelationsResultDataValue", string(data)}, " ")
}
