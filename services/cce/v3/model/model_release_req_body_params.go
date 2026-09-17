package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReleaseReqBodyParams **参数解释：** 模板实例参数 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type ReleaseReqBodyParams struct {

	// **参数解释：** 开启后，仅验证模板参数，不进行安装 **约束限制：** 不涉及 **取值范围：** - true：仅验证 - false：正常安装  **默认取值：** false
	DryRun *bool `json:"dry_run,omitempty"`

	// **参数解释：** 实例名称模板 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	NameTemplate *string `json:"name_template,omitempty"`

	// **参数解释：** 安装时是否禁用hooks **约束限制：** 不涉及 **取值范围：** - true：禁用hooks - false：不禁用hooks  **默认取值：** false
	NoHooks *bool `json:"no_hooks,omitempty"`

	// **参数解释：** 模板实例更新时是否保留values，该字段仅在更新指定模板实例时生效 **约束限制：** 不涉及 **取值范围：** - true：保留values - false：不保留values  **默认取值：** false
	Replace *bool `json:"replace,omitempty"`

	// **参数解释：** 模板实例更新时是否重建实例，该字段仅在更新指定模板实例时生效 **约束限制：** 不涉及 **取值范围：** - true：重建实例 - false：不重建实例  **默认取值：** false
	Recreate *bool `json:"recreate,omitempty"`

	// **参数解释：** 模板实例更新时是否重置values，该字段仅在更新指定模板实例时生效 **约束限制：** 不涉及 **取值范围：** - true：重置values - false：不重置values  **默认取值：** false
	ResetValues *bool `json:"reset_values,omitempty"`

	// **参数解释：** 回滚实例的版本 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ReleaseVersion *int32 `json:"release_version,omitempty"`

	// **参数解释：** 更新或者删除时启用hooks **约束限制：** 不涉及 **取值范围：** - true：启用hooks - false：不启用hooks  **默认取值：** false
	IncludeHooks *bool `json:"include_hooks,omitempty"`
}

func (o ReleaseReqBodyParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReleaseReqBodyParams struct{}"
	}

	return strings.Join([]string{"ReleaseReqBodyParams", string(data)}, " ")
}
