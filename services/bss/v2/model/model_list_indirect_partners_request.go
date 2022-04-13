package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListIndirectPartnersRequest struct {
	Body *QueryIndirectPartnersReq `json:"body,omitempty"`
}

func (o ListIndirectPartnersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIndirectPartnersRequest struct{}"
	}

	return strings.Join([]string{"ListIndirectPartnersRequest", string(data)}, " ")
}
