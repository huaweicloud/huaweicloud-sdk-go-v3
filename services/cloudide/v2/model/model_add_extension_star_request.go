package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddExtensionStarRequest Request Object
type AddExtensionStarRequest struct {
	Body *ExtensionStar `json:"body,omitempty"`
}

func (o AddExtensionStarRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddExtensionStarRequest struct{}"
	}

	return strings.Join([]string{"AddExtensionStarRequest", string(data)}, " ")
}
