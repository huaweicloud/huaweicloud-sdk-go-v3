package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAccountAssignmentResponse Response Object
type DeleteAccountAssignmentResponse struct {
	AccountAssignmentDeletionStatus *AccountAssignmentOperationStatusDto `json:"account_assignment_deletion_status,omitempty"`
	HttpStatusCode                  int                                  `json:"-"`
}

func (o DeleteAccountAssignmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAccountAssignmentResponse struct{}"
	}

	return strings.Join([]string{"DeleteAccountAssignmentResponse", string(data)}, " ")
}
