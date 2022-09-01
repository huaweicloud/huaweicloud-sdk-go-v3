package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 维度对象结构体
type DimChild struct {

	// 维度名称，当前支持维度有dcs_instance_id、dcs_cluster_redis_node、 dcs_cluster_proxy_node和dcs_memcached_instance_id。
	DimName *string `json:"dim_name,omitempty" xml:"dim_name"`

	// 维度的路由，结构为主维度名称,当前维度名称，比如： dim_name字段为dcs_cluster_redis_node时，这个字段的值为dcs_instance_id,dcs_cluster_redis_node。
	DimRoute *string `json:"dim_route,omitempty" xml:"dim_route"`
}

func (o DimChild) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DimChild struct{}"
	}

	return strings.Join([]string{"DimChild", string(data)}, " ")
}
