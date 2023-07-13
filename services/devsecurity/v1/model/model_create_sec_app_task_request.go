package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSecAppTaskRequest Request Object
type CreateSecAppTaskRequest struct {
	Body *CreateSecAppTaskRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o CreateSecAppTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecAppTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateSecAppTaskRequest", string(data)}, " ")
}
