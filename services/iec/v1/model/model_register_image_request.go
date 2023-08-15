package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterImageRequest Request Object
type RegisterImageRequest struct {
	Body *RegisterImageRequestBody `json:"body,omitempty"`
}

func (o RegisterImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterImageRequest struct{}"
	}

	return strings.Join([]string{"RegisterImageRequest", string(data)}, " ")
}
