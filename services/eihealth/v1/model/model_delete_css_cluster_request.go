package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCssClusterRequest Request Object
type DeleteCssClusterRequest struct {

	// css集群id
	CssClusterId string `json:"css_cluster_id"`
}

func (o DeleteCssClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCssClusterRequest struct{}"
	}

	return strings.Join([]string{"DeleteCssClusterRequest", string(data)}, " ")
}
