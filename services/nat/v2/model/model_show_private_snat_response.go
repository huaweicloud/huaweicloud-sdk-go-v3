package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowPrivateSnatResponse struct {
	SnatRule *PrivateSnat `json:"snat_rule,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPrivateSnatResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPrivateSnatResponse struct{}"
	}

	return strings.Join([]string{"ShowPrivateSnatResponse", string(data)}, " ")
}
