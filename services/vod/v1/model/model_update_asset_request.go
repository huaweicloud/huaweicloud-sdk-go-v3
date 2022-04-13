package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateAssetRequest struct {
	Body *UploadAssetReq `json:"body,omitempty"`
}

func (o UpdateAssetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAssetRequest struct{}"
	}

	return strings.Join([]string{"UpdateAssetRequest", string(data)}, " ")
}
