package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateConnectionResponse Response Object
type UpdateConnectionResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	Asset          *AssetInfoResponseBody `json:"asset,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o UpdateConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConnectionResponse struct{}"
	}

	return strings.Join([]string{"UpdateConnectionResponse", string(data)}, " ")
}
