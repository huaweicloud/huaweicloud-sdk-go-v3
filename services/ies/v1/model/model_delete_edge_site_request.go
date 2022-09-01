package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteEdgeSiteRequest struct {

	// 边缘小站ID
	SiteId string `json:"site_id" xml:"site_id"`
}

func (o DeleteEdgeSiteRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEdgeSiteRequest struct{}"
	}

	return strings.Join([]string{"DeleteEdgeSiteRequest", string(data)}, " ")
}
