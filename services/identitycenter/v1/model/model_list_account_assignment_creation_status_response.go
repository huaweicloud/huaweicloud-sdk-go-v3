package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAccountAssignmentCreationStatusResponse Response Object
type ListAccountAssignmentCreationStatusResponse struct {

	// 操作状态列表.
	AccountAssignmentsCreationStatus *[]AccountAssignmentOperationStatusMetadataDto `json:"account_assignments_creation_status,omitempty"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListAccountAssignmentCreationStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccountAssignmentCreationStatusResponse struct{}"
	}

	return strings.Join([]string{"ListAccountAssignmentCreationStatusResponse", string(data)}, " ")
}
