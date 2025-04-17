package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RefreshAssetRequest Request Object
type RefreshAssetRequest struct {
	Body *RefreshTaskReq `json:"body,omitempty"`
}

func (o RefreshAssetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RefreshAssetRequest struct{}"
	}

	return strings.Join([]string{"RefreshAssetRequest", string(data)}, " ")
}
