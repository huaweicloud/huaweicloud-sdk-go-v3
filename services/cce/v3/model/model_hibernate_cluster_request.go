package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type HibernateClusterRequest struct {
	// 集群 ID，获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。

	ClusterId string `json:"cluster_id"`
}

func (o HibernateClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HibernateClusterRequest struct{}"
	}

	return strings.Join([]string{"HibernateClusterRequest", string(data)}, " ")
}
