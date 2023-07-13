package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSetAttributesResponse Response Object
type BatchSetAttributesResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchSetAttributesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetAttributesResponse struct{}"
	}

	return strings.Join([]string{"BatchSetAttributesResponse", string(data)}, " ")
}
