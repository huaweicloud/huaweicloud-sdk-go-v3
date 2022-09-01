package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateImageRequest struct {
	Body *CreateImageRequestBody `json:"body,omitempty" xml:"body"`
}

func (o CreateImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateImageRequest struct{}"
	}

	return strings.Join([]string{"CreateImageRequest", string(data)}, " ")
}
