package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdatePortalInfoRequest struct {

	// 主页ID。
	PortalId string `json:"portal_id"`

	Body *UpdatePortalInfoRequestBody `json:"body,omitempty"`
}

func (o UpdatePortalInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePortalInfoRequest struct{}"
	}

	return strings.Join([]string{"UpdatePortalInfoRequest", string(data)}, " ")
}
