package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DescribeAccountAssignmentCreationStatusResponse Response Object
type DescribeAccountAssignmentCreationStatusResponse struct {
	AccountAssignmentCreationStatus *AccountAssignmentOperationStatusDto `json:"account_assignment_creation_status,omitempty"`
	HttpStatusCode                  int                                  `json:"-"`
}

func (o DescribeAccountAssignmentCreationStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DescribeAccountAssignmentCreationStatusResponse struct{}"
	}

	return strings.Join([]string{"DescribeAccountAssignmentCreationStatusResponse", string(data)}, " ")
}
