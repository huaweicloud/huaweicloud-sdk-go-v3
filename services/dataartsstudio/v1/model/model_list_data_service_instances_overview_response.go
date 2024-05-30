package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDataServiceInstancesOverviewResponse Response Object
type ListDataServiceInstancesOverviewResponse struct {

	// 集群数量。
	Total *int32 `json:"total,omitempty"`

	// 是否支持缩容。
	ScaleDown *bool `json:"scale_down,omitempty"`

	// 是否支持扩容。
	ScaleOut *bool `json:"scale_out,omitempty"`

	// 集群概览信息。
	Instances      *[]InstanceOverviewDto `json:"instances,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListDataServiceInstancesOverviewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataServiceInstancesOverviewResponse struct{}"
	}

	return strings.Join([]string{"ListDataServiceInstancesOverviewResponse", string(data)}, " ")
}
