package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteTopicAttributeByNameRequest struct {
	// Topic的唯一的资源标识，可通过[查询主题列表](https://support.huaweicloud.com/api-smn/smn_api_51004.html)获取该标识。

	TopicUrn string `json:"topic_urn"`
	// 主题策略名称。 只支持特定的策略名称，请参见[Topic属性表](https://support.huaweicloud.com/intl/zh-cn/api-smn/smn_api_a1000.html)。

	Name string `json:"name"`
}

func (o DeleteTopicAttributeByNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTopicAttributeByNameRequest struct{}"
	}

	return strings.Join([]string{"DeleteTopicAttributeByNameRequest", string(data)}, " ")
}
