package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTagStatusRequest Request Object
type ShowTagStatusRequest struct {

	// 集群ID。
	ClusterId string `json:"cluster_id"`
}

func (o ShowTagStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTagStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowTagStatusRequest", string(data)}, " ")
}
