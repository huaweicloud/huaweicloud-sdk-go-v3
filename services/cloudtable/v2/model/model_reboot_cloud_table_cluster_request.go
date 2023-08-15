package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RebootCloudTableClusterRequest Request Object
type RebootCloudTableClusterRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 语言类型
	XLanguage *string `json:"X-Language,omitempty"`

	Body *HbaseClusterActionReq `json:"body,omitempty"`
}

func (o RebootCloudTableClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RebootCloudTableClusterRequest struct{}"
	}

	return strings.Join([]string{"RebootCloudTableClusterRequest", string(data)}, " ")
}
