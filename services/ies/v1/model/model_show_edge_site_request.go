package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowEdgeSiteRequest struct {

	// 边缘小站ID
	SiteId string `json:"site_id" xml:"site_id"`
}

func (o ShowEdgeSiteRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEdgeSiteRequest struct{}"
	}

	return strings.Join([]string{"ShowEdgeSiteRequest", string(data)}, " ")
}
