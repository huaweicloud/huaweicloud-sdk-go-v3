package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConnectionResponse Response Object
type CreateConnectionResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	Asset          *AssetInfoResponseBody `json:"asset,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o CreateConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConnectionResponse struct{}"
	}

	return strings.Join([]string{"CreateConnectionResponse", string(data)}, " ")
}
