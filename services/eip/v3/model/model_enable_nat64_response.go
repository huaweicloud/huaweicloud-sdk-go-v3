package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableNat64Response Response Object
type EnableNat64Response struct {
	Publicip *PublicipResp `json:"publicip,omitempty"`

	// 本次请求编号
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o EnableNat64Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableNat64Response struct{}"
	}

	return strings.Join([]string{"EnableNat64Response", string(data)}, " ")
}
