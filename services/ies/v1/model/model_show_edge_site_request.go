package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEdgeSiteRequest Request Object
type ShowEdgeSiteRequest struct {

	// 边缘小站ID
	SiteId string `json:"site_id"`
}

func (o ShowEdgeSiteRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEdgeSiteRequest struct{}"
	}

	return strings.Join([]string{"ShowEdgeSiteRequest", string(data)}, " ")
}
