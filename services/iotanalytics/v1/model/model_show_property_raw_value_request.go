package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPropertyRawValueRequest Request Object
type ShowPropertyRawValueRequest struct {

	// 资产ID
	AssetId string `json:"asset_id"`

	Body *RawRequest `json:"body,omitempty"`
}

func (o ShowPropertyRawValueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPropertyRawValueRequest struct{}"
	}

	return strings.Join([]string{"ShowPropertyRawValueRequest", string(data)}, " ")
}
