package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPublicationsRequest Request Object
type ListPublicationsRequest struct {

	// 语言。默认en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 发布名（模糊匹配）
	PublicationName *string `json:"publication_name,omitempty"`

	// 发布数据库名（模糊匹配）
	PublicationDbName *string `json:"publication_db_name,omitempty"`

	// 订阅实例id
	SubscriberInstanceId *string `json:"subscriber_instance_id,omitempty"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。默认为10，不能为负数，最小值为1，最大值为100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListPublicationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPublicationsRequest struct{}"
	}

	return strings.Join([]string{"ListPublicationsRequest", string(data)}, " ")
}
