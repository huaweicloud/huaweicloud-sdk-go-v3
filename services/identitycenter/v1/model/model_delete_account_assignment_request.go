package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAccountAssignmentRequest Request Object
type DeleteAccountAssignmentRequest struct {
	InstanceId string `json:"instance_id"`

	Body *DeleteAccountAssignmentReqBody `json:"body,omitempty"`
}

func (o DeleteAccountAssignmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAccountAssignmentRequest struct{}"
	}

	return strings.Join([]string{"DeleteAccountAssignmentRequest", string(data)}, " ")
}
