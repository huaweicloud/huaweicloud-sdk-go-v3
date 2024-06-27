package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAssetResponse Response Object
type ShowAssetResponse struct {
	Code *string `json:"code,omitempty"`

	Data *[]Asset `json:"data,omitempty"`

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAssetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAssetResponse struct{}"
	}

	return strings.Join([]string{"ShowAssetResponse", string(data)}, " ")
}
