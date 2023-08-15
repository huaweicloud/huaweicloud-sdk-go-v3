package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WebImageImageSize
type WebImageImageSize struct {

	// 传入image_size时的返回，为图像高度。
	Height *int32 `json:"height,omitempty"`

	// 传入image_size时的返回，为图像宽度。
	Width *int32 `json:"width,omitempty"`
}

func (o WebImageImageSize) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebImageImageSize struct{}"
	}

	return strings.Join([]string{"WebImageImageSize", string(data)}, " ")
}
