package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionModelViewDto struct {
	Branch *VersionModelBranchViewDto `json:"branch,omitempty"`

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

	// **参数解释：**  KIA密级。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Kiaguid *string `json:"kiaguid,omitempty"`

	// **参数解释：**  最后更新时间。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// **参数解释：**  是否为最新版本。  **取值范围：**  - true：是最新版本。 - false：不是最新版本。  **默认取值：**  false。
	Latest *bool `json:"latest,omitempty"`

	// **参数解释：**  是否为最新迭代版本。  **取值范围：**  - true：是最新迭代版本。 - false：不是最新迭代版本。  **默认取值：**  不涉及。
	LatestIteration *bool `json:"latestIteration,omitempty"`

	// **参数解释：**  是否为最新修订版本。  **取值范围：**  - true：是最新修订版本。 - false：不是最新修订版本。  **默认取值：**  不涉及。
	LatestVersion *bool `json:"latestVersion,omitempty"`

	Master *VersionModelMasterViewDto `json:"master,omitempty"`

	// **参数解释：**  修改人。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`

	// **参数解释：**  中文名称。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：**  前序版本实例ID。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	PreVersionId *string `json:"preVersionId,omitempty"`

	// **参数解释：**  软删除标识。  **取值范围：**  - 0：表示未删除。 - 1：表示已删除。  **默认取值：**  0。
	RdmDeleteFlag *int32 `json:"rdmDeleteFlag,omitempty"`

	// **参数解释：**  扩展类型。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	RdmExtensionType *string `json:"rdmExtensionType,omitempty"`

	// **参数解释：**  系统版本。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	RdmVersion *int32 `json:"rdmVersion,omitempty"`

	// **参数解释：**  安全密级。  **取值范围：**  - INTERNAL：内部公开。 - SECRET：秘密。 - CONFIDENTIAL：机密。 - TOP_SECRET：绝密。  **默认取值：**  不涉及。
	SecurityLevel *string `json:"securityLevel,omitempty"`

	Tenant *TenantViewDto `json:"tenant,omitempty"`

	// **参数解释：**  版本号。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Version *string `json:"version,omitempty"`

	// **参数解释：**  业务版本内码。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	VersionCode *int32 `json:"versionCode,omitempty"`

	// **参数解释：**  是否已检出。  **取值范围：**  - true：已检出。 - false：未检出。  **默认取值：**  false。
	WorkingCopy *bool `json:"workingCopy,omitempty"`

	WorkingState *WorkingState `json:"workingState,omitempty"`
}

func (o VersionModelViewDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionModelViewDto struct{}"
	}

	return strings.Join([]string{"VersionModelViewDto", string(data)}, " ")
}
