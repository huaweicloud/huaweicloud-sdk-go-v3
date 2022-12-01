package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListDataSourceRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`
}

func (o ListDataSourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataSourceRequest struct{}"
	}

	return strings.Join([]string{"ListDataSourceRequest", string(data)}, " ")
}
