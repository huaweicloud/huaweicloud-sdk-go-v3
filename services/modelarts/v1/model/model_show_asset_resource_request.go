package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAssetResourceRequest Request Object
type ShowAssetResourceRequest struct {

	// **参数解释**：资产id。 **约束限制**：1-36字符，英文字母，数字，下划线，中划线。 **取值范围**：不涉及。 **默认取值**：不涉及。
	AssetId string `json:"asset_id"`

	// 工作空间ID。[获取方法请参见[查询工作空间列表](ListWorkspace.xml)。](tag:hc)未创建工作空间时默认值为“0”，存在创建并使用的工作空间，以实际取值为准。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释**：静态资源名称。 **约束限制**：当前只支持readme文件查询（docs/RELEASE.md）。 **取值范围**：不涉及。 **默认取值**：不涉及。
	ResourceName string `json:"resource_name"`

	// **参数解释**：查询资源时是否拼接区域信息。如果为false，则直接查询docs/RELEASE.md文件；如果为true，则会优先匹配带有区域信息的文件（请求头传递x-language，如英文查询docs/RELEASE_en-US.md）。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	CheckLocale *bool `json:"check_locale,omitempty"`
}

func (o ShowAssetResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAssetResourceRequest struct{}"
	}

	return strings.Join([]string{"ShowAssetResourceRequest", string(data)}, " ")
}
