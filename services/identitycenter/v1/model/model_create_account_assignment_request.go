package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAccountAssignmentRequest Request Object
type CreateAccountAssignmentRequest struct {
	InstanceId string `json:"instance_id"`

	Body *CreateAccountAssignmentReqBody `json:"body,omitempty"`
}

func (o CreateAccountAssignmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAccountAssignmentRequest struct{}"
	}

	return strings.Join([]string{"CreateAccountAssignmentRequest", string(data)}, " ")
}
