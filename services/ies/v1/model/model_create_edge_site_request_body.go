package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateEdgeSiteRequestBody struct {
	EdgeSite *CreateEdgeSite `json:"edge_site" xml:"edge_site"`
}

func (o CreateEdgeSiteRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEdgeSiteRequestBody struct{}"
	}

	return strings.Join([]string{"CreateEdgeSiteRequestBody", string(data)}, " ")
}
