package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckEdgeSiteResourcesRequest Request Object
type CheckEdgeSiteResourcesRequest struct {
	Body *CheckEdgeSiteResourcesReq `json:"body,omitempty"`
}

func (o CheckEdgeSiteResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckEdgeSiteResourcesRequest struct{}"
	}

	return strings.Join([]string{"CheckEdgeSiteResourcesRequest", string(data)}, " ")
}
