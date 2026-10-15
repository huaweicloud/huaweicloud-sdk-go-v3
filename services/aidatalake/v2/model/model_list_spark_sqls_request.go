package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSparkSqlsRequest Request Object
type ListSparkSqlsRequest struct {

	// **参数解释**：工作空间的ID。 **约束限制**：不涉及。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**：作业ID游标位置，用于分页查询。 **约束限制**：不涉及。 **取值范围**：采用UUID格式，长度为36个字符。 **默认取值**：首次查询可不传或传空字符串，后续查询传入上次返回的next_marker值。
	Marker *string `json:"marker,omitempty"`

	// **参数解释**：查询记录数，用于指定返回的作业数量。 **约束限制**：不涉及。 **取值范围**：1~100。 **默认取值**：10。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：是否反向分页查询。默认为false，表示正向分页查询。 当设置为true时，表示查询上一页数据，需要配合marker参数使用。
	PageReverse *bool `json:"page_reverse,omitempty"`

	// **参数解释**：端点名称，用于指定查询的SparkSql执行环境。 **约束限制**：不涉及。 **取值范围**：只能由英文小写字母、数字及中划线组成，以英文小写字母开头，以英文小写字母或数字结尾，且长度为1~63个字符。 **默认取值**：不涉及。
	EndpointName *string `json:"endpoint_name,omitempty"`

	// **参数解释**：SQL片段，用于模糊查询包含指定SQL片段的作业。 **约束限制**：不涉及。 **取值范围**：长度为1~128个字符。 **默认取值**：不涉及。
	Statement *string `json:"statement,omitempty"`

	// **参数解释**：用于查询创建时间在该时间点之后的作业。 **约束限制**：不涉及。 **取值范围**：采用unix时间戳格式，单位为毫秒，最小值为1764061598000。 **默认取值**：不涉及。
	CreateTimeAfter *int64 `json:"create_time_after,omitempty"`

	// **参数解释**：用于查询创建时间在该时间点之前的作业。 **约束限制**：不涉及。 **取值范围**：采用unix时间戳格式，单位为毫秒，最小值为1764061598000。 **默认取值**：不涉及。
	CreateTimeBefore *int64 `json:"create_time_before,omitempty"`

	// **参数解释**：SparkSql作业的状态列表，用于按状态过滤作业。 **约束限制**：数组元素数量不能超过10个。
	States *[]string `json:"states,omitempty"`

	// **参数解释**：作业类型列表，用于按作业类型过滤作业。 **约束限制**：数组元素数量不能超过4个。
	StatementTypes *[]string `json:"statement_types,omitempty"`

	// **参数解释**：SparkSql作业ID，用于查询指定作业。 **约束限制**：不涉及。 **取值范围**：只能由英文字母（不区分大小写）、数字、下划线及中划线组成，且长度为1~64个字符。 **默认取值**：不涉及。
	StatementId *string `json:"statement_id,omitempty"`

	// **参数解释**：作业创建者ID，用于按创建者过滤作业。 **约束限制**：不涉及。 **取值范围**：长度为1~256个字符。 **默认取值**：不涉及。
	CreateUserId *string `json:"create_user_id,omitempty"`

	// **参数解释**：作业创建者名称，用于按创建者精确过滤作业。 **约束限制**：不涉及。 **取值范围**：长度为1~256个字符。 **默认取值**：不涉及。
	CreateUserName *string `json:"create_user_name,omitempty"`

	// **参数解释**：作业标签，用于按标签过滤作业。格式为“key=value”，其中“=”需要URL编码为“%3D”。 **约束限制**：不涉及。 **取值范围**：长度为1~1024个字符。 **默认取值**：不涉及。
	Labels *string `json:"labels,omitempty"`
}

func (o ListSparkSqlsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSparkSqlsRequest struct{}"
	}

	return strings.Join([]string{"ListSparkSqlsRequest", string(data)}, " ")
}
