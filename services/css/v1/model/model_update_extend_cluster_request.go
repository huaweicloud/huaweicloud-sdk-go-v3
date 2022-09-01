package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateExtendClusterRequest struct {

	// 指定待扩容的集群ID。
	ClusterId string `json:"cluster_id" xml:"cluster_id"`

	Body *ExtendClusterReq `json:"body,omitempty" xml:"body"`
}

func (o UpdateExtendClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateExtendClusterRequest struct{}"
	}

	return strings.Join([]string{"UpdateExtendClusterRequest", string(data)}, " ")
}
