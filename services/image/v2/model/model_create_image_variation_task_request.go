package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateImageVariationTaskRequest struct {
	Body *CreateImageVariationTaskRequestBody `json:"body,omitempty"`
}

func (o CreateImageVariationTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateImageVariationTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateImageVariationTaskRequest", string(data)}, " ")
}
