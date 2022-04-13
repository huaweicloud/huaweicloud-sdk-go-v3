package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateOndemandClusterToPeriodRequest struct {
	// 指定待更改的集群ID。

	ClusterId string `json:"cluster_id"`

	Body *PeriodReq `json:"body,omitempty"`
}

func (o UpdateOndemandClusterToPeriodRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOndemandClusterToPeriodRequest struct{}"
	}

	return strings.Join([]string{"UpdateOndemandClusterToPeriodRequest", string(data)}, " ")
}
