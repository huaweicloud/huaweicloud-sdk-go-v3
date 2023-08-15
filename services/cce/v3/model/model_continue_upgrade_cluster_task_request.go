package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ContinueUpgradeClusterTaskRequest Request Object
type ContinueUpgradeClusterTaskRequest struct {

	// 集群ID，获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。
	ClusterId string `json:"cluster_id"`
}

func (o ContinueUpgradeClusterTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContinueUpgradeClusterTaskRequest struct{}"
	}

	return strings.Join([]string{"ContinueUpgradeClusterTaskRequest", string(data)}, " ")
}
