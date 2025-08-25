package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusterUriRequest Request Object
type ShowClusterUriRequest struct {

	// 密码集群ID
	ClusterId string `json:"cluster_id"`
}

func (o ShowClusterUriRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterUriRequest struct{}"
	}

	return strings.Join([]string{"ShowClusterUriRequest", string(data)}, " ")
}
