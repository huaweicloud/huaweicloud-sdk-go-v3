package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListReplicationErrorsRequest Request Object
type ListReplicationErrorsRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 订阅id。 若该参数为空，则查询当前实例作为分发服务器时，分发下的所有报错信息。 若该参数不为空，则查询该订阅的所有报错信息。
	SubscriptionId *string `json:"subscription_id,omitempty"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。默认为10，不能为负数，最小值为1，最大值为100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListReplicationErrorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListReplicationErrorsRequest struct{}"
	}

	return strings.Join([]string{"ListReplicationErrorsRequest", string(data)}, " ")
}
