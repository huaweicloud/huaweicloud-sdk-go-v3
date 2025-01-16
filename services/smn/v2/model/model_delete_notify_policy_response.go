package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNotifyPolicyResponse Response Object
type DeleteNotifyPolicyResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteNotifyPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNotifyPolicyResponse struct{}"
	}

	return strings.Join([]string{"DeleteNotifyPolicyResponse", string(data)}, " ")
}
