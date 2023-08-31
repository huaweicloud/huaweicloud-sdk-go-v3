package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetPrimaryVideoThumbnailRequest Request Object
type SetPrimaryVideoThumbnailRequest struct {

	// 请求体参数类型，该字段必须设置为：application/json。
	ContentType string `json:"Content-Type"`

	Body *SetPrimaryVideoThumbnailRequestBody `json:"body,omitempty"`
}

func (o SetPrimaryVideoThumbnailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetPrimaryVideoThumbnailRequest struct{}"
	}

	return strings.Join([]string{"SetPrimaryVideoThumbnailRequest", string(data)}, " ")
}
