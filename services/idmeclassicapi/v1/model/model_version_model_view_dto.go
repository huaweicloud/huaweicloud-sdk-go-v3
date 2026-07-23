package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionModelViewDto struct {
	Branch *VersionModelBranchViewDto `json:"branch,omitempty"`

	// **参数解释：**  检出时间，使用UTC+0时间格式，格式为yyyy-MM-ddTHH:mm:ss.SSSZ。若实例未检出，返回null。  **取值范围：**  不涉及。
	CheckOutTime *string `json:"checkOutTime,omitempty"`

	// **参数解释：**  检出人名称。若实例未检出，返回null。  **取值范围：**  不涉及。
	CheckOutUserName *string `json:"checkOutUserName,omitempty"`

	// **参数解释：**  版本实例的类名。  **取值范围：**  不涉及。
	ClassName *string `json:"className,omitempty"`

	// **参数解释：**  版本实例的创建时间，使用UTC+0时间格式，格式为yyyy-MM-ddTHH:mm:ss.SSSZ。  **取值范围：**  不涉及。
	CreateTime *string `json:"createTime,omitempty"`

	// **参数解释：**  版本实例的创建者账号。  **取值范围：**  不涉及。
	Creator *string `json:"creator,omitempty"`

	// **参数解释：**  版本实例的描述信息，如工艺变更说明、设备规格备注等。  **取值范围：**  不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释：**  版本实例的唯一标识。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。
	Id *string `json:"id,omitempty"`

	// **参数解释：**  版本实例的迭代版本号，用于标识同一版本下的迭代次数。  **取值范围：**  不涉及。
	Iteration *int32 `json:"iteration,omitempty"`

	// **参数解释：**  KIA密级，用于数据资产的安全标识。  **取值范围：**  不涉及。
	Kiaguid *string `json:"kiaguid,omitempty"`

	// **参数解释：**  版本实例的最后更新时间，使用UTC+0时间格式，格式为yyyy-MM-ddTHH:mm:ss.SSSZ。  **取值范围：**  不涉及。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// **参数解释：**  是否为最新版本。  **取值范围：**  - true：是最新版本。 - false：不是最新版本。
	Latest *bool `json:"latest,omitempty"`

	// **参数解释：**  是否为最新迭代版本。同一版本号下可能存在多次迭代，此字段标识是否为最后一次迭代。  **取值范围：**  - true：是最新迭代版本。 - false：不是最新迭代版本。
	LatestIteration *bool `json:"latestIteration,omitempty"`

	// **参数解释：**  是否为最新修订版本。  **取值范围：**  - true：是最新修订版本。 - false：不是最新修订版本。
	LatestVersion *bool `json:"latestVersion,omitempty"`

	Master *VersionModelMasterViewDto `json:"master,omitempty"`

	// **参数解释：**  版本实例的更新者账号。  **取值范围：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`

	// **参数解释：**  版本实例的中文名称。  **取值范围：**  不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：**  前序版本实例ID，用于追溯版本变更历史。  **取值范围：**  不涉及。
	PreVersionId *string `json:"preVersionId,omitempty"`

	// **参数解释：**  软删除标识，用于M-V模型实例的逻辑删除管理。  **取值范围：**  - 0：表示未删除。 - 1：表示已删除。
	RdmDeleteFlag *int32 `json:"rdmDeleteFlag,omitempty"`

	// **参数解释：**  扩展类型，用于区分不同业务域的M-V模型实例。  **取值范围：**  不涉及。
	RdmExtensionType *string `json:"rdmExtensionType,omitempty"`

	// **参数解释：**  系统版本号，用于数据版本控制与并发冲突检测。每次更新后自动递增。  **取值范围：**  不涉及。
	RdmVersion *int32 `json:"rdmVersion,omitempty"`

	// **参数解释：**  安全密级，标识当前版本实例的访问控制级别。  **取值范围：**  - internal：内部公开。 - secret：秘密。 - confidential：机密。 - top_secret：绝密。
	SecurityLevel *string `json:"securityLevel,omitempty"`

	Tenant *TenantViewDto `json:"tenant,omitempty"`

	// **参数解释：**  版本号。  **取值范围：**  不涉及。
	Version *string `json:"version,omitempty"`

	// **参数解释：**  业务版本内码，用于系统内部的版本排序与比较。  **取值范围：**  不涉及。
	VersionCode *int32 `json:"versionCode,omitempty"`

	// **参数解释：**  是否已检出。本接口要求实例必须处于未检出状态（workingCopy=false）方可更新。  **取值范围：**  - true：已检出（当前被某用户锁定编辑中）。 - false：未检出（处于检入状态，可被管理员更新）。
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
