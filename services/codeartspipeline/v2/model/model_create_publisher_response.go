package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePublisherResponse Response Object
type CreatePublisherResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePublisherResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePublisherResponse struct{}"
	}

	return strings.Join([]string{"CreatePublisherResponse", string(data)}, " ")
}
