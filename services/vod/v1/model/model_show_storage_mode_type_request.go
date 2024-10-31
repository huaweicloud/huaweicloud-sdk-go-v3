package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStorageModeTypeRequest Request Object
type ShowStorageModeTypeRequest struct {
}

func (o ShowStorageModeTypeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStorageModeTypeRequest struct{}"
	}

	return strings.Join([]string{"ShowStorageModeTypeRequest", string(data)}, " ")
}
