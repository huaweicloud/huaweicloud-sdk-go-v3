package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStackInstancesResponse Response Object
type ListStackInstancesResponse struct {

	// 资源栈实例列表
	StackInstances *[]StackInstance `json:"stack_instances,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListStackInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStackInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListStackInstancesResponse", string(data)}, " ")
}
