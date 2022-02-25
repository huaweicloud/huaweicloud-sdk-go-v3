package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateStructConfigResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateStructConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStructConfigResponse struct{}"
	}

	return strings.Join([]string{"CreateStructConfigResponse", string(data)}, " ")
}
