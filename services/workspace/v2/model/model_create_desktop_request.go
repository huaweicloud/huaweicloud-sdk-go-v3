package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDesktopRequest Request Object
type CreateDesktopRequest struct {
	Body *CreateDesktopReq `json:"body,omitempty"`
}

func (o CreateDesktopRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDesktopRequest struct{}"
	}

	return strings.Join([]string{"CreateDesktopRequest", string(data)}, " ")
}
