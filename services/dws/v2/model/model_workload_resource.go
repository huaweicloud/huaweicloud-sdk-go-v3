package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 工作负载队列资源项
type WorkloadResource struct {

	// 资源名称。
	ResourceName *string `json:"resource_name,omitempty"`

	// 资源属性值。
	ResourceValue *int32 `json:"resource_value,omitempty"`
}

func (o WorkloadResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadResource struct{}"
	}

	return strings.Join([]string{"WorkloadResource", string(data)}, " ")
}
