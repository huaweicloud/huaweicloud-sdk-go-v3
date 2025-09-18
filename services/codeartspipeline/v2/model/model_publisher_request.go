package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PublisherRequest struct {

	// **参数解释**： 发布商名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 用户ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	UserId *string `json:"user_id,omitempty"`

	// **参数解释**： 发布商描述。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 发布商图标URL。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	LogoUrl *string `json:"logo_url,omitempty"`

	// **参数解释**： 发布商网页地址。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Website *string `json:"website,omitempty"`

	// **参数解释**： 发布商帮助地址。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	SupportUrl string `json:"support_url"`

	// **参数解释**： 发布商源码地址。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	SourceUrl *string `json:"source_url,omitempty"`

	// **参数解释**： 发布商英文名。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	EnName string `json:"en_name"`

	// **参数解释**： 发布商ID。可通过[查询发布商详情](ShowPublisher.xml)查询。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	PublisherUniqueId *string `json:"publisher_unique_id,omitempty"`
}

func (o PublisherRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublisherRequest struct{}"
	}

	return strings.Join([]string{"PublisherRequest", string(data)}, " ")
}
