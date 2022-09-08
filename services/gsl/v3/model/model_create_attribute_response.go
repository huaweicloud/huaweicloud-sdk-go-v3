package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateAttributeResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAttributeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAttributeResponse struct{}"
	}

	return strings.Join([]string{"CreateAttributeResponse", string(data)}, " ")
}
