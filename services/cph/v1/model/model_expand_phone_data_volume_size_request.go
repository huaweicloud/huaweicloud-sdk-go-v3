package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExpandPhoneDataVolumeSizeRequest Request Object
type ExpandPhoneDataVolumeSizeRequest struct {
	Body *ExpandPhoneDataVolumeSizeRequestBody `json:"body,omitempty"`
}

func (o ExpandPhoneDataVolumeSizeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandPhoneDataVolumeSizeRequest struct{}"
	}

	return strings.Join([]string{"ExpandPhoneDataVolumeSizeRequest", string(data)}, " ")
}
