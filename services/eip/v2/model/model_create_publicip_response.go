package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreatePublicipResponse struct {
	Publicip       *PublicipCreateResp `json:"publicip,omitempty" xml:"publicip"`
	HttpStatusCode int                 `json:"-"`
}

func (o CreatePublicipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePublicipResponse struct{}"
	}

	return strings.Join([]string{"CreatePublicipResponse", string(data)}, " ")
}
