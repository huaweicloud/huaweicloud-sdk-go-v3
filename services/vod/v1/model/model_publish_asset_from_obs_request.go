package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type PublishAssetFromObsRequest struct {
	Body *PublishAssetFromObsReq `json:"body,omitempty"`
}

func (o PublishAssetFromObsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishAssetFromObsRequest struct{}"
	}

	return strings.Join([]string{"PublishAssetFromObsRequest", string(data)}, " ")
}
