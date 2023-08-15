package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEdgeAppRequest Request Object
type CreateEdgeAppRequest struct {
	Body *CreateEdgeApplicationRequestDto `json:"body,omitempty"`
}

func (o CreateEdgeAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEdgeAppRequest struct{}"
	}

	return strings.Join([]string{"CreateEdgeAppRequest", string(data)}, " ")
}
