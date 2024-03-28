package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListElasticResourcePoolsResponse Response Object
type ListElasticResourcePoolsResponse struct {

	// 是否成功
	IsSuccess *bool `json:"is_success,omitempty"`

	// 消息
	Message *string `json:"message,omitempty"`

	// 数量
	Count *int32 `json:"count,omitempty"`

	// 弹性资源池列表
	ElasticResourcePools *[]ElasticResourcePool `json:"elastic_resource_pools,omitempty"`
	HttpStatusCode       int                    `json:"-"`
}

func (o ListElasticResourcePoolsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListElasticResourcePoolsResponse struct{}"
	}

	return strings.Join([]string{"ListElasticResourcePoolsResponse", string(data)}, " ")
}
