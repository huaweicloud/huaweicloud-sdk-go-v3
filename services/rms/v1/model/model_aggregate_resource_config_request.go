package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询源帐号中的特定资源聚合的配置项请求体。
type AggregateResourceConfigRequest struct {

	// 资源聚合器ID。
	AggregatorId string `json:"aggregator_id"`

	ResourceIdentifier *ResourceIdentifier `json:"resource_identifier"`
}

func (o AggregateResourceConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AggregateResourceConfigRequest struct{}"
	}

	return strings.Join([]string{"AggregateResourceConfigRequest", string(data)}, " ")
}
