package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateQueueToElasticResourcePoolRequest Request Object
type AssociateQueueToElasticResourcePoolRequest struct {

	// 弹性资源池名称
	ElasticResourcePoolName string `json:"elastic_resource_pool_name"`

	Body *AssociateQueueToElasticResourcePoolRequestBody `json:"body,omitempty"`
}

func (o AssociateQueueToElasticResourcePoolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateQueueToElasticResourcePoolRequest struct{}"
	}

	return strings.Join([]string{"AssociateQueueToElasticResourcePoolRequest", string(data)}, " ")
}
