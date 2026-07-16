package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Flavor struct {

	// **参数解释**：硬件架构类型。 **取值范围**：枚举类型，取值如下： - x86_64：X86架构。 - aarch64：ARM架构。
	Arch *string `json:"arch,omitempty"`

	Billing *BillingInfo `json:"billing,omitempty"`

	// **参数解释**：处理器类型。 **取值范围**：枚举类型，取值如下： - CPU - GPU - ASCEND
	Category *string `json:"category,omitempty"`

	// **参数解释**：规格描述信息。 **取值范围**：不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：规格包含EVS时，EVS存储创建的最大上限(单位：GB)。 **取值范围**：不涉及。
	EvsMaxSize *string `json:"evs_max_size,omitempty"`

	// **参数解释**：规格包含EVS时，EVS存储的sku编码。 **取值范围**：不涉及。
	EvsSkuCode *string `json:"evs_sku_code,omitempty"`

	// **参数解释**：规格类别。 **取值范围**：枚举类型，取值如下： - DEFAULT：CodeLab规格。 - NOTEBOOK：Notebook规格。
	Feature *string `json:"feature,omitempty"`

	// **参数解释**：是否为免费规格。 **取值范围**：布尔类型： - true：免费规格。 - false：非免费规格。
	Free *bool `json:"free,omitempty"`

	// **参数解释**：支持站点类型。 **取值范围**：枚举类型，取值如下： - COMMON：国内与国际站都支持。 - NATIONAL：仅在国内站支持。 - INTERNATIONAL：仅在国际站支持。 - NONE：国内与国际站都不支持。
	GrowSupportType *string `json:"grow_support_type,omitempty"`

	// **参数解释**：规格ID。 **取值范围**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：内存大小。 **取值范围**：不涉及。
	Memory *string `json:"memory,omitempty"`

	// **参数解释**：规格名称。 **取值范围**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：是否告罄。 **取值范围**：布尔类型： - true：告罄。 - false：未告罄。
	SoldOut *bool `json:"sold_out,omitempty"`

	// **参数解释**：规格支持的存储类型。 **取值范围**：枚举类型，取值如下： - EVS：云硬盘。 - OBS：对象存储服务。 - OBSFS：并行文件系统。 - EFS：弹性文件服务（SFS Turbo）
	Storages *[]string `json:"storages,omitempty"`

	// **参数解释**：CPU核数。 **取值范围**：不涉及。
	Vcpus *int32 `json:"vcpus,omitempty"`

	Gpu *GpUsInfo `json:"gpu,omitempty"`

	Ascend *AscendInfo `json:"ascend,omitempty"`
}

func (o Flavor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Flavor struct{}"
	}

	return strings.Join([]string{"Flavor", string(data)}, " ")
}
