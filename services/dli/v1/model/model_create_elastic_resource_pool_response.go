package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateElasticResourcePoolResponse Response Object
type CreateElasticResourcePoolResponse struct {

	// 是否成功
	IsSuccess *bool `json:"is_success,omitempty"`

	// 响应消息
	Message *string `json:"message,omitempty"`

	// 创建成功的弹性资源池名称
	ElasticResourcePoolName *string `json:"elastic_resource_pool_name,omitempty"`
	HttpStatusCode          int     `json:"-"`
}

func (o CreateElasticResourcePoolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateElasticResourcePoolResponse struct{}"
	}

	return strings.Join([]string{"CreateElasticResourcePoolResponse", string(data)}, " ")
}
