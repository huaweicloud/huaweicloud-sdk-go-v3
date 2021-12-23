package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAuditlogLinksRequest struct {
	// 实例ID，可以调用“查询实例列表”接口获取。如果未申请实例，可以调用“创建实例”接口创建。

	InstanceId string `json:"instance_id"`

	Body *ProduceAuditlogLinksRequestBody `json:"body,omitempty"`
}

func (o ListAuditlogLinksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditlogLinksRequest struct{}"
	}

	return strings.Join([]string{"ListAuditlogLinksRequest", string(data)}, " ")
}
