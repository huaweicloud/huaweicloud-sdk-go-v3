package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAssetTagsInfoRequest Request Object
type ShowAssetTagsInfoRequest struct {

	// **参数解释**：标签分类。 **约束限制**：1-128位字符，仅支持中文、英文字母、数字、下划线和中划线。 **取值范围**：不涉及。 **默认取值**：不涉及。
	TagCategory string `json:"tag_category"`

	// **参数解释**：请求头（Header）参数，主要用于控制API返回内容的语言。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	XLanguage *string `json:"X-Language,omitempty"`

	// 工作空间ID。[获取方法请参见[查询工作空间列表](ListWorkspace.xml)。](tag:hc)未创建工作空间时默认值为“0”，存在创建并使用的工作空间，以实际取值为准。
	WorkspaceId *string `json:"workspace_id,omitempty"`
}

func (o ShowAssetTagsInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAssetTagsInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowAssetTagsInfoRequest", string(data)}, " ")
}
