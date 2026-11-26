package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDataServiceApproversResponse Response Object
type ListDataServiceApproversResponse struct {

	// 总记录数
	TotalRecords *string `json:"total_records,omitempty"`

	// 单页记录数
	Size *int32 `json:"size,omitempty"`

	// 总页数
	TotalPages *int32 `json:"total_pages,omitempty"`

	// 审批人列表信息
	Approvers      *[]ApproverListApprovers `json:"approvers,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListDataServiceApproversResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataServiceApproversResponse struct{}"
	}

	return strings.Join([]string{"ListDataServiceApproversResponse", string(data)}, " ")
}
