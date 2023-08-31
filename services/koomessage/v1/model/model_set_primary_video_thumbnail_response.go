package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetPrimaryVideoThumbnailResponse Response Object
type SetPrimaryVideoThumbnailResponse struct {

	// 请求状态，固定200。
	Status *string `json:"status,omitempty"`

	// 状态描述。
	Message *string `json:"message,omitempty"`

	// 固定为null。
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o SetPrimaryVideoThumbnailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetPrimaryVideoThumbnailResponse struct{}"
	}

	return strings.Join([]string{"SetPrimaryVideoThumbnailResponse", string(data)}, " ")
}
