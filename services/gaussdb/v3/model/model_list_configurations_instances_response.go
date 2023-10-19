package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConfigurationsInstancesResponse Response Object
type ListConfigurationsInstancesResponse struct {

	// 可以应用的实例列表。
	Instances *[]ApplicableInstances `json:"instances,omitempty"`

	// 可应用的实例列表数量。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListConfigurationsInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigurationsInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListConfigurationsInstancesResponse", string(data)}, " ")
}
