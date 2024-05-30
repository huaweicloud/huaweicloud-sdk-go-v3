package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApproversResultDataValue value，统一的返回结果的外层数据结构。
type ListApproversResultDataValue struct {

	// 总量。
	Total *int32 `json:"total,omitempty"`

	// 查询到的审批人对象（ApproverVO）数组。
	Records *[]ApproverVo `json:"records,omitempty"`
}

func (o ListApproversResultDataValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApproversResultDataValue struct{}"
	}

	return strings.Join([]string{"ListApproversResultDataValue", string(data)}, " ")
}
