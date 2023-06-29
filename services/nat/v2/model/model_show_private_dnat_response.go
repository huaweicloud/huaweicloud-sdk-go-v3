package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPrivateDnatResponse Response Object
type ShowPrivateDnatResponse struct {
	DnatRule *PrivateDnat `json:"dnat_rule,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPrivateDnatResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPrivateDnatResponse struct{}"
	}

	return strings.Join([]string{"ShowPrivateDnatResponse", string(data)}, " ")
}
