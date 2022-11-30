package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateAssociatePublicipResponse struct {

	// 本次请求的编号
	RequestId *string `json:"request_id,omitempty"`

	Publicip       *PublicipInstanceResp `json:"publicip,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o UpdateAssociatePublicipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAssociatePublicipResponse struct{}"
	}

	return strings.Join([]string{"UpdateAssociatePublicipResponse", string(data)}, " ")
}
