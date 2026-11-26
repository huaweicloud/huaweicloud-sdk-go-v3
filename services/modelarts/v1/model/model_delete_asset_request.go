package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAssetRequest Request Object
type DeleteAssetRequest struct {

	// **参数解释**：资产id。 **约束限制**：1-36字符，英文字母，数字，下划线，中划线。 **取值范围**：不涉及。 **默认取值**：不涉及。
	AssetId string `json:"asset_id"`

	// 工作空间ID。[获取方法请参见[查询工作空间列表](ListWorkspace.xml)。](tag:hc)未创建工作空间时默认值为“0”，存在创建并使用的工作空间，以实际取值为准。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释**：数据集资产是否删除OBS桶中的数据，boolean类型，取值true：删除OBS桶中的数据，取值false：仅删除记录不删除OBS桶中的数据。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	ForceDelete *bool `json:"force_delete,omitempty"`
}

func (o DeleteAssetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAssetRequest struct{}"
	}

	return strings.Join([]string{"DeleteAssetRequest", string(data)}, " ")
}
