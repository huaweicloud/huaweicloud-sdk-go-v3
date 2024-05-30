package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchApprovalsResultDataValue value，统一的返回结果的外层数据结构。
type SearchApprovalsResultDataValue struct {

	// 总量。
	Total *int32 `json:"total,omitempty"`

	// 查询到的审批单对象（ApprovalVO）数组。
	Records *[]ApprovalVo `json:"records,omitempty"`
}

func (o SearchApprovalsResultDataValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchApprovalsResultDataValue struct{}"
	}

	return strings.Join([]string{"SearchApprovalsResultDataValue", string(data)}, " ")
}
