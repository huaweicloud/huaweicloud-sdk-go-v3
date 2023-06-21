package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type EnableComponentRequest struct {

	// 租户ID
	ProjectId string `json:"projectId"`

	// 集群ID
	ClusterId string `json:"clusterId"`

	// 组件类型，取值为tsdb
	ComponentName string `json:"componentName"`

	// 语言类型
	XLanguage *string `json:"X-Language,omitempty"`

	Body *AddComponentReq `json:"body,omitempty"`
}

func (o EnableComponentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableComponentRequest struct{}"
	}

	return strings.Join([]string{"EnableComponentRequest", string(data)}, " ")
}
