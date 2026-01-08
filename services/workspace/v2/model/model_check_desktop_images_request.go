package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckDesktopImagesRequest Request Object
type CheckDesktopImagesRequest struct {
	Body *CheckDesktopImagesReq `json:"body,omitempty"`
}

func (o CheckDesktopImagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckDesktopImagesRequest struct{}"
	}

	return strings.Join([]string{"CheckDesktopImagesRequest", string(data)}, " ")
}
