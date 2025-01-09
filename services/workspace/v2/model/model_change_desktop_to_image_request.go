package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeDesktopToImageRequest Request Object
type ChangeDesktopToImageRequest struct {
	Body *DesktopToImageReq `json:"body,omitempty"`
}

func (o ChangeDesktopToImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeDesktopToImageRequest struct{}"
	}

	return strings.Join([]string{"ChangeDesktopToImageRequest", string(data)}, " ")
}
