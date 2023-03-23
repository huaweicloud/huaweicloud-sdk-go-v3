package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdatePrivateDnatResponse struct {
	DnatRule *PrivateDnat `json:"dnat_rule,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdatePrivateDnatResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrivateDnatResponse struct{}"
	}

	return strings.Join([]string{"UpdatePrivateDnatResponse", string(data)}, " ")
}
