package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResourceTagResponse Response Object
type CreateResourceTagResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateResourceTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceTagResponse struct{}"
	}

	return strings.Join([]string{"CreateResourceTagResponse", string(data)}, " ")
}
