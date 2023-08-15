package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateComponentRequest Request Object
type CreateComponentRequest struct {
	Body *ComponentParam `json:"body,omitempty"`
}

func (o CreateComponentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateComponentRequest struct{}"
	}

	return strings.Join([]string{"CreateComponentRequest", string(data)}, " ")
}
