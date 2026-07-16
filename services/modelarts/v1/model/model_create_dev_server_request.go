package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDevServerRequest Request Object
type CreateDevServerRequest struct {
	Body *ServerCreateRequest `json:"body,omitempty"`
}

func (o CreateDevServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDevServerRequest struct{}"
	}

	return strings.Join([]string{"CreateDevServerRequest", string(data)}, " ")
}
