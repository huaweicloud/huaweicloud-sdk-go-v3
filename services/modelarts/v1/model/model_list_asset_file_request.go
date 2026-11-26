package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAssetFileRequest Request Object
type ListAssetFileRequest struct {

	// 工作空间ID。[获取方法请参见[查询工作空间列表](ListWorkspace.xml)。](tag:hc)未创建工作空间时默认值为“0”，存在创建并使用的工作空间，以实际取值为准。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释**：资产id。 **约束限制**：1-36字符，英文字母，数字，下划线，中划线。 **取值范围**：不涉及。 **默认取值**：不涉及。
	AssetId string `json:"asset_id"`

	// **参数解释**：列举桶内对象列表时，指定一个前缀，限定返回的对象名必须带有prefix前缀。 **约束限制**：1-1024位字符。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Prefix *string `json:"prefix,omitempty"`

	// **参数解释**：列举对象的最大数目，返回的对象列表将是按照字典顺序的最多前max_keys个对象。 **约束限制**：1-1000。 **取值范围**：不涉及。 **默认取值**：1000。
	MaxKeys *int32 `json:"max_keys,omitempty"`
}

func (o ListAssetFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssetFileRequest struct{}"
	}

	return strings.Join([]string{"ListAssetFileRequest", string(data)}, " ")
}
