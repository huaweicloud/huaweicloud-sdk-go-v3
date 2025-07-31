package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunHostAssetManualCollectResponse Response Object
type RunHostAssetManualCollectResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RunHostAssetManualCollectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunHostAssetManualCollectResponse struct{}"
	}

	return strings.Join([]string{"RunHostAssetManualCollectResponse", string(data)}, " ")
}
