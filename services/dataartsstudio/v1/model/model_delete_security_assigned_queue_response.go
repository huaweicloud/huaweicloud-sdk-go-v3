package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSecurityAssignedQueueResponse Response Object
type DeleteSecurityAssignedQueueResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSecurityAssignedQueueResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecurityAssignedQueueResponse struct{}"
	}

	return strings.Join([]string{"DeleteSecurityAssignedQueueResponse", string(data)}, " ")
}
