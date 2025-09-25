package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeRequestBody GaussDB实例版本升级接口传参参数。
type UpgradeRequestBody struct {

	// **参数解释**: 实例升级类型，区分大小写。包括就地升级，灰度升级，热补丁升级三种，三种升级方式的异同，详见接口描述。 **约束限制**: 不涉及。 **取值范围**: - inplace：就地升级。 - grey：灰度升级。 - hotfix：热补丁升级。  **默认取值**:   不涉及。
	UpgradeType string `json:"upgrade_type"`

	// **参数解释**: 实例升级操作，区分大小写。 实例升级类型是就地升级时非必选。 灰度升级包括升级自动提交、升级待观察、提交升级和升级回退四种。 热补丁升级包括升级自动提交，升级回退两种。 **约束限制**: 不涉及。 **取值范围**: - upgradeAutoCommit：升级自动提交。 - upgrade：升级待观察。 - commit：提交升级。 - rollback：升级回退。  **默认取值**:   不涉及。
	UpgradeAction *string `json:"upgrade_action,omitempty"`

	// **参数解释**: 实例升级目标版本。 **约束限制**: - 热补丁升级场景下支持传入多个热补丁版本。  **取值范围**:   不涉及。 **默认取值**: 不涉及。
	TargetVersion *string `json:"target_version,omitempty"`

	// **参数解释**: 分布式实例灰度升级，滚动升级分片数。分布式实例灰度升级，升级待观察必填。该值不能大于实例未升级分片总数。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	UpgradeShardNum *int32 `json:"upgrade_shard_num,omitempty"`

	// **参数解释**: 灰度升级，升级待观察AZ值，升级待观察时必填。可以支持多AZ一起升级，AZ之间以英文逗号分隔。不能填入不属于该实例的AZ值。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	UpgradeAz *string `json:"upgrade_az,omitempty"`
}

func (o UpgradeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeRequestBody struct{}"
	}

	return strings.Join([]string{"UpgradeRequestBody", string(data)}, " ")
}
