package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConnectionResponse Response Object
type ShowConnectionResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	Asset          *AssetInfoResponseBody `json:"asset,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConnectionResponse struct{}"
	}

	return strings.Join([]string{"ShowConnectionResponse", string(data)}, " ")
}
