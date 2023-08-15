package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteElasticResourcePoolRequest Request Object
type DeleteElasticResourcePoolRequest struct {

	// 弹性资源池名称
	ElasticResourcePoolName string `json:"elastic_resource_pool_name"`
}

func (o DeleteElasticResourcePoolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteElasticResourcePoolRequest struct{}"
	}

	return strings.Join([]string{"DeleteElasticResourcePoolRequest", string(data)}, " ")
}
