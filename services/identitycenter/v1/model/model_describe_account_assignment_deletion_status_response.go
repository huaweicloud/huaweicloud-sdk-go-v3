package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DescribeAccountAssignmentDeletionStatusResponse Response Object
type DescribeAccountAssignmentDeletionStatusResponse struct {
	AccountAssignmentDeletionStatus *AccountAssignmentOperationStatusDto `json:"account_assignment_deletion_status,omitempty"`
	HttpStatusCode                  int                                  `json:"-"`
}

func (o DescribeAccountAssignmentDeletionStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DescribeAccountAssignmentDeletionStatusResponse struct{}"
	}

	return strings.Join([]string{"DescribeAccountAssignmentDeletionStatusResponse", string(data)}, " ")
}
