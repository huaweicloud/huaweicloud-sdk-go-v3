package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowPrivateNatResponse struct {
	Gateway *PrivateNat `json:"gateway,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPrivateNatResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPrivateNatResponse struct{}"
	}

	return strings.Join([]string{"ShowPrivateNatResponse", string(data)}, " ")
}
