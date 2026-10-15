package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRayJobsRequest Request Object
type ListRayJobsRequest struct {

	// **参数解释**：工作空间的ID。 **约束限制**：不涉及。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**：上一页中最后一条记录id，查询第一页时传空值。 **约束限制**：不涉及。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	Marker *string `json:"marker,omitempty"`

	// **参数解释**：指定每一页返回的最大条目数。 **约束限制**：不涉及。 **取值范围**：1~100。 **默认取值**：10。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：通过名字搜索作业。 **约束限制**：不涉及。 **取值范围**：长度为1~64的中文、字母、数字、下划线、中划线的组合。 **默认取值**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：通过作业id检索。 **约束限制**：不涉及。 **取值范围**：长度为1~36的英文字符、数字和中划线的组合。 **默认取值**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：通过端点名称检索的参数。 **约束限制**：不涉及。 **取值范围**：长度为1~64的中文、英文字母、数字、下划线、中划线、点号、空格的组合。 **默认取值**：不涉及。
	EndpointName *string `json:"endpoint_name,omitempty"`

	// **参数解释**：该参数用于按联邦名称搜索端点。 **约束限制**：不涉及。 **取值范围**：长度为1~64的中文、英文字母、数字、下划线、中划线、点号、空格的组合。 **默认取值**：不涉及。
	FederationName *string `json:"federation_name,omitempty"`

	// **参数解释**：用于查询创建时间在该时间点之前的作业。unix时间戳，单位：毫秒。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	CreateTimeBefore *int64 `json:"create_time_before,omitempty"`

	// **参数解释**：用于查询创建时间在该时间点之后的作业。unix时间戳，单位：毫秒。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	CreateTimeAfter *int64 `json:"create_time_after,omitempty"`

	// **参数解释**：状态过滤，支持多种状态查询，默认查询所有。 **约束限制**：不涉及。 **取值范围**：可选值有：   - QUEUED：排队中。   - PENDING：待处理。   - RUNNING：运行中。   - CANCELING：取消中。   - CANCELED：已取消。   - FAILED：失败。   - QUEUED_TIMEOUT：排队超时。   - RUNNING_TIMEOUT：运行超时。   - SUCCEEDED：成功。 **默认取值**：不涉及。
	Status *[]string `json:"status,omitempty"`
}

func (o ListRayJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRayJobsRequest struct{}"
	}

	return strings.Join([]string{"ListRayJobsRequest", string(data)}, " ")
}
