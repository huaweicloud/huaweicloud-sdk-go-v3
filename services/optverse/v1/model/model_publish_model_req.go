package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishModelReq **参数解释**： 发布模型请求体。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
type PublishModelReq struct {

	// **参数解释**： 训练任务ID。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	TrainJobId *string `json:"train_job_id,omitempty"`

	// **参数解释**： 模型名称。 **约束限制**： 不涉及 **取值范围**： 长度为[1-128]个字符。 **默认取值**： 不涉及
	Name string `json:"name"`

	// **参数解释**： 模型名称。 **约束限制**： 不涉及 **取值范围**： 长度为[1-2048]个字符。 **默认取值**： 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释**： 模型在OBS存储位置。 **约束限制**： 不涉及 **取值范围**： 长度为[1-512]个字符。 **默认取值**： 不涉及
	AssetLocation *string `json:"asset_location,omitempty"`

	// **参数解释**： 对话id。 **约束限制**： 不涉及 **取值范围**： 长度为[1-64]个字符。 **默认取值**： 不涉及
	ChatId *string `json:"chat_id,omitempty"`

	// **参数解释**： 模型版本。 **约束限制**： 不涉及 **取值范围**： 长度为[1-64]个字符。 **默认取值**： 不涉及
	AssetVersion *string `json:"asset_version,omitempty"`

	// **参数解释**： 模型对外显示版本。 **约束限制**： 不涉及 **取值范围**： 长度为[1-64]个字符。 **默认取值**： 不涉及
	ExternalVersion *string `json:"external_version,omitempty"`

	// **参数解释**： 模型类型。 **约束限制**： 不涉及 **取值范围**： 长度为[1-64]个字符。 **默认取值**： 不涉及
	AssetType *string `json:"asset_type,omitempty"`

	// **参数解释**： 模型子类型。 **约束限制**： 不涉及 **取值范围**： 长度为[1-64]个字符。 **默认取值**： 不涉及
	SubAssetType *string `json:"sub_asset_type,omitempty"`
}

func (o PublishModelReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishModelReq struct{}"
	}

	return strings.Join([]string{"PublishModelReq", string(data)}, " ")
}
