package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateProjectRequest Request Object
type CreateProjectRequest struct {
	Body *CreateProjectRequestBody `json:"body,omitempty"`
}

func (o CreateProjectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProjectRequest struct{}"
	}

	return strings.Join([]string{"CreateProjectRequest", string(data)}, " ")
}
