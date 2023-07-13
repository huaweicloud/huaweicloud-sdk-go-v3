package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBusinessAssetsRequest Request Object
type ShowBusinessAssetsRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *BusinessAssetRequest `json:"body,omitempty"`
}

func (o ShowBusinessAssetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBusinessAssetsRequest struct{}"
	}

	return strings.Join([]string{"ShowBusinessAssetsRequest", string(data)}, " ")
}
