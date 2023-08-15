package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateQueueToElasticResourcePoolRequestBody 队列关联弹性资源池
type AssociateQueueToElasticResourcePoolRequestBody struct {

	// 队列名称
	QueueName string `json:"queue_name"`
}

func (o AssociateQueueToElasticResourcePoolRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateQueueToElasticResourcePoolRequestBody struct{}"
	}

	return strings.Join([]string{"AssociateQueueToElasticResourcePoolRequestBody", string(data)}, " ")
}
