package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAccountAssignmentResponse Response Object
type CreateAccountAssignmentResponse struct {
	AccountAssignmentCreationStatus *AccountAssignmentOperationStatusDto `json:"account_assignment_creation_status,omitempty"`
	HttpStatusCode                  int                                  `json:"-"`
}

func (o CreateAccountAssignmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAccountAssignmentResponse struct{}"
	}

	return strings.Join([]string{"CreateAccountAssignmentResponse", string(data)}, " ")
}
