package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateDependencyRequest struct {
	Body *CreateDependencyRequestBody `json:"body,omitempty"`
}

func (o CreateDependencyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDependencyRequest struct{}"
	}

	return strings.Join([]string{"CreateDependencyRequest", string(data)}, " ")
}
