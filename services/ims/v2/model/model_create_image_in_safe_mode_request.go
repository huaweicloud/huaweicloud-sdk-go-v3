package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateImageInSafeModeRequest Request Object
type CreateImageInSafeModeRequest struct {
	Body *CreateImageRequestBody `json:"body,omitempty"`
}

func (o CreateImageInSafeModeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateImageInSafeModeRequest struct{}"
	}

	return strings.Join([]string{"CreateImageInSafeModeRequest", string(data)}, " ")
}
