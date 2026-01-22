package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceMessagesRequest Request Object
type ShowInstanceMessagesRequest struct {

	// **参数解释**： 实例ID。获取方法如下：调用“查询所有实例列表”接口，从响应体中获取实例ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**： Topic名称。 **约束限制**： Topic名称必须以字母开头且只支持大小写字母、中横线、下划线以及数字。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Topic string `json:"topic"`

	// **参数解释**： 是否按照时间排序。 **约束限制**： 不涉及。 **取值范围**： - true：按照时间排序。 - false：不按照时间排序。 **默认取值**： 不涉及。
	Asc *bool `json:"asc,omitempty"`

	// **参数解释**： 开始时间。  Unix毫秒时间戳。  **约束限制**： 按创建时间查询时，为必选参数。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**： 结束时间。  Unix毫秒时间戳。  **约束限制**： 按创建时间查询时，为必选参数。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**： 每一页显示的消息数量。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Limit *string `json:"limit,omitempty"`

	// **参数解释**： 页数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Offset *string `json:"offset,omitempty"`

	// **参数解释**： 是否下载消息到本地。 **约束限制**： 不涉及。 **取值范围**： - true：下载。 - false：不下载。 **默认取值**： 不涉及。
	Download *bool `json:"download,omitempty"`

	// **参数解释**： 消息偏移量。 **约束限制**： 按偏移量查询时，为必选参数。  若start_time、end_time参数不为空，该参数无效。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	MessageOffset *string `json:"message_offset,omitempty"`

	// **参数解释**： 分区。 **约束限制**： 按偏移量查询时，为必选参数。  若start_time、end_time参数不为空，该参数无效。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Partition *string `json:"partition,omitempty"`

	// **参数解释**： 设置查询消息的关键词。 **约束限制**： 不涉及。 **取值范围**： 0~50字符。 **默认取值**： 不涉及。
	Keyword *string `json:"keyword,omitempty"`

	// **参数解释**： 设置消息的KEY，查询结果为包含KEY的所有消息。 **约束限制**： 由于查询资源和性能限制，最大搜索10000条消息且所有消息总大小不超过200MB，最多返回包含KEY的前10条消息。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Key *string `json:"key,omitempty"`

	// **参数解释**： 设置消息正文中包含的关键词，查询结果为包含此关键词的消息。 **约束限制**： 多个关键字用%2C隔开，%2C是“,”的URL编码形式。 **取值范围**： include与exclude的关键词总数不得超过20个。 **默认取值**： 不涉及。
	Include *string `json:"include,omitempty"`

	// **参数解释**： 设置消息正文中需要排除的关键词，查询结果为不包含此关键词的消息。 **约束限制**： 多个关键字用%2C隔开，%2C是“,”的URL编码形式。 **取值范围**： include与exclude的关键词总数不得超过20个。 **默认取值**： 不涉及。
	Exclude *string `json:"exclude,omitempty"`
}

func (o ShowInstanceMessagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceMessagesRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceMessagesRequest", string(data)}, " ")
}
