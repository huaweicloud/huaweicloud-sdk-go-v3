package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowExtensionDetailRequest Request Object
type ShowExtensionDetailRequest struct {
	Body *ExtensionSearchUserInputParamCustomizeForDetail `json:"body,omitempty"`
}

func (o ShowExtensionDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowExtensionDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowExtensionDetailRequest", string(data)}, " ")
}
