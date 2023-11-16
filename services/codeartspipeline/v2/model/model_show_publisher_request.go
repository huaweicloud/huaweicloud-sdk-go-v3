package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPublisherRequest Request Object
type ShowPublisherRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	Body *[]string `json:"body,omitempty"`
}

func (o ShowPublisherRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPublisherRequest struct{}"
	}

	return strings.Join([]string{"ShowPublisherRequest", string(data)}, " ")
}
