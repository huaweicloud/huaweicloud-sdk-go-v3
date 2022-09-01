package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateEdgeAppRequest struct {
	Body *CreateEdgeApplicationRequestDto `json:"body,omitempty" xml:"body"`
}

func (o CreateEdgeAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEdgeAppRequest struct{}"
	}

	return strings.Join([]string{"CreateEdgeAppRequest", string(data)}, " ")
}
