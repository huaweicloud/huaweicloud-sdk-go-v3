package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDataImageInSafeModeRequest Request Object
type CreateDataImageInSafeModeRequest struct {
	Body *CreateDataImageRequestBody `json:"body,omitempty"`
}

func (o CreateDataImageInSafeModeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataImageInSafeModeRequest struct{}"
	}

	return strings.Join([]string{"CreateDataImageInSafeModeRequest", string(data)}, " ")
}
