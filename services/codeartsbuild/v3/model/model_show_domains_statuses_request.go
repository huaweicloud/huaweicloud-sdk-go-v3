package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainsStatusesRequest Request Object
type ShowDomainsStatusesRequest struct {
	Body *DomainsStatusesRequestBody `json:"body,omitempty"`
}

func (o ShowDomainsStatusesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainsStatusesRequest struct{}"
	}

	return strings.Join([]string{"ShowDomainsStatusesRequest", string(data)}, " ")
}
