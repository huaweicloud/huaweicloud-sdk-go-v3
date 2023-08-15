package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDisassociatePublicipResponse Response Object
type UpdateDisassociatePublicipResponse struct {

	// 本次请求的编号
	RequestId *string `json:"request_id,omitempty"`

	Publicip       *PublicipInstanceResp `json:"publicip,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o UpdateDisassociatePublicipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDisassociatePublicipResponse struct{}"
	}

	return strings.Join([]string{"UpdateDisassociatePublicipResponse", string(data)}, " ")
}
