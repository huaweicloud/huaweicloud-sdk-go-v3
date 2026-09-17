package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListPipelinesPageTagList struct {

	// **参数解释**： 标签ID。 **取值范围**： 不涉及。
	TagId *string `json:"tag_id,omitempty"`

	// **参数解释**： 标签名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 标签颜色。 **取值范围**： 不涉及。
	Color *string `json:"color,omitempty"`

	// **参数解释**： 项目ID。 **取值范围**： 不涉及。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**： 租户ID。 **取值范围**： 不涉及。
	DomainId *string `json:"domain_id,omitempty"`

	// **参数解释**： 创建人ID。 **取值范围**： 不涉及。
	CreatorId *string `json:"creator_id,omitempty"`

	// **参数解释**： 更新人ID。 **取值范围**： 不涉及。
	UpdaterId *string `json:"updater_id,omitempty"`

	// **参数解释**： 创建时间。 **取值范围**： 不涉及。
	CreateTime *int64 `json:"create_time,omitempty"`

	// **参数解释**： 更新时间。 **取值范围**： 不涉及。
	UpdateTime *int64 `json:"update_time,omitempty"`
}

func (o ListPipelinesPageTagList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipelinesPageTagList struct{}"
	}

	return strings.Join([]string{"ListPipelinesPageTagList", string(data)}, " ")
}
