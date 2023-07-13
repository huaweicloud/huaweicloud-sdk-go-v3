package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDependencyVersionRequest Request Object
type CreateDependencyVersionRequest struct {
	Body *CreateDependencyRequestBody `json:"body,omitempty"`
}

func (o CreateDependencyVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDependencyVersionRequest struct{}"
	}

	return strings.Join([]string{"CreateDependencyVersionRequest", string(data)}, " ")
}
