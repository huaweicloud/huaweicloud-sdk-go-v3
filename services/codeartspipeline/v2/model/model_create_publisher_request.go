package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePublisherRequest Request Object
type CreatePublisherRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	Body *PublisherRequest `json:"body,omitempty"`
}

func (o CreatePublisherRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePublisherRequest struct{}"
	}

	return strings.Join([]string{"CreatePublisherRequest", string(data)}, " ")
}
