package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchDeleteResourceTagResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchDeleteResourceTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteResourceTagResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteResourceTagResponse", string(data)}, " ")
}
