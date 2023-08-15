package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPublicIpTypeResponse Response Object
type ShowPublicIpTypeResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPublicIpTypeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPublicIpTypeResponse struct{}"
	}

	return strings.Join([]string{"ShowPublicIpTypeResponse", string(data)}, " ")
}
