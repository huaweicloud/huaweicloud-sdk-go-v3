package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ExpandClusterComponentRequest struct {

	// 租户ID
	ProjectId string `json:"projectId"`

	// 集群ID
	ClusterId string `json:"clusterId"`

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
