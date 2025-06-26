package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubResourceInstancesResponse Response Object
type ListSubResourceInstancesResponse struct {

	// 资源实例列表
	Resources *[]Resource `json:"resources,omitempty"`

	// 资源实例数量
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSubResourceInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubResourceInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListSubResourceInstancesResponse", string(data)}, " ")
}
