package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AssociatePublicipsResponse struct {

	// 本次请求的编号
	RequestId *string `json:"request_id,omitempty" xml:"request_id"`

	Publicip       *PublicipInstanceResp `json:"publicip,omitempty" xml:"publicip"`
	HttpStatusCode int                   `json:"-"`
}

func (o AssociatePublicipsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociatePublicipsResponse struct{}"
	}

	return strings.Join([]string{"AssociatePublicipsResponse", string(data)}, " ")
}
