package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionModelQueryViewDto struct {
	Branch *VersionObjectBranchQueryViewDto `json:"branch,omitempty"`

	// **参数解释：**  检出时间。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	CheckOutTime *string `json:"checkOutTime,omitempty"`

	// **参数解释：**  检出人。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	CheckOutUserName *string `json:"checkOutUserName,omitempty"`

	// **参数解释：**  类名。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	ClassName *string `json:"className,omitempty"`

	// **参数解释：**  创建时间。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	CreateTime *string `json:"createTime,omitempty"`

	// **参数解释：**  创建者。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Creator *string `json:"creator,omitempty"`

	// **参数解释：**  描述信息。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释：**  唯一标识。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：**  迭代版本。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Iteration *int32 `json:"iteration,omitempty"`

	// **参数解释：**  最后更新时间。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// **参数解释：**  是否为最新版本。  **取值范围：**  - true：是最新版本。 - false：不是最新版本。  **默认取值：**  false。
	Latest *bool `json:"latest,omitempty"`

	// **参数解释：**  是否为最新迭代版本。  **取值范围：**  - true：是最新迭代版本。 - false：不是最新迭代版本。  **默认取值：**  false。
	LatestIteration *bool `json:"latestIteration,omitempty"`

	// **参数解释：**  是否为最新修订版本。  **取值范围：**  - true：是最新修订版本。 - false：不是最新修订版本。  **默认取值：**  false。
	LatestVersion *bool `json:"latestVersion,omitempty"`

	Master *MasterObjectQueryViewDto `json:"master,omitempty"`

	// **参数解释：**  修改人。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`

	// **参数解释：**  中文名称。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：**  前序版本实例ID。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	PreVersionId *string `json:"preVersionId,omitempty"`

	// **参数解释：**  扩展类型。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	RdmExtensionType *string `json:"rdmExtensionType,omitempty"`

	Tenant *TenantQueryViewDto `json:"tenant,omitempty"`

	// **参数解释：**  版本号。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Version *string `json:"version,omitempty"`

	// **参数解释：**  业务版本内码。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	VersionCode *int32 `json:"versionCode,omitempty"`

	// **参数解释：**  是否已检出。  **取值范围：**  - true：已检出。 - false：未检出。  **默认取值：**  不涉及。
	WorkingCopy *bool `json:"workingCopy,omitempty"`

	WorkingState *WorkingState `json:"workingState,omitempty"`
}

func (o VersionModelQueryViewDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionModelQueryViewDto struct{}"
	}

	return strings.Join([]string{"VersionModelQueryViewDto", string(data)}, " ")
}
