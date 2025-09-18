package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PublishPluginDto struct {

	// **参数解释**： 插件名。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	PluginName string `json:"plugin_name"`

	// **参数解释**： 版本。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Version string `json:"version"`

	// **参数解释**： 发布商ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	PublisherUniqueId string `json:"publisher_unique_id"`
}

func (o PublishPluginDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishPluginDto struct{}"
	}

	return strings.Join([]string{"PublishPluginDto", string(data)}, " ")
}
