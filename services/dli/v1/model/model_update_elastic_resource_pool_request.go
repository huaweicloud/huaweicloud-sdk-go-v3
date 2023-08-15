package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateElasticResourcePoolRequest Request Object
type UpdateElasticResourcePoolRequest struct {

	// 弹性资源池名称
	ElasticResourcePoolName string `json:"elastic_resource_pool_name"`

	Body *UpdateElasticResourcePoolRequestBody `json:"body,omitempty"`
}

func (o UpdateElasticResourcePoolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateElasticResourcePoolRequest struct{}"
	}

	return strings.Join([]string{"UpdateElasticResourcePoolRequest", string(data)}, " ")
}
