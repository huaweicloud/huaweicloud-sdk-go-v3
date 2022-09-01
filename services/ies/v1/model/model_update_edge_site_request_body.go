package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateEdgeSiteRequestBody struct {
	EdgeSite *UpdateEdgeSite `json:"edge_site" xml:"edge_site"`
}

func (o UpdateEdgeSiteRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEdgeSiteRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateEdgeSiteRequestBody", string(data)}, " ")
}
