package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPublisherRequest Request Object
type ListPublisherRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	// 偏移
	Offset string `json:"offset"`

	// 大小
	Limit string `json:"limit"`

	// 名称
	Name *string `json:"name,omitempty"`
}

func (o ListPublisherRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPublisherRequest struct{}"
	}

	return strings.Join([]string{"ListPublisherRequest", string(data)}, " ")
}
