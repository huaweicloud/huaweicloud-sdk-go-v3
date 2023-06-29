package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableComponentRequest Request Object
type EnableComponentRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 组件类型，取值为tsdb
	ComponentName string `json:"component_name"`

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
