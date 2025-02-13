package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNotifyPolicyResponse Response Object
type UpdateNotifyPolicyResponse struct {

	// 请求的唯一标识ID
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateNotifyPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNotifyPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateNotifyPolicyResponse", string(data)}, " ")
}
