package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDesktopPoolRequest Request Object
type CreateDesktopPoolRequest struct {
	Body *CreateDesktopPoolReq `json:"body,omitempty"`
}

func (o CreateDesktopPoolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDesktopPoolRequest struct{}"
	}

	return strings.Join([]string{"CreateDesktopPoolRequest", string(data)}, " ")
}
