package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPublicipResponse Response Object
type ShowPublicipResponse struct {

	// 本次请求的编号
	RequestId *string `json:"request_id,omitempty"`

	Publicip       *PublicipSingleShowResp `json:"publicip,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ShowPublicipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPublicipResponse struct{}"
	}

	return strings.Join([]string{"ShowPublicipResponse", string(data)}, " ")
}
