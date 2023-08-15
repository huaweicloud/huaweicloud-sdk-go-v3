package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAssetRequest Request Object
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
