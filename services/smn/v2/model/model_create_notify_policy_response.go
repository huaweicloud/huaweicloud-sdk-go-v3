package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNotifyPolicyResponse Response Object
type CreateNotifyPolicyResponse struct {

	// 请求的唯一标识ID
	RequestId *string `json:"request_id,omitempty"`

	// 通知策略ID
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateNotifyPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNotifyPolicyResponse struct{}"
	}

	return strings.Join([]string{"CreateNotifyPolicyResponse", string(data)}, " ")
}
