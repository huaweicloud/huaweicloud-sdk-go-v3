package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateClusterNameRequest Request Object
type UpdateClusterNameRequest struct {

	// 集群ID。获取方法，请参见[获取集群ID](https://support.huaweicloud.com/api-mrs/mrs_02_9001.html)。
	ClusterId string `json:"cluster_id"`

	Body *UpdateClusterReq `json:"body,omitempty"`
}

func (o UpdateClusterNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClusterNameRequest struct{}"
	}

	return strings.Join([]string{"UpdateClusterNameRequest", string(data)}, " ")
}
