package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateStorageModeTypeRequest Request Object
type UpdateStorageModeTypeRequest struct {
	Body *UpdateStorageModeTypeReq `json:"body,omitempty"`
}

func (o UpdateStorageModeTypeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStorageModeTypeRequest struct{}"
	}

	return strings.Join([]string{"UpdateStorageModeTypeRequest", string(data)}, " ")
}
