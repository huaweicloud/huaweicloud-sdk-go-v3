package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdatePublicBandWidthRequest struct {
	// 指定查询集群ID。

	ClusterId string `json:"cluster_id"`

	Body *BindPublicReqEip `json:"body,omitempty"`
}

func (o UpdatePublicBandWidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePublicBandWidthRequest struct{}"
	}

	return strings.Join([]string{"UpdatePublicBandWidthRequest", string(data)}, " ")
}
