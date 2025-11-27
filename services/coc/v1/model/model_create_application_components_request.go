package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateApplicationComponentsRequest Request Object
type CreateApplicationComponentsRequest struct {
	Body *ComponentCreateRequest `json:"body,omitempty"`
}

func (o CreateApplicationComponentsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApplicationComponentsRequest struct{}"
	}

	return strings.Join([]string{"CreateApplicationComponentsRequest", string(data)}, " ")
}
