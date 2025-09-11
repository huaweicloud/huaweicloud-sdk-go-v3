package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceResourceInstancesResponse Response Object
type ListInstanceResourceInstancesResponse struct {

	// 资源列表
	Resources *[]Resource `json:"resources,omitempty"`

	// 资源数量
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstanceResourceInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceResourceInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceResourceInstancesResponse", string(data)}, " ")
}
