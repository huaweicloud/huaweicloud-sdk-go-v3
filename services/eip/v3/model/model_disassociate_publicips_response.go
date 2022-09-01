package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DisassociatePublicipsResponse struct {

	// 本次请求的编号
	RequestId *string `json:"request_id,omitempty" xml:"request_id"`

	Publicip       *PublicipInstanceResp `json:"publicip,omitempty" xml:"publicip"`
	HttpStatusCode int                   `json:"-"`
}

func (o DisassociatePublicipsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociatePublicipsResponse struct{}"
	}

	return strings.Join([]string{"DisassociatePublicipsResponse", string(data)}, " ")
}
