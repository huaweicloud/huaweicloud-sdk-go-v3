package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAssetTreeResponse Response Object
type ShowAssetTreeResponse struct {
	Code *string `json:"code,omitempty"`

	Data *interface{} `json:"data,omitempty"`

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAssetTreeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAssetTreeResponse struct{}"
	}

	return strings.Join([]string{"ShowAssetTreeResponse", string(data)}, " ")
}
