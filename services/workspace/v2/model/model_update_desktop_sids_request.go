package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDesktopSidsRequest Request Object
type UpdateDesktopSidsRequest struct {
	Body *UpdateDesktopSidReq `json:"body,omitempty"`
}

func (o UpdateDesktopSidsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDesktopSidsRequest struct{}"
	}

	return strings.Join([]string{"UpdateDesktopSidsRequest", string(data)}, " ")
}
