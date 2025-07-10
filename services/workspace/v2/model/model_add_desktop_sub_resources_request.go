package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddDesktopSubResourcesRequest Request Object
type AddDesktopSubResourcesRequest struct {
	Body *AddDesktopSubResourcesReq `json:"body,omitempty"`
}

func (o AddDesktopSubResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDesktopSubResourcesRequest struct{}"
	}

	return strings.Join([]string{"AddDesktopSubResourcesRequest", string(data)}, " ")
}
