package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModelConfigDto 模型配置数据对象
type ModelConfigDto struct {

	// **参数解释**： 工作项属性。 **取值范围**： 不涉及。
	Categories *[]BaseCategory `json:"categories,omitempty"`

	// **参数解释**： 工作项层级关系。 **取值范围**： 不涉及。
	CategoryLayerConfig *[]CategoryLayerDto `json:"category_layer_config,omitempty"`

	// **参数解释**： 工作项功能页面跳转链接模板。 **取值范围**： 不涉及。
	FeaturePageLinkTemplate *string `json:"feature_page_link_template,omitempty"`
}

func (o ModelConfigDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelConfigDto struct{}"
	}

	return strings.Join([]string{"ModelConfigDto", string(data)}, " ")
}
