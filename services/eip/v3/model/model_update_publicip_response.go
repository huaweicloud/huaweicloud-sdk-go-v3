package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePublicipResponse Response Object
type UpdatePublicipResponse struct {

	// 本次请求的编号
	RequestId *string `json:"request_id,omitempty"`

	Publicip       *PublicipUpdateResp `json:"publicip,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o UpdatePublicipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePublicipResponse struct{}"
	}

	return strings.Join([]string{"UpdatePublicipResponse", string(data)}, " ")
}
