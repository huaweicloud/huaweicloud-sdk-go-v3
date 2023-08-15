package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePublicBandWidthResponse Response Object
type UpdatePublicBandWidthResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdatePublicBandWidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePublicBandWidthResponse struct{}"
	}

	return strings.Join([]string{"UpdatePublicBandWidthResponse", string(data)}, " ")
}
