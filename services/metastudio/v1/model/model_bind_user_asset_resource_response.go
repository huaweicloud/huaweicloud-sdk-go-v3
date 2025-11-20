package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BindUserAssetResourceResponse Response Object
type BindUserAssetResourceResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BindUserAssetResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindUserAssetResourceResponse struct{}"
	}

	return strings.Join([]string{"BindUserAssetResourceResponse", string(data)}, " ")
}
