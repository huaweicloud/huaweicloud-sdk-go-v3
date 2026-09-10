package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowModelAssetDetailResponse Response Object
type ShowModelAssetDetailResponse struct {

	// **参数解释**： 资产ID。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	AssetId *string `json:"asset_id,omitempty"`

	// **参数解释**： 根资产ID。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	RootAssetId *string `json:"root_asset_id,omitempty"`

	// **参数解释**： 资产名称。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	AssetName *string `json:"asset_name,omitempty"`

	// **参数解释**： 资产内部版本。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	AssetVersion *string `json:"asset_version,omitempty"`

	// **参数解释**： 资产对外版本。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	ExternalVersion *string `json:"external_version,omitempty"`

	// **参数解释**： 对外是否可见。 **约束限制**： 不涉及 **取值范围**： * 1：可见 * 0：不可见 **默认取值**： 不涉及
	IsAvailable *string `json:"is_available,omitempty"`

	// **参数解释**： OBS存储位置。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	AssetLocation *string `json:"asset_location,omitempty"`

	// **参数解释**： 资产描述。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	AssetDesc *string `json:"asset_desc,omitempty"`

	// **参数解释**： 资产类型。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	AssetType *string `json:"asset_type,omitempty"`

	// **参数解释**： 资产子类型。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	SubAssetType *string `json:"sub_asset_type,omitempty"`

	// **参数解释**： 资产编码。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	AssetCode *string `json:"asset_code,omitempty"`

	// **参数解释**： 资产来源。 **约束限制**： 不涉及 **取值范围**： * Preset：预置 * AIGallery：订阅 * Import：导入 * Publish：发布 **默认取值**： 不涉及
	AssetSource *string `json:"asset_source,omitempty"`

	// **参数解释**： 资产应用场景，多个场景以英文逗号分隔。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	AssetActions *string `json:"asset_actions,omitempty"`

	// **参数解释**： 更新时间（毫秒时间戳）。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	UpdateTime *int64 `json:"update_time,omitempty"`

	// **参数解释**： 创建时间（毫秒时间戳）。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	CreateTime *int64 `json:"create_time,omitempty"`

	// **参数解释**： 模型创建者名称。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Creator *string `json:"creator,omitempty"`

	// **参数解释**： 模型创建者ID。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	UserId *string `json:"user_id,omitempty"`

	// **参数解释**： 生成该模型时用到的数据集。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	TrainObsUrl    *string `json:"train_obs_url,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowModelAssetDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowModelAssetDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowModelAssetDetailResponse", string(data)}, " ")
}
