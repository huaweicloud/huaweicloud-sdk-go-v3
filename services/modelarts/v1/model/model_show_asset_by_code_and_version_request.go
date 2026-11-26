package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAssetByCodeAndVersionRequest Request Object
type ShowAssetByCodeAndVersionRequest struct {

	// **参数解释**：资产code。 **约束限制**：1-60字符，英文字母，数字，中划线，下划线，点。 **取值范围**：不涉及。 **默认取值**：不涉及。
	AssetCode string `json:"asset_code"`

	// **参数解释**：资产version。 **约束限制**：数字或为格式x.y.z， x,y,z均为数字。 **取值范围**：不涉及。 **默认取值**：不涉及。
	AssetVersion string `json:"asset_version"`

	// 工作空间ID。[获取方法请参见[查询工作空间列表](ListWorkspace.xml)。](tag:hc)未创建工作空间时默认值为“0”，存在创建并使用的工作空间，以实际取值为准。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释**:资产类型。 **约束限制**:不涉及。 **取值范围**:不涉及。 **默认取值**:MODEL。
	AssetType *string `json:"asset_type,omitempty"`
}

func (o ShowAssetByCodeAndVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAssetByCodeAndVersionRequest struct{}"
	}

	return strings.Join([]string{"ShowAssetByCodeAndVersionRequest", string(data)}, " ")
}
