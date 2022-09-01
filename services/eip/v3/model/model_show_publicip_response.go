package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowPublicipResponse struct {

	// 本次请求的编号
	RequestId *string `json:"request_id,omitempty" xml:"request_id"`

	Publicip       *PublicipSingleShowResp `json:"publicip,omitempty" xml:"publicip"`
	HttpStatusCode int                     `json:"-"`
}

func (o ShowPublicipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPublicipResponse struct{}"
	}

	return strings.Join([]string{"ShowPublicipResponse", string(data)}, " ")
}
