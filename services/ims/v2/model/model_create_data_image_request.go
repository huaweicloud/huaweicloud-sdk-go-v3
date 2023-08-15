package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDataImageRequest Request Object
type CreateDataImageRequest struct {
	Body *CreateDataImageRequestBody `json:"body,omitempty"`
}

func (o CreateDataImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataImageRequest struct{}"
	}

	return strings.Join([]string{"CreateDataImageRequest", string(data)}, " ")
}
