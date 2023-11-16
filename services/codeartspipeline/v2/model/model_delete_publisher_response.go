package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePublisherResponse Response Object
type DeletePublisherResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeletePublisherResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePublisherResponse struct{}"
	}

	return strings.Join([]string{"DeletePublisherResponse", string(data)}, " ")
}
