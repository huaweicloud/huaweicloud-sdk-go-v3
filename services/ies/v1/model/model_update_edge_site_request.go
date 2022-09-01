package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateEdgeSiteRequest struct {

	// 边缘小站ID
	SiteId string `json:"site_id" xml:"site_id"`

	Body *UpdateEdgeSiteRequestBody `json:"body,omitempty" xml:"body"`
}

func (o UpdateEdgeSiteRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEdgeSiteRequest struct{}"
	}

	return strings.Join([]string{"UpdateEdgeSiteRequest", string(data)}, " ")
}
