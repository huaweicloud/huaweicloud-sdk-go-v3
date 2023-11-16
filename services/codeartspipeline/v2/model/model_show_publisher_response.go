package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPublisherResponse Response Object
type ShowPublisherResponse struct {
	Body           map[string]PublisherVo `json:"body,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowPublisherResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPublisherResponse struct{}"
	}

	return strings.Join([]string{"ShowPublisherResponse", string(data)}, " ")
}
