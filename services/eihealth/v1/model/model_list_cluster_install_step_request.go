package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterInstallStepRequest Request Object
type ListClusterInstallStepRequest struct {

	// 计算集群ID。
	ClusterId string `json:"cluster_id"`
}

func (o ListClusterInstallStepRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterInstallStepRequest struct{}"
	}

	return strings.Join([]string{"ListClusterInstallStepRequest", string(data)}, " ")
}
