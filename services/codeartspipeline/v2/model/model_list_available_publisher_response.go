package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAvailablePublisherResponse Response Object
type ListAvailablePublisherResponse struct {
	Body           *[]PublisherVo `json:"body,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListAvailablePublisherResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailablePublisherResponse struct{}"
	}

	return strings.Join([]string{"ListAvailablePublisherResponse", string(data)}, " ")
}
