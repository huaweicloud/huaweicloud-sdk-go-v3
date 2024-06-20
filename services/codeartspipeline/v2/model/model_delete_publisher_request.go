package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePublisherRequest Request Object
type DeletePublisherRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	// 发布商ID
	PublisherUniqueId string `json:"publisher_unique_id"`
}

func (o DeletePublisherRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePublisherRequest struct{}"
	}

	return strings.Join([]string{"DeletePublisherRequest", string(data)}, " ")
}
