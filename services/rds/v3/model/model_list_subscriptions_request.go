package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubscriptionsRequest Request Object
type ListSubscriptionsRequest struct {

	// 语言。默认en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 发布id。不为空则查询该发布下的订阅；为空（null）则查询实例本地订阅。
	PublicationId *string `json:"publication_id,omitempty"`

	// 订阅服务器来源  true：订阅服务器为云上实例  false：订阅服务器非云上实例  不传该参数（null）：全部订阅
	IsCloud *bool `json:"is_cloud,omitempty"`

	// 发布名（模糊匹配）
	PublicationName *string `json:"publication_name,omitempty"`

	// 订阅数据库名（模糊匹配）
	SubscriptionDbName *string `json:"subscription_db_name,omitempty"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。默认为10，不能为负数，最小值为1，最大值为100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListSubscriptionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubscriptionsRequest struct{}"
	}

	return strings.Join([]string{"ListSubscriptionsRequest", string(data)}, " ")
}
