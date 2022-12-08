package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchAddResourceTagResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchAddResourceTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddResourceTagResponse struct{}"
	}

	return strings.Join([]string{"BatchAddResourceTagResponse", string(data)}, " ")
}
