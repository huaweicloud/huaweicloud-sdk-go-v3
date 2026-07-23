package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MultiViewModelViewDto struct {

	// **参数解释：**  视图实例的唯一标识。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。
	Id *string `json:"id,omitempty"`

	// **参数解释：**  视图中文名称。  **取值范围：**  不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：**  视图描述信息。  **取值范围：**  不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释：**  数据模型的类名称。  **取值范围：**  不涉及。
	ClassName *string `json:"className,omitempty"`

	// **参数解释：**  视图的创建者账号。  **取值范围：**  不涉及。
	Creator *string `json:"creator,omitempty"`

	// **参数解释：**  视图的创建时间，使用UTC+0时间格式，格式为yyyy-MM-ddTHH:mm:ss.SSSZ。  **取值范围：**  不涉及。
	CreateTime *string `json:"createTime,omitempty"`

	// **参数解释：**  视图的更新者账号。  **取值范围：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`

	// **参数解释：**  视图最后更新时间。使用UTC+0时间格式，格式为yyyy-MM-ddTHH:mm:ss.SSSZ。  **取值范围：**  不涉及。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// **参数解释：**  系统版本号。  **取值范围：**  不涉及。
	RdmVersion *int32 `json:"rdmVersion,omitempty"`

	// **参数解释：**  软删除标识。  **取值范围：**  - 0：表示未删除。 - 1：表示已删除。
	RdmDeleteFlag *int32 `json:"rdmDeleteFlag,omitempty"`

	// **参数解释：**  扩展类型，标识数据模型的扩展分类。  **取值范围：**  不涉及。
	RdmExtensionType *string `json:"rdmExtensionType,omitempty"`

	// **参数解释：**  KIA密级。  **取值范围：**
	Kiaguid *string `json:"kiaguid,omitempty"`

	// **参数解释：**  安全密级。  **取值范围：**  - internal：内部公开。 - secret：秘密。 - confidential：机密。 - top_secret：绝密。
	SecurityLevel *string `json:"securityLevel,omitempty"`

	// **参数解释：**  版本号，标识该视图的版本标识。  **取值范围：**  不涉及。
	Version *string `json:"version,omitempty"`

	// **参数解释：**  业务版本内码，用于系统内部版本管理。  **取值范围：**  不涉及。
	VersionCode *int32 `json:"versionCode,omitempty"`

	// **参数解释：**  迭代版本号，标识该视图在当前版本下的迭代次数。  **取值范围：**  不涉及。
	Iteration *int32 `json:"iteration,omitempty"`

	// **参数解释：**  是否为最新版本。  **取值范围：**  - true：是最新版本。 - false：不是最新版本。
	Latest *bool `json:"latest,omitempty"`

	// **参数解释：**  是否为最新修订版本。  **取值范围：**  - true：是最新修订版本。 - false：不是最新修订版本。
	LatestVersion *bool `json:"latestVersion,omitempty"`

	// **参数解释：**  是否为最新迭代版本。  **取值范围：**  - true：是最新迭代版本。 - false：不是最新迭代版本。
	LatestIteration *bool `json:"latestIteration,omitempty"`

	// **参数解释：**  是否已检出。  **取值范围：**  - true：已检出。 - false：未检出。
	WorkingCopy *bool `json:"workingCopy,omitempty"`

	WorkingState *WorkingState `json:"workingState,omitempty"`

	// **参数解释：**  检出人。  **取值范围：**  不涉及。
	CheckOutUserName *string `json:"checkOutUserName,omitempty"`

	// **参数解释：**  检出时间。使用UTC+0时间格式，格式为yyyy-MM-ddTHH:mm:ss.SSSZ。  **取值范围：**  不涉及。
	CheckOutTime *string `json:"checkOutTime,omitempty"`

	// **参数解释：**  前序版本实例ID。  **取值范围：**  不涉及。
	PreVersionId *string `json:"preVersionId,omitempty"`

	Master *MultiViewModelMasterViewDto `json:"master,omitempty"`

	Branch *MultiViewModelBranchViewDto `json:"branch,omitempty"`

	Item *MultiViewItemViewDto `json:"item,omitempty"`

	Tenant *TenantViewDto `json:"tenant,omitempty"`
}

func (o MultiViewModelViewDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiViewModelViewDto struct{}"
	}

	return strings.Join([]string{"MultiViewModelViewDto", string(data)}, " ")
}
