package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdatePublicIpResponse struct {
	Publicip       *PublicIp `json:"publicip,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o UpdatePublicIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePublicIpResponse struct{}"
	}

	return strings.Join([]string{"UpdatePublicIpResponse", string(data)}, " ")
}
