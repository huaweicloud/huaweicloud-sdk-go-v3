package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIndirectPartnersRequest Request Object
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
