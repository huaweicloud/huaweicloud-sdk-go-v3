package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateWholeImageRequest struct {
	Body *CreateWholeImageRequestBody `json:"body,omitempty"`
}

func (o CreateWholeImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWholeImageRequest struct{}"
	}

	return strings.Join([]string{"CreateWholeImageRequest", string(data)}, " ")
}
