package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstancesApiResponse Response Object
type ListInstancesApiResponse struct {

	// 实例列表
	InstanceInfos *[]DasInstanceInfoDto `json:"instance_infos,omitempty"`

	// 总数
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstancesApiResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesApiResponse struct{}"
	}

	return strings.Join([]string{"ListInstancesApiResponse", string(data)}, " ")
}
