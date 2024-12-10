package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GenerateRandomPasswordRequest Request Object
type GenerateRandomPasswordRequest struct {
	Body *CreatePasswordRequestBody `json:"body,omitempty"`
}

func (o GenerateRandomPasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GenerateRandomPasswordRequest struct{}"
	}

	return strings.Join([]string{"GenerateRandomPasswordRequest", string(data)}, " ")
}
