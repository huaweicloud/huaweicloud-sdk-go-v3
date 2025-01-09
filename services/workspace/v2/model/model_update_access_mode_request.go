package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAccessModeRequest Request Object
type UpdateAccessModeRequest struct {

	// 站点ID。
	SiteId string `json:"site_id"`

	Body *UpdateAccessModeReq `json:"body,omitempty"`
}

func (o UpdateAccessModeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAccessModeRequest struct{}"
	}

	return strings.Join([]string{"UpdateAccessModeRequest", string(data)}, " ")
}
