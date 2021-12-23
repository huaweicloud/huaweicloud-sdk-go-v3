package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateResourceRequest struct {
	Body *ResourceInfo `json:"body,omitempty"`
}

func (o CreateResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceRequest struct{}"
	}

	return strings.Join([]string{"CreateResourceRequest", string(data)}, " ")
}
