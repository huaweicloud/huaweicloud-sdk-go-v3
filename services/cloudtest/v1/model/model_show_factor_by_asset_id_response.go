package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFactorByAssetIdResponse Response Object
type ShowFactorByAssetIdResponse struct {
	Code *string `json:"code,omitempty"`

	Data *interface{} `json:"data,omitempty"`

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowFactorByAssetIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFactorByAssetIdResponse struct{}"
	}

	return strings.Join([]string{"ShowFactorByAssetIdResponse", string(data)}, " ")
}
