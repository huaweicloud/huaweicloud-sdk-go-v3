package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExpandClusterComponentRequest Request Object
type ExpandClusterComponentRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 语言类型
	XLanguage string `json:"X-Language"`

	Body *GrowNodeReq `json:"body,omitempty"`
}

func (o ExpandClusterComponentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandClusterComponentRequest struct{}"
	}

	return strings.Join([]string{"ExpandClusterComponentRequest", string(data)}, " ")
}
