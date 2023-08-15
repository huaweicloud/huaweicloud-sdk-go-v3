package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePrivateNatResponse Response Object
type CreatePrivateNatResponse struct {
	Gateway *PrivateNat `json:"gateway,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePrivateNatResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateNatResponse struct{}"
	}

	return strings.Join([]string{"CreatePrivateNatResponse", string(data)}, " ")
}
