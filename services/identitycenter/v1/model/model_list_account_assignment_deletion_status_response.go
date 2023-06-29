package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAccountAssignmentDeletionStatusResponse Response Object
type ListAccountAssignmentDeletionStatusResponse struct {

	// 操作状态集合
	AccountAssignmentsDeletionStatus *[]AccountAssignmentOperationStatusMetadataDto `json:"account_assignments_deletion_status,omitempty"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListAccountAssignmentDeletionStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccountAssignmentDeletionStatusResponse struct{}"
	}

	return strings.Join([]string{"ListAccountAssignmentDeletionStatusResponse", string(data)}, " ")
}
