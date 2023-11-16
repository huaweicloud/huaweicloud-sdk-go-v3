package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAvailablePublisherRequest Request Object
type ListAvailablePublisherRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`
}

func (o ListAvailablePublisherRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailablePublisherRequest struct{}"
	}

	return strings.Join([]string{"ListAvailablePublisherRequest", string(data)}, " ")
}
