package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableNat64Response Response Object
type DisableNat64Response struct {
	Publicip *PublicipResp `json:"publicip,omitempty"`

	// 本次请求编号
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DisableNat64Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableNat64Response struct{}"
	}

	return strings.Join([]string{"DisableNat64Response", string(data)}, " ")
}
