package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ResizeDesktopRequest struct {
	Body *ResizeDesktopReq `json:"body,omitempty"`
}

func (o ResizeDesktopRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeDesktopRequest struct{}"
	}

	return strings.Join([]string{"ResizeDesktopRequest", string(data)}, " ")
}
