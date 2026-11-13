package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllTypeInstancesResponse Response Object
type ListAllTypeInstancesResponse struct {

	// 实例列表
	Instances *[]DasCommonInstanceDto `json:"instances,omitempty"`

	// 总记录数
	TotalCount     *int64 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAllTypeInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllTypeInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListAllTypeInstancesResponse", string(data)}, " ")
}
