package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateImageRequest struct {
	Body *CreateImageRequestBody `json:"body,omitempty"`
}

func (o CreateImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateImageRequest struct{}"
	}

	return strings.Join([]string{"CreateImageRequest", string(data)}, " ")
}
