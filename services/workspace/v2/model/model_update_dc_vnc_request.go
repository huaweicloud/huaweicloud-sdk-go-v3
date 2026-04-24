package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDcVncRequest Request Object
type UpdateDcVncRequest struct {

	// 站点ID。
	SiteId string `json:"site_id"`

	Body *UpdateDcVncRequestBody `json:"body,omitempty"`
}

func (o UpdateDcVncRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDcVncRequest struct{}"
	}

	return strings.Join([]string{"UpdateDcVncRequest", string(data)}, " ")
}
