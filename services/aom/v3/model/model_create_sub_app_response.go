package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSubAppResponse Response Object
type CreateSubAppResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSubAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubAppResponse struct{}"
	}

	return strings.Join([]string{"CreateSubAppResponse", string(data)}, " ")
}
